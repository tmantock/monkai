package parser

import (
	"testing"

	"github.com/tmantock/monkai/ast"
	"github.com/tmantock/monkai/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	foobar;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. Got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. Got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. Got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. Got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. Got=%s", "foobar", ident.TokenLiteral())
	}

	// if program == nil {
	// 	t.Fatalf("ParseProgram() returned nil")
	// }

	// if len(program.Statements) != 3 {
	// 	t.Fatalf("program.Statements does not contain 3 statements. Got %d", len(program.Statements))
	// }

	// for _, stmt := range program.Statements {
	// 	returnStmt, ok := stmt.(*ast.ReturnStatement)
	// 	if !ok {
	// 		t.Errorf("stmt not *ast.returnStatement. Got = %T", stmt)
	// 		continue
	// 	}

	// 	if returnStmt.TokenLiteral() != "return" {
	// 		t.Errorf("returnStmt.TokenLiteral not 'return.' Got %q", returnStmt.TokenLiteral())
	// 	}
	// }

	// tests := []struct {
	// 	expectedIdentifier string
	// }{
	// 	{"x"},
	// 	{"y"},
	// 	{"foobar"},
	// }

	// for i, tt := range tests {
	// 	stmt := program.Statements[i]
	// 	if !testLetStatement(t, stmt, tt.expectedIdentifier) {
	// 		return
	// 	}
	// }
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not let. Got %q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. Got %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. Got %s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. Got %s", name, letStmt.Name)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}
	t.FailNow()
}
