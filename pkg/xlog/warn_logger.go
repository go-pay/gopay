package xlog

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type WarnLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (i *WarnLogger) LogOut(col *ColorType, format *string, v ...interface{}) {
	i.once.Do(func() {
		i.init()
	})
	if Level >= WarnLevel {
		if col != nil {
			if format != nil {
				i.logger.Output(3, string(*col)+fmt.Sprintf(*format, v...)+string(Reset))
				return
			}
			i.logger.Output(3, string(*col)+fmt.Sprintln(v...)+string(Reset))
			return
		}
		if format != nil {
			i.logger.Output(3, fmt.Sprintf(*format, v...))
			return
		}
		i.logger.Output(3, fmt.Sprintln(v...))
	}
}

func (i *WarnLogger) init() {
	if Level == 0 {
		Level = DebugLevel
	}
	i.logger = log.New(os.Stdout, "[WARN] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
}
