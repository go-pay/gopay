package xlog

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type ErrorLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (e *ErrorLogger) LogOut(col *ColorType, format *string, v ...any) {
	e.once.Do(func() {
		e.init()
	})
	if Level >= ErrorLevel {
		if col != nil {
			if format != nil {
				_ = e.logger.Output(3, string(*col)+fmt.Sprintf(*format, v...)+string(Reset))
				return
			}
			_ = e.logger.Output(3, string(*col)+fmt.Sprintln(v...)+string(Reset))
			return
		}
		if format != nil {
			_ = e.logger.Output(3, fmt.Sprintf(*format, v...))
			return
		}
		_ = e.logger.Output(3, fmt.Sprintln(v...))
	}
}

func (e *ErrorLogger) init() {
	if Level == 0 {
		Level = DebugLevel
	}
	e.logger = log.New(os.Stdout, "[ERROR] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
}
