package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	cnt := count(os.Stdin, *lines)
	fmt.Println(cnt)
}
