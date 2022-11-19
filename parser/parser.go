package parser

import (
	"fmt"
	"github.com/shinp09/monkey/ast"
	"github.com/shinp09/monkey/lexer"
	"github.com/shinp09/monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	errors    []string
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// 2つトークンを読み込み、curToken と peekToken の両方がセットされる
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// 字句解析器インスタンスへのポインタの NextToken() を繰り返し呼び、入力から次のトークンを取得する
func (p *Parser) nextToken() {
	// 現在調べているトークン
	p.curToken = p.peekToken
	// 次のトークン
	p.peekToken = p.l.NextToken()
}

// ParseProgram ast を作成する
func (p *Parser) ParseProgram() *ast.Program {
	// ルートノードの作成
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// 入力されたトークンを EOF になるまで読み込む
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// parseLetStatement 現在見ているトークンに基づいて *ast.LetStatement を構築
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek peekToken の型が正しい場合は nextToken を実行する
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}

}

// parseReturnStatement SEMICOLON がきたら stmt を返し、きてなければ次のtokenを読み込む
func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{
		Token: p.curToken,
	}
	p.nextToken()
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
