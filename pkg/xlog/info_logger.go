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

type InfoLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (i *InfoLogger) LogOut(col *ColorType, format *string, v ...interface{}) {
	i.once.Do(func() {
		i.init()
	})
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

func (i *InfoLogger) init() {
	intNum, _ := strconv.Atoi(strings.Split(runtime.Version(), ".")[1])
	if intNum >= 14 {
		i.logger = log.New(os.Stdout, "[INFO] >> ", 64|log.Llongfile|log.Ldate|log.Lmicroseconds)
		return
	}
	i.logger = log.New(os.Stdout, "[INFO] ", log.Llongfile|log.Ldate|log.Lmicroseconds)
}
