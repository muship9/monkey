package repl

import (
	"bufio"
	"fmt"
	"github.com/shinp09/monkey/lexer"
	"github.com/shinp09/monkey/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	// 入力ソースから読み込み
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// 読み込んだ行を字句解析器のインスタンスに渡す
		line := scanner.Text()
		l := lexer.New(line)

		// 字句解析器が返す全てのトークンを表示
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
