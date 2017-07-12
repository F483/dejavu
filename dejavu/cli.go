package main

import (
	"bufio"
	// "github.com/f483/dejavu"
	"log"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		log.Println("line", s.Text())
	}
}
