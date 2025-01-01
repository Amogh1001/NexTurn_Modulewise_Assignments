package main

import (
	"os"
	questions "question-management/questions"

	"github.com/sirupsen/logrus"
)

// AppLogger is a globally accessible logger
var AppLogger *logrus.Logger

func init() {
	AppLogger = logrus.New()
	AppLogger.SetFormatter(&logrus.JSONFormatter{}) // Set the log format to JSON
	AppLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	}) // Set the log format to text
	AppLogger.SetOutput(os.Stdout)        // Set the output to stdout
	AppLogger.SetLevel(logrus.DebugLevel) // Set the log level to debug
}

func main() {
	questions.AddQuestions()
	questions.GetAllQuestions()
}
