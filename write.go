package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	InfoColor    = "\033[1;34m"
	NoticeColor  = "\033[1;36m"
	WarningColor = "\033[1;33m"
	ErrorColor   = "\033[1;31m"
	DebugColor   = "\033[0;36m"
	SuccessColor = "\033[1;32m"
	ResetColor   = "\033[0;0m"
)

func Write(msg string, args... interface{}) {
	fmt.Printf(msg + "\n", args...)
}

func Prompt(defaultChoice bool, msg string, args... interface{}) bool {
	var dcText string
	if defaultChoice {
		dcText = "[Y/n]"
	} else {
		dcText = "[y/N]"
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(msg + " " + dcText + "\n", args...)
	str, _ := reader.ReadString('\n')
	switch strings.ToLower(strings.TrimSpace(str)) {
	case "y", "Y":
		return true
	case "n", "N":
		return false
	default:
		return defaultChoice
	}
}