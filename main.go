package main

/*

#include "listen.h"

#cgo LDFLAGS: -framework CoreFoundation -framework SystemConfiguration
*/
import "C"
import (
	"fmt"
	"time"
)

func main() {
	go C.registerIPChanges()
	fmt.Println("IP Change Listner started")
	time.Sleep(100 * time.Second)
}

//export goCallback
func goCallback() {
	fmt.Println("Detected IP Change!")
}
