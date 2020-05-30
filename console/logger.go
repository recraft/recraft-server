package console

import (
	"fmt"
	"log"
)

//LogLevel definition
// 0 - Info (and Errors)
// 1 - Warning
// 2 - Debug
type LogLevel int8

//Logger wip
type Console struct {
	Loglevel LogLevel
}

// Info to the console
func (logger *Console) Error(data ...interface{}) {
	log.Println("[ERROR] ", fmt.Sprint(data...))
}

// Info to the console
func (logger *Console) Info(data ...interface{}) {
	log.Println("[INFO] ", fmt.Sprint(data...))
}

// Warning to the console
func (logger *Console) Warning(data ...interface{}) {
	if logger.Loglevel == 1 {
		log.Println("[WARNING] ", fmt.Sprint(data...))
	}

}

// Debug info to the console
func (logger *Console) Debug(data ...interface{}) {
	if logger.Loglevel == 2 {
		log.Println("[DEBUG] ", fmt.Sprint(data...))
	}

}
