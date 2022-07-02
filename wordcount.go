package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}

func readInput() (string, error) {
	args := os.Args[1:]
	if len(args) == 0 {
		return "", nil
	}
	return args[0], nil
}

func countWords(src string) (int, error) {
	var result int = len(strings.Split(src, " "))
	if reflect.DeepEqual(src, "") {
		result = 0
	}
	return result, nil
}

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}

	count, err := countWords(src)
	if err != nil {
		fail(err)
	}
	fmt.Println(count)
}
