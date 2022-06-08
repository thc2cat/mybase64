package main

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

var (
	equal                = []byte{'='}
	ErrCorruptInputError = base64.CorruptInputError(63)
)

func main() {
	input := stdinToChanByteArray(1)
	for i := range input {
		try(string(i))
	}
}

func try(i string) {
	rawDecodedText, err := base64.StdEncoding.DecodeString(i)
	if err != nil {
		switch {
		case errors.Is(err, ErrCorruptInputError):
			// fmt.Printf("Appending %s to '%s'\n", equal, i)
			tmp := append([]byte(i), equal...)
			try(string(tmp))
		default:
			fmt.Printf("Error decoding %v with '%s'\n", err, i)
		}
	} else {
		fmt.Printf("%s converted to %s\n", i, rawDecodedText)
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
