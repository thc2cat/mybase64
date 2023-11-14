package main

// mybase64 read stdin and output decoded base64 if ascii text
//
// WARNING : output is different from linux base64 command
//
// Versions :
//
// v0.3 : use isASCII
// v0.2 : add '=' to corrupted entries
// v0.1 : first running

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
	// MaxASCII             = (byte)(unicode.MaxASCII)
	MaxASCII = (byte)(127)
)

func main() {
	input := stdinToChanByteArray(1)
	for i := range input {
		try(string(i))
	}
}

// try to decode based64 encodings.
// if error is a Corrupted input at 63, try to add "="
func try(input string) {
	rawDecodedText, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		switch {
		case errors.Is(err, ErrCorruptInputError):
			// fmt.Printf("Appending %s to '%s'\n", equal, i)
			tmp := append([]byte(input), equal...)
			try(string(tmp))
		default:
			fmt.Printf("Error decoding %v with '%s'\n", err, input)
		}
	} else {
		if isASCII(rawDecodedText) {
			fmt.Printf("%v converted to %s\n", input, rawDecodedText)
		} else {
			fmt.Printf("%s converted to non ASCII output", input)
		}
	}
}

// Standard stdin to chan ( should be in a utils lib )
func stdinToChanByteArray(length int) chan []byte {
	myoutput := make(chan []byte, length)
	tmp := make([]byte, length)

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

// testing non ASCII []byte array like 🧡💛💚💙💜
func isASCII(s []byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > MaxASCII {
			return false
		}
	}
	return true
}
