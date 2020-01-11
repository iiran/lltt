package helper

import "os"

func MakeDir(relatePath string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+relatePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}