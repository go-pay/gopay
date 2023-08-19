package xlog

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type InfoLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (i *InfoLogger) LogOut(col *ColorType, format *string, v ...any) {
	i.once.Do(func() {
		i.init()
	})
	if Level >= InfoLevel {
		if col != nil {
			if format != nil {
				_ = i.logger.Output(3, string(*col)+fmt.Sprintf(*format, v...)+string(Reset))
				return
			}
			_ = i.logger.Output(3, string(*col)+fmt.Sprintln(v...)+string(Reset))
			return
		}
		if format != nil {
			_ = i.logger.Output(3, fmt.Sprintf(*format, v...))
			return
		}
		_ = i.logger.Output(3, fmt.Sprintln(v...))
	}
}

func (i *InfoLogger) init() {
	if Level == 0 {
		Level = DebugLevel
	}
	i.logger = log.New(os.Stdout, "[INFO] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
}
