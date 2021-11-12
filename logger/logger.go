package logger

import (
	"fmt"
	"time"

	"github.com/gookit/color"
)

var dtFmt = "2006-01-02 15:04:05"

func Info(msg string) {
	fmt.Printf("[%s] %s\n", color.White.Render(time.Now().Format(dtFmt)), color.Cyan.Render(msg))
}

func Debug(msg string) {
	fmt.Printf("[%s] %s\n", color.White.Render(time.Now().Format(dtFmt)), color.Magenta.Render(msg))
}

func Error(msg string) {
	fmt.Printf("[%s] %s\n", color.White.Render(time.Now().Format(dtFmt)), color.Red.Render(msg))
}

func Err(msg string, err error) {
	fmt.Printf("[%s] %s\n", color.White.Render(time.Now().Format(dtFmt)), color.Red.Render(msg, fmt.Sprintf(" (%s)", err.Error())))
}