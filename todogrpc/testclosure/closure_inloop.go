package main

import (
	"log"
)

func main() {

	funcs := make([]func(), 0)

	for i := 1; i < 10; i++ {
		//icopy := i
		funcs = append(funcs, func() {
			log.Printf("i:%d", i)
		})
	}

	for _, f := range funcs {
		f()
	}
}
