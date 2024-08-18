package utils

import (
	"log"
	"os"
)

func InitLogger() {
	log.SetOutput(os.Stdout)
	// Lshortfile 只会显示文件名和行号
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}
