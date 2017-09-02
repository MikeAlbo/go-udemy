package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// the filename should be passed as a argument to the program when ran
// ex. go run main.go myFile.txt
// use os.args to access the arguments passed to the program  which is of type []string
// use os.Open to open a file
// use theio.Copy function to output the text content to the console

type printFile struct{}

func main() {

	if len(os.Args) != 2 {
		log.Fatal("invalid number of arguments")
		os.Exit(1)
	}
	//t := os.Args[1] // get the file argument
	f, err := os.Open(os.Args[1])

	pf := printFile{}

	if err != nil {
		log.Fatal("Error accessing file")
		os.Exit(1)
	} else {
		io.Copy(pf, f)
	}

	//fmt.Println("Input file", f)
}

func (printFile) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
