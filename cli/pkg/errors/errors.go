package errors

import (
	"errors"
	"fmt"
	"log"
)

const colorRed = "\033[31m"
const colorReset = "\033[0m"
const errorPrefix = "CLI [error] > "

type CmdRunnerErr struct {
	StatusCode string
	Err        error
}

func (me *CmdRunnerErr) Error() string {
	return fmt.Sprintf("status %s: err %v", me.StatusCode, me.Err)
}

func Generic(code, msg string) *CmdRunnerErr {
	return &CmdRunnerErr{
		StatusCode: code,
		Err:        errors.New(msg),
	}
}

func FatalErr(msg error) {
	log.Fatalf("%s %s %s %s\n", colorRed, errorPrefix, colorReset, msg)
}

func PrintStr(msg string) {
	log.Printf("%s %s %s %s\n", colorRed, errorPrefix, colorReset, msg)
}

func Print(errorMsg error) {
	log.Printf("%s %s %s %s\n", colorRed, errorPrefix, colorReset, errorMsg)
}
