// MFP   - Miulti-Function Printers and scanners toolkit
// mains - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// mfp-shell command implementation

package main

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/alexpevzner/mfp/mains"
	"github.com/peterh/liner"
)

// main function for the mfp-shell command
func main() {
	// Setup liner library
	liner := liner.NewLiner()
	defer liner.Close()

	liner.SetCtrlCAborts(true)

	// Read and execute line by line
	for {
		line, err := liner.Prompt("MFP> ")
		if err != nil {
			fmt.Printf("\n")
			break
		}

		err = exec(line)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
}

// exec parses and executes the command
func exec(line string) error {
	// Tokenize string
	tokens, err := tokenize(line)
	if err != nil {
		return err
	}

	// Ignore empty lines
	if len(tokens) == 0 {
		return nil
	}

	// Lookup the command
	cmd := mains.CommandByName(tokens[0])
	if cmd == nil {
		err = fmt.Errorf("%q: command not found", tokens[0])
		return err
	}

	cmd.Main(tokens[1:])

	return nil
}

// tokenize splits string into tokens.
func tokenize(line string) ([]string, error) {
	type tkState int
	const (
		tkSpace   tkState = iota
		tkWord            // Got non-space
		tkQuote           // Got "
		tkQuoteBs         // Got " ... \
		tkHex1            // Got " ... \x
		tkHex2            // Got " ... \xN
		tkOct1            // Got " ... \N
		tkOct2            // Got " ... \NN
	)

	state := tkSpace
	token := ""
	acc := 0
	tokens := []string{}

	for _, c := range line {

		switch state {
		case tkSpace, tkWord:
			if c == '"' {
				state = tkQuote
			} else if unicode.IsSpace(c) {
				if state != tkSpace {
					tokens = append(tokens, token)
					token = ""
					state = tkSpace
				}
			} else {
				state = tkWord
				token += string(c)
			}

		case tkQuote:
			if c == '\\' {
				state = tkQuoteBs
			} else if c == '"' {
				state = tkWord
			} else {
				token += string(c)
			}

		case tkQuoteBs:
			switch c {
			case 'x', 'X':
				acc = 0
				state = tkHex1

			case '0', '1', '2', '3', '4', '5', '6', '7':
				acc = int(c - '0')
				state = tkOct1

			case 'a':
				token += "\a"

			case 'b':
				token += "\b"

			case 'f':
				token += "\f"

			case 'n':
				token += "\n"

			case 'r':
				token += "\r"

			case 't':
				token += "\t"

			case 'v':
				token += "\v"

			default:
				token += string(c)
			}

			if state == tkQuoteBs {
				state = tkQuote
			}

		case tkHex1, tkHex2:
			if n := hexadecimal(c); n >= 0 {
				acc = (acc << 4) | n
				if state == tkHex1 {
					state = tkHex2
				} else {
					token += string(rune(acc))
					state = tkQuote
				}
			} else {
				token += string(rune(acc))
				if c == '"' {
					state = tkWord
				} else {
					token += string(c)
					state = tkQuote
				}
			}

		case tkOct1, tkOct2:
			if n := octal(c); n >= 0 {
				acc = (acc << 3) | n
				if state == tkOct1 {
					state = tkOct2
				} else {
					token += string(rune(acc))
					state = tkQuote
				}
			} else {
				token += string(rune(acc))
				if c == '"' {
					state = tkWord
				} else {
					token += string(c)
					state = tkQuote
				}
			}
		}
	}

	switch state {
	case tkWord:
		tokens = append(tokens, token)

	case tkQuote, tkQuoteBs, tkHex1, tkHex2, tkOct1, tkOct2:
		return nil, errors.New("unterminated string")
	}

	return tokens, nil
}

// octal returns numerical value of octal digit c.
// If c is not octal digit, it returns -1.
func octal(c rune) int {
	if '0' <= c && c <= '9' {
		return int(c - '0')
	}
	return -1
}

// hexadecimal returns numerical value of hexadecimal digit c.
// If c is not hexadecimal digit, it returns -1.
func hexadecimal(c rune) int {
	switch {
	case '0' <= c && c <= '9':
		return int(c - '0')
	case 'a' <= c && c <= 'f':
		return int(c-'a') + 10
	case 'A' <= c && c <= 'F':
		return int(c-'A') + 10
	}
	return -1
}
