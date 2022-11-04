package ast_test

import (
	"testing"

	"github.com/a-yohe1/writing_an_interpreter_in_go/monkey/ast"
	"github.com/a-yohe1/writing_an_interpreter_in_go/monkey/token"
)

func TestString(t *testing.T) {
	// let myVar = anotherVar;
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
