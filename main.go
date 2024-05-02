package main

import (
	"github.com/avearmin/go-json-parser/internal/lexer"
	"github.com/avearmin/go-json-parser/internal/parser"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	var input []byte
	var err error
	switch len(args) {
	case 0:
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalln(err)
		}
	case 1:
		file, err := os.Open(args[0])
		if err != nil {
			log.Fatalln(err)
		}

		input, err = io.ReadAll(file)
		if err != nil {
			log.Fatalln(err)
		}

		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalln("too many arguments, need 1")
	}

	l := lexer.New(string(input))
	p := parser.New(l)

	if _, err := p.ParseJSON(); err != nil {
		log.Fatalln("invalid JSON: " + err.Error())
	}

	log.Println("JSON valid")
	os.Exit(0)
}
