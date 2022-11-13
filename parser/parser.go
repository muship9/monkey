package parser

import (
	"github.com/shinp09/monkey/ast"
	"github.com/shinp09/monkey/lexer"
	"github.com/shinp09/monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 2つトークンを読み込み、curToken と peekToken の両方がセットされる
	p.nextToken()
	p.nextToken()

	return p
}

// 字句解析器インスタンスへのポインタの NextToken() を繰り返し呼び、入力から次のトークンを取得する
func (p *Parser) nextToken() {
	// 現在調べているトークン
	p.curToken = p.peekToken
	// 次のトークン
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
