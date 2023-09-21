package gologger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

const version float32 = 1.1

var lock = &sync.Mutex{}

type logger struct {
	logType []LogLevel
}

var loggerInstance *logger

func getInstance() *logger {
	if loggerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if loggerInstance == nil {
			fmt.Println("Creating single instance now.")
			loggerInstance = &logger{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return loggerInstance
}

func (l GoLogger) Error(msg string) {

}

// func (t LogType) String() string {
// 	switch t {
// 	case Debug:
// 		return "debug"
// 	case Info:
// 		return "info"
// 	case Error:
// 		return "error"
// 	case Fatal:
// 		return "fatal"
// 	}
// 	return "unknown"
// }

type GoLogger interface {
	Debug(msg string, err error) error
	Info(msg string, err error) error
	Error(msg string, err error) error
	Fatal(msg string, err error) error
}

func registerLog() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func Register() {
	registerLog()
}
