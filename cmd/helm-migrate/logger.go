package main

import (
	"log"
	"os"

	"github.com/kyokomi/emoji/v2"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoMessage := emoji.Sprint("INFO: :white_check_mark: ")
	WarningPrefix := emoji.Sprint("WARNING: :warning: ")
	ErrorPrefix := emoji.Sprint("ERROR: :fire:")

	InfoLogger = log.New(os.Stdout, InfoMessage, log.Ldate|log.Ltime)
	WarningLogger = log.New(os.Stdout, WarningPrefix, log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stdout, ErrorPrefix, log.Ldate|log.Ltime)
}
