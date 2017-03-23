package repl

import (
	"bufio"
	"fmt"
	"github.com/jamesroutley/monkey/lexer"
	"github.com/jamesroutley/monkey/token"
	"io"
)

// PROMPT is the prompt string to print at the repl.
const PROMPT = ">> "

// Start starts the Monkey repl.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}
}
