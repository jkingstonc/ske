package src

import (
	"fmt"
)

/*
simple debug wrapper for debugging interfaces
*/
func Log(msg ...interface{}) {
	fmt.Println(msg...)
}

/*
logs to the in-game console (not implemented yet)
*/
func ConsoleLog(msg interface{}) {
}

/*
if the assertion condition is false, then we have an error
*/
func Assert(assertion bool, msg string) {
	if !assertion {
		panic(msg)
	}
}
