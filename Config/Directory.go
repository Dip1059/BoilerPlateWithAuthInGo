package Config

import (
	"log"
	"os"
)

func DirectoryConfig() {
	CreateSessionDirectory()
}

func CreateSessionDirectory() {
	err := os.MkdirAll("./Storage/Session", 0777)
	if err != nil {
		log.Println(err.Error())
	}
}