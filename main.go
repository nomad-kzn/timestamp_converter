package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const convertArg = "convert"

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("invalid args")
		return
	}
	if strings.ToLower(args[1]) == convertArg {
		timestamp, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			fmt.Println("invalid timestamp")
			return
		}

		tm := time.Unix(timestamp, 0)
		fmt.Println(tm)
		return
	}

	fmt.Println("invalid args")
}
