package logging

import (
	// "io"
	"log"
	"os"
	// "os"
)

var (
	Info *log.Logger
	Error *log.Logger
	Serve *log.Logger
)

func init() {
	
	// if _, err := os.Stat("/stor_data/logs.log"); os.IsNotExist(err) {
	// 	_, err = os.Create("/stor_data/logs.log")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// file, err := os.OpenFile("stor_data/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalln("Failed to open log file:", err)
	// }
	// Info = log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	// Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Serve = log.New(os.Stdout, "", log.Ltime)
}
