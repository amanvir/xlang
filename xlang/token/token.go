package token

type TokenType string

type Token struct {
  Type TokenType
  Literal String
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  IDENT = "IDENT"
  INT = "INT"

  ASSIGN = "="
  PLUS = "+"

  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  LBRACE = "}"

  FUNCTION = "FUNCTION"
  LET = "LET"
)
