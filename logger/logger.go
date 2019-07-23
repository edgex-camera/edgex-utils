package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
)

type printClient struct{}

var myLogger = log.New(os.Stdout, "", log.Llongfile|log.Ldate|log.Ltime)

func NewPrintClient() logger.LoggingClient {
	return &printClient{}
}

func print(level string, msg string, args ...interface{}) {
	myLogger.Output(3, fmt.Sprintf(level+": "+msg, args...))
}

func (c *printClient) Debug(msg string, args ...interface{}) {
	print("Debug", msg, args...)
}

func (c *printClient) Error(msg string, args ...interface{}) {
	print("Error", msg, args...)
}

func (c *printClient) Info(msg string, args ...interface{}) {
	print("Info", msg, args...)
}

func (c *printClient) Trace(msg string, args ...interface{}) {
	print("Trace", msg, args...)
}

func (c *printClient) Warn(msg string, args ...interface{}) {
	print("Warn", msg, args...)
}

func (c *printClient) SetLogLevel(logLevel string) error {
	return nil
}
