package main

const (
	TOKEN_ILLEGAL  = "ILLEGAL"
	TOKEN_EOF      = "EOF"
	TOKEN_INT      = "INT"
	TOKEN_PLUS     = "+"
	TOKEN_MINUS    = "-"
	TOKEN_ASTERISK = "*"
	TOKEN_SLASH    = "/"
	TOKEN_LPAREN   = "("
	TOKEN_RPAREN   = ")"
	TOKEN_IDENT    = "IDENT"
	TOKEN_ASSIGN   = "="
)

type Token struct {
	Type    string
	Literal string
}
