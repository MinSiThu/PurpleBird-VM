package lexer

import "go/token"

type Lexer struct {
	position     int //current Character Position
	nextPosition int //next Chararcter Position
	ch           rune
	characters   []rune
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{characters: []rune(input)}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.nextPosition >= len(lexer.characters) {
		lexer.ch = rune(0)
	} else {
		lexer.ch = lexer.characters[lexer.nextPosition]
	}
	lexer.position = lexer.nextPosition
	lexer.nextPosition++
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token
	lexer.skipWhiteSpace()

	if lexer.ch == rune('#') {
		if !isDigit(lexer.peekChar()) {
			lexer.skipComment()
			return lexer.NextToken()
		}
	}

	switch lexer.ch {
	case rune(','):
		tok = newToken(token.COMMA, lexer.ch)

	case rune(0):
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isDigit(lexer.ch) {
			return lexer.readDecimal()
		}

		tok.Literal = lexer.readIdentifier()
		return tok
	}
	lexer.readChar()

	return nil
}

func newToken(tokenType token.Type, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) skipWhiteSpace() {
	for isWhitespace(lexer.ch) {
		lexer.readChar()
	}
}

func (lexer *Lexer) peekChar() rune {
	if lexer.nextPosition >= len(lexer.characters) {
		return rune(0)
	}
	return lexer.characters[lexer.nextPosition]
}

func (lexer *Lexer) skipComment() {
	for lexer.ch != '\n' && lexer.ch != rune(0) {
		lexer.readChar()
	}
	lexer.skipWhiteSpace()
}

//utils
func isDigit(ch rune) bool {
	return rune('0') <= ch && ch <= rune('9')
}

func isWhitespace(ch rune) bool {
	return ch == rune(' ') || ch == rune('\t') || ch == rune('\n') || ch == rune('\r')
}
