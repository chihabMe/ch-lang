package lexer_test

import (
	"testing"

	"github.com/chihabMe/mg-lang/lexer"
	"github.com/chihabMe/mg-lang/token"
)

func TestSimpleString(t *testing.T) {
	input := "=+(),;"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.PLUS, expectedLiteral: "+"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
	}
	l := lexer.New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tt.expectedType != tok.Type {
			t.Fatalf(
				"tests[%d] - token type wrong . exptected=%q , got=%q",
				i,
				tt.expectedType,
				tok.Type,
			)
		}
		if tt.expectedLiteral != tok.Literal {
			t.Fatalf(
				"tests[%d] - token litral . exptected=%q , got=%q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}

func TestFuncDeclartion(t *testing.T) {
	input := `
  fnc max(n1,n2){
  }
  `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{expectedType: token.FUNCTION, expectedLiteral: "fnc"},
		{expectedType: token.IDENT, expectedLiteral: "max"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.IDENT, expectedLiteral: "n1"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.IDENT, expectedLiteral: "n2"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.LBRACE, expectedLiteral: "{"},
		{expectedType: token.RBRACE, expectedLiteral: "}"},
	}
	l := lexer.New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tt.expectedType != tok.Type {
			t.Fatalf(
				"tests[%d] - tokentype wrong. expected=%q, got=%q",
				i,
				tt.expectedType,
				tok.Type,
			)
		}

		if tt.expectedLiteral != tok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt, tok)
		}
	}
}

// func TestNextToken(t *testing.T) {
// 	input := "=+(),;"
// 	tests := []struct {
// 		expectedType    token.TokenType
// 		expectedLiteral string
// 	}{

// 		{token.ASSIGN, "="},
// 		{token.PLUS, "+"},
// 		{token.LPAREN, "("},
// 		{token.RPAREN, ")"},
// 		{token.COMMA, ","},
// 		{token.SEMICOLON, ";"},
// 		{token.EOF, ""},
// 	}
// 	l := lexer.New(input)
// 	for i, tt := range tests {
// 		tok := l.NextToken()
// 		if tok.Type != tt.expectedType {
// 			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
// 				i, tt.expectedType, tok.Type)
// 		}
// 		if tok.Literal != tt.expectedLiteral {
// 			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
// 				i, tt.expectedLiteral, tok.Literal)
// 		}
// 	}

// }

// func TestNextToken2(t *testing.T) {

// 	input := `
// 	let five = 5;
// 	`
// 	tests := []struct {
// 		expectedType    token.TokenType
// 		expectedLiteral string
// 	}{
// 		{token.LET, "let"},
// 		{token.IDENT, "five"},
// 		{token.ASSIGN, "="},
// 		{token.INT, "5"},
// 		{token.SEMICOLON, ";"},
// 	}
// 	l := lexer.New(input)
// 	for i, tt := range tests {
// 		tok := l.NextToken()
// 		log.Println(i, tok, tt)
// 		if tok.Type != tt.expectedType {
// 			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
// 				i, tt.expectedType, tok.Type)
// 		}
// 		if tok.Literal != tt.expectedLiteral {
// 			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
// 				i, tt.expectedLiteral, tok.Literal)
// 		}
// 	}
// }
//
// func TestNextToken3(t *testing.T) {
// 	input := `let five = 5;
//     `
// 	tests := []struct {
// 		expectedType    token.TokenType
// 		expectedLiteral string
// 	}{
// 		{token.LET, "let"},
// 		{token.IDENT, "five"},
// 		{token.ASSIGN, "="},
// 		{token.INT, "5"},
// 		{token.SEMICOLON, ";"}, // Update this line to expect a semicolon token
// 		// Add other test cases as needed
// 	}
//
// 	l := lexer.New(input)
//
// 	for i, tt := range tests {
// 		tok := l.NextToken()
// 		fmt.Println(i, tok)
//
// 		if tok.Type != tt.expectedType {
// 			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
// 				i, tt.expectedType, tok.Type)
// 		}
//
// 		if tok.Literal != tt.expectedLiteral {
// 			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
// 				i, tt.expectedLiteral, tok.Literal)
// 		}
// 	}
// }
