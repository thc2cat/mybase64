package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {

	var rawDecodedText []byte
	var err error

	input := stdinToChanByteArray(100)
	for i := range input {

		// fmt.Printf("Starting  decoding %v\n", string(i))
		rawDecodedText, err = base64.StdEncoding.DecodeString(string(i))
		if err != nil {
			fmt.Printf("Error decoding %s\n", string(i))
		}
		fmt.Printf("%s\n", rawDecodedText)
	}
}

func stdinToChanByteArray(length int) chan []byte {
	myoutput := make(chan []byte, length)
	tmp := make([]byte, 128)

	go func(c chan []byte) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			tmp = scanner.Bytes()
			passwd := make([]byte, len(tmp))
			copy(passwd, tmp)
			c <- passwd
		}
		close(c) // close all workers
	}(myoutput)
	return myoutput
}
