package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func newTextFromFile(filename string) string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
	}
	return string(bs)
}

func main() {

	//Option 1:
	var s = os.Args[1]
	q, err := os.Open(s)
	if err != nil {
		fmt.Println("error")
	}
	io.Copy(os.Stdout, q)

	//Option 2:
	fmt.Println(newTextFromFile(os.Args[1]))

}
