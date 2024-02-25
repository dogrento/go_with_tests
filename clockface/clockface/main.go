package main

import (
	"os"
	"time"

  "go_with_tests/clockface"
)

func main() {
	t := time.Now()

  clockface.SVGWriter(os.Stdout, t)
}

