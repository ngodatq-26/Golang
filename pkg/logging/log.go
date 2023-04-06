package logging

import (
	"log"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Fatal(v ...interface{}) {
	log.Fatalln()
}
