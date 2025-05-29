package parser_test

import (
	"testing"

	"github.com/timocheu/kalayo/ast"
	"github.com/timocheu/kalayo/lexer"
	"github.com/timocheu/kalayo/parser"
)

func TestVarStatements(t *testing.T) {
	input := `
	var a 10;
	var = 1000;
	var 44223;
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParserProgram() returned nil or failed")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program Statements does not contain 3 Statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"a"},
		{"b"},
		{"c"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not 'var'. got=%q", s.TokenLiteral())
	}

	// Assert that s is an letStatement which is under the Statement interface
	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if varStmt.Name.Value != name {
		t.Errorf("varStmt.Name.Value not '%s'. got=%s",
			name, varStmt.Name.Value)
		return false
	}

	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("varStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, varStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParseErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
