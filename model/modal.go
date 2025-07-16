package model

import (
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

const WorkerCount = 2

var (
	Dl  = filepath.Join(os.Getenv("USERPROFILE"), "Downloads")
	Ds  = filepath.Join(os.Getenv("USERPROFILE"), "Desktop")
	Typ = map[string][]string{
		"Resimler":    {".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff", ".svg", ".ico"},
		"Videolar":    {".mp4", ".avi", ".mkv", ".mov", ".flv", ".wmv", ".webm", ".mpeg"},
		"Belgeler":    {".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".txt", ".odt", ".rtf", ".csv"},
		"Kurulumlar":  {".exe", ".msi", ".zip", ".rar", ".7z", ".tar", ".gz", ".deb", ".rpm"},
		"Müzik":       {".mp3", ".wav", ".flac", ".aac", ".ogg", ".wma", ".m4a"},
		"Kodlar":      {".go", ".py", ".js", ".ts", ".html", ".css", ".java", ".c", ".cpp", ".cs", ".json", ".xml", ".sh", ".bat", ".jar"},
		"Tasarımlar":  {".psd", ".xd", ".fig", ".sketch", ".ai", ".indd"},
		"E-kitaplar":  {".epub", ".mobi", ".azw3", ".fb2"},
		"Disk Kalıpları": {".iso", ".img", ".bin", ".cue"},
		"Veritabanı":  {".sql", ".db", ".sqlite", ".mdb", ".accdb"},
		"Yedekler":    {".bak", ".tmp", ".old"},
	}
	
)
type File struct {
	gorm.Model
	Time string
	Name string
	Type string
	Size int64
	Path string
}