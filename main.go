package main

import (
	"fmt"
)

func main() {
	msg := seyHello("go")
	
	println(msg)

}

func seyHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}