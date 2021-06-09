package xlog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/go-pay/gopay/pkg/util"
)

type DebugLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (i *DebugLogger) logOut(col *ColorType, format *string, v ...interface{}) {
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

func (i *DebugLogger) init() {
	if util.String2Int(strings.Split(runtime.Version(), ".")[1]) >= 14 {
		i.logger = log.New(os.Stdout, "[DEBUG] >> ", 64|log.Llongfile|log.Ldate|log.Lmicroseconds)
		return
	}
	i.logger = log.New(os.Stdout, "[DEBUG] ", log.Llongfile|log.Ldate|log.Lmicroseconds)
}
