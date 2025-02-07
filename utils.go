package main

import (
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
