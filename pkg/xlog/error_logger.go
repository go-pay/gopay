package xlog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type ErrorLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (e *ErrorLogger) LogOut(col *ColorType, format *string, v ...interface{}) {
	e.once.Do(func() {
		e.init()
	})
	if col != nil {
		if format != nil {
			e.logger.Output(3, string(*col)+fmt.Sprintf(*format, v...)+string(Reset))
			return
		}
		e.logger.Output(3, string(*col)+fmt.Sprintln(v...)+string(Reset))
		return
	}
	if format != nil {
		e.logger.Output(3, fmt.Sprintf(*format, v...))
		return
	}
	e.logger.Output(3, fmt.Sprintln(v...))
}

func (e *ErrorLogger) init() {
	intNum, _ := strconv.Atoi(strings.Split(runtime.Version(), ".")[1])
	if intNum >= 14 {
		e.logger = log.New(os.Stdout, "[ERROR] >> ", 64|log.Llongfile|log.Ldate|log.Lmicroseconds)
		return
	}
	e.logger = log.New(os.Stdout, "[ERROR] ", log.Llongfile|log.Ldate|log.Lmicroseconds)
}
