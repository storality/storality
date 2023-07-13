package shout

import (
	"io"
	"log"
	"os"
)

var (
	Info *log.Logger
	Error *log.Logger
)

func init() {
	logsFile := "./stor_data/logs.log"
	if _, er := os.Stat(logsFile); os.IsNotExist(er) {
		_, er = os.Create(logsFile)
		if er != nil {
			log.Fatal(er)
		}
	}

	file, er := os.OpenFile(logsFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if er != nil {
		log.Fatalln("Failed to open log file:", er)
	}

	multiWriter := io.MultiWriter(file, os.Stdout)
	Info = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}