package main

import (
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
	for cat := range typ {
		ch := make(chan fs.FileInfo, 50)
		queues[cat] = ch
		for i := 0; i < workerCount; i++ {
			go worker(cat, ch)
		}
	}

	t := time.NewTicker(10 * time.Second)
	defer t.Stop()

	for {
		files, err := os.ReadDir(dl)
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
			for cat, exts := range typ {
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
	dest := filepath.Join(ds, cat)
	_ = os.MkdirAll(dest, 0755)

	for file := range ch {
		src := filepath.Join(dl, file.Name())
		dst := filepath.Join(dest, file.Name())
		dst = avoidConflict(dst)

		if err := os.Rename(src, dst); err == nil {
			fmt.Println("Taşındı:", file.Name(), "->", cat)
		}
	}
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
