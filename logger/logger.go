package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLogger() {
	console := zerolog.ConsoleWriter{
		Out: os.Stdout,
	}
	file, err := os.OpenFile(
		"logs/app.log",
		os.O_APPEND|os.O_CREATE|os.O_RDWR,
		0666,
	)
	if err != nil {
		fmt.Println("init logger err: " + err.Error())
	}
	multi := zerolog.MultiLevelWriter(
		console,
		file,
	)
	Log = zerolog.New(multi).
		With().
		Timestamp().
		Logger()

}
