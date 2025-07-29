package main

import "os"

func put_str(s string) string {
	return s
}

func main() {
	word := os.Args[1]
	os.Stdout.Write([]byte(word + "\n"))
}
