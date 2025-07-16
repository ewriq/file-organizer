package main

import (
	database "file-organizer/database"
	"file-organizer/model"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Cronjob() {
	queues := make(map[string]chan fs.FileInfo)
	for cat := range  model.Typ {
		ch := make(chan fs.FileInfo, 50)
		queues[cat] = ch
		for i := 0; i < model.WorkerCount; i++ {
			go worker(cat, ch)
		}
	}

	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		files, err := os.ReadDir(model.Dl)
		if err != nil {
			log.Println("Read error:", err)
			continue
		}

		for _, entry := range files {
			info, err := entry.Info()
			if err != nil || info.IsDir() {
				continue
			}
			ext := strings.ToLower(filepath.Ext(info.Name()))
			for cat, exts := range  model.Typ {
				if contains(exts, ext) {
					queues[cat] <- info
					break
				}
			}
		}
		<-t.C
	}
}

func worker(cat string, ch chan fs.FileInfo) {
	dest := filepath.Join(model.Ds, cat)
	_ = os.MkdirAll(dest, 0755)

	for file := range ch {
		src := filepath.Join(model.Dl, file.Name())
		dst := filepath.Join(dest, file.Name())
		dst = avoidConflict(dst)

		if err := os.Rename(src, dst); err == nil {
			fmt.Println("Taşındı:", file.Name(), "->", cat)

			ext := strings.ToLower(filepath.Ext(file.Name()))
			typeName := getCategory(ext) 
			fmt.Println(typeName)
			err := database.Add(file.Name(), typeName, file.Size(), dst)
			if err != nil {
				log.Println("DB'ye eklenemedi:", err)
			}
		} else {
			log.Println("Dosya taşıma hatası:", err)
		}
	}
}


func getCategory(ext string) string {
	for cat, exts := range model.Typ {
		for _, e := range exts {
			if e == ext {
				return cat
			}
		}
	}
	return "Bilinmeyen"
}



func contains(list []string, ext string) bool {
	for _, e := range list {
		if e == ext {
			return true
		}
	}
	return false
}

func avoidConflict(p string) string {
	ext := filepath.Ext(p)
	base := strings.TrimSuffix(filepath.Base(p), ext)
	dir := filepath.Dir(p)

	for i := 1; ; i++ {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			return p
		}
		p = filepath.Join(dir, fmt.Sprintf("%s (%d)%s", base, i, ext))
	}
}
