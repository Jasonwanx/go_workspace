package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	FileName string = "demo_exp.xml"
	InputFile *os.File
	InputReader *bufio.Reader
	InputError error
	Count int
	ReadCh = make(chan string)
	)

func ReadFile(fileName string) {
	InputFile, InputError = os.Open(FileName)
	if InputError != nil{
		fmt.Printf("An error occurred on opening the inputfile, error:%s\n", InputError.Error())
		return
	}

	defer func() {
		InputFile.Close()
		close(ReadCh)
	}()

	InputReader = bufio.NewReader(InputFile)
	for{
		Count++
		inputString, readError := InputReader.ReadString('\n')
		if readError == io.EOF{
			return
		}
		ReadCh <- inputString
		fmt.Printf("The %d line is:%s", Count, inputString)
	}
}

func main(){
	go ReadFile(FileName)
	for str := range ReadCh{
		fmt.Printf("---------line:%s", str)
	}
}