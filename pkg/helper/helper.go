package helper

import (
	"os"
	"strconv"
)

// Stoi parse a string to int, return to default if error.
func Stoi(s string, defaultNum int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultNum
	}
	return i
}

// PanicIfErr will panic if param err is not nil
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func MakeDir(relatePath string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+relatePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
