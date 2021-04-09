package main

/*
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Hydrogen/src/evaluator"
	object "github.com/Hydrogen/src/evaluator/objects"
	"github.com/Hydrogen/src/lexer"
	"github.com/Hydrogen/src/parser"
)

const PROMPT = ">>"

func main() {
	fmt.Println("Hydrogen Beta 0.0.1v")
	fmt.Println("Avalaible commands: ")
	fmt.Println("run ${fileáth}  -  : run example1.hg")
	in := os.Stdin
	sto := os.Stdout
	scanner := bufio.NewScanner(in)

	env := object.NewEnvironemnt()
	for {
		fmt.Fprintf(sto, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			continue
		}
		line := scanner.Text()

		textPush := strings.Split(line, " ")

		switch textPush[0] {
		case "run":
			filepathh := textPush[1]
			fileRead, err := ioutil.ReadFile(filepathh)
			if err != nil {
				fmt.Printf("Dont exist File: %s \n", filepathh)
				continue
			}
			fileExt := filepath.Ext(filepathh)
			if fileExt != ".hg" {
				fmt.Printf("The file: %s has no extension .hg \n", filepathh)
				continue
			}
			readStr := string(fileRead)
			//fmt.Println(readStr)
			l := lexer.New(readStr)
			p := parser.New(l)
			program := p.ParseProgram()

			if len(p.Errors()) != 0 {
				printParserErrors(sto, p.Errors())
				continue
			}
			evaluated := evaluator.Eval(program, env)
			if evaluated != nil {
				io.WriteString(sto, evaluated.Inspect())
				io.WriteString(sto, "\n")
			}
		default:
			continue
		}
		fmt.Printf("\n")
	}

}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Hydrogen v0.0.1")
	io.WriteString(out, "Woops! there is a typing error here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
*/
