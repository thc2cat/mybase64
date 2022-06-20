package main

// mybase64 read stdin and output decoded base64 if ascii text
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
	"unicode"
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

// try to decode based64 encodings.
// if error is a Corrupted input at 63, try to add "="
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
		if isASCII(rawDecodedText) {
			fmt.Printf("%s converted to %s\n", i, rawDecodedText)
		}
	}
}

// Standard stdin to chan ( should be in a utils lib )
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

// testing non ASCII []byte array like 🧡💛💚💙💜
func isASCII(s []byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > (byte)(unicode.MaxASCII) {
			return false
		}
	}
	return true
}
