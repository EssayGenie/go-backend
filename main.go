package main

import (
	"go-backend/cmd"
	"log"
	"os"
)

func main() {
	// save here in the directory as logfile
	logFile, err := os.OpenFile("logfile", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Unable to set logfile:", err.Error())
	}
	log.SetOutput(logFile)
	cmd.Execute()
}
