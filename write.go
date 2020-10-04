package gocli

import "fmt"

const (
	InfoColor    = "\033[1;34m"
	NoticeColor  = "\033[1;36m"
	WarningColor = "\033[1;33m"
	ErrorColor   = "\033[1;31m"
	DebugColor   = "\033[0;36m"
	SuccessColor = "\033[1;32m"
	ResetColor   = "\033[0;0m"
)

func Write (msg string, args... interface{}) {
	fmt.Printf(msg, args...)
}

