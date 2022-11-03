package lexer

import (
	"github.com/shinp09/monkey/token"
)

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置 ( 現在の文字を指し示す )
	readPosition int  // これから読み込む位置 ( 現在の文字の次 )
	ch           byte // 現在検査中の文字
}

// New 入力された文字を Lexer 構造体に格納し、位置情報を更新する
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.leadChar()
	return l
}

// leadChar 現在の読み込み位置を確認し、現在位置と次に対象となる位置を取得する
func (l *Lexer) leadChar() {

	// 入力が終端に到達したかをチェックし、してなければ次の文字をセットする
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken 受け取った文字に応じて必要な token を返す
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.leadChar()
	return tok
}

// newToken Token 構造体に情報をマッピング
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
