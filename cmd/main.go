package main

import (
	"os"

	"github.com/qerdcv/aoc/internal/generic"
)

func main() {
	var (
		arg string
		err error
	)

	args, _ := generic.PopStart(os.Args)

	args, arg = generic.PopStart(args)
	switch arg {
	case "init":
		err = initDay(args)
	case "run":
		err = run()
	}

	if err != nil {
		panic(err)
	}
}
