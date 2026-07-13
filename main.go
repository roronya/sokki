package main

import (
	"log"
	"os"

	"github.com/roronya/sokki/evaluator"
	"github.com/roronya/sokki/lexer"
	"github.com/roronya/sokki/parser"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("invalid arguments. usage: sokki input_file output_file")
	}

	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := string(bs)

	l := lexer.New(input)
	p := parser.New(l)
	ast := p.ParseDocument()
	evaluated := evaluator.Eval(ast)

	err = os.WriteFile(os.Args[2], []byte(evaluated), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
