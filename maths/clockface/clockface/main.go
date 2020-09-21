package main

import (
	"os"
	"time"

	"github.com/qiwenmin/lgwt/maths/clockface/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
