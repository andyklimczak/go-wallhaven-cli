package logger

import (
	"fmt"
	"os"
)

type GoHavenLogger struct {
	Verbose bool
}

func DefaultGoHavenLogger(verbose bool) *GoHavenLogger {
	return &GoHavenLogger{
		Verbose: verbose,
	}
}

func (l *GoHavenLogger) Debug(format string, v ...any) {
	if l.Verbose {
		fmt.Println()
		fmt.Printf(format, v...)
		fmt.Println()
	}
}

func (l *GoHavenLogger) Info(format string, v ...any) {
	fmt.Println()
	fmt.Printf(format, v...)
	fmt.Println()
}

func (l *GoHavenLogger) Error(format string, v ...any) {
	fmt.Println()
	fmt.Printf(format, v...)
	fmt.Println()
}

func (l *GoHavenLogger) Fatal(format string, v ...any) {
	fmt.Println()
	fmt.Printf(format, v...)
	fmt.Println()
	os.Exit(-1)
}
