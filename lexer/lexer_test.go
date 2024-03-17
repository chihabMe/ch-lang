package lexer_test

import (
	"testing"

	"github.com/chihabMe/mg-lang/lexer"
	"github.com/chihabMe/mg-lang/token"
)

func TestFuncDeclaration(t *testing.T) {
	input := `
  fnc add(n1,n2){
    n1+n2;
  }
  set num1 = 5 ;
  set num2 = 10 ;
  set result = add(num1,num2);
  `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{expectedType: token.FUNCTION, expectedLiteral: "fnc"},
		{expectedType: token.IDENT, expectedLiteral: "add"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.IDENT, expectedLiteral: "n1"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.IDENT, expectedLiteral: "n2"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.LBRACE, expectedLiteral: "{"},
		{expectedType: token.IDENT, expectedLiteral: "n1"},
		{expectedType: token.PLUS, expectedLiteral: "+"},
		{expectedType: token.IDENT, expectedLiteral: "n2"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"}, // This line was corrected
		{expectedType: token.RBRACE, expectedLiteral: "}"},
		// end of func
		// set num1 = 5 ;
		// set num2 = 10;
		// set result = add(num1,num2);

		{expectedType: token.SET, expectedLiteral: "set"},
		{expectedType: token.IDENT, expectedLiteral: "num1"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.INT, expectedLiteral: "5"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},

		{expectedType: token.SET, expectedLiteral: "set"},
		{expectedType: token.IDENT, expectedLiteral: "num2"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.INT, expectedLiteral: "10"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},

		{expectedType: token.SET, expectedLiteral: "set"},
		{expectedType: token.IDENT, expectedLiteral: "result"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.IDENT, expectedLiteral: "add"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.IDENT, expectedLiteral: "num1"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.IDENT, expectedLiteral: "num2"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.EOF, expectedLiteral: ""},
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
			t.Fatalf(
				"tests[%d] - literal wrong. expected=%q, got=%q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}

func TestOperators(t *testing.T) {
	input := `
  set num = 10 ;
  *-+/ ;
  5 < 20 > 10 ;
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SET, "set"},
		{token.IDENT, "num"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.ASTERISK, "*"},
		{token.MINUS, "-"},
		{token.PLUS, "+"},
		{token.SLASH, "/"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.ST, "<"},
		{token.INT, "20"},
		{token.LT, ">"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
	}
	l := lexer.New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tt.expectedType != tok.Type {
			t.Fatalf("tests[%d] - wrong type expected=%q , got=%q", i, tt.expectedType, tok.Type)
		}
		if tt.expectedLiteral != tok.Literal {
			t.Fatalf(
				"tests[%d] -  wrong literal expected=%q , got=%q ",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}

	}
}

func TestLongKeywords(t *testing.T) {
	input := `
  fnc isPositive(num){
    if num < 0 {
      return true ;
    } else {
      return false ;
    }
  }
    set valid = true == true ;
    set Invalid = false != false ;
  `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FUNCTION, "fnc"},
		{token.IDENT, "isPositive"},
		{token.LPAREN, "("},
		{token.IDENT, "num"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.IDENT, "num"},
		{token.ST, "<"},
		{token.INT, "0"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.SET, "set"},
		{token.IDENT, "valid"},
		{token.ASSIGN, "="},
		{token.TRUE, "true"},
		{token.EQ, "=="},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.SET, "set"},
		{token.IDENT, "Invalid"},
		{token.ASSIGN, "="},
		{token.FALSE, "false"},
		{token.NOT_EQ, "!="},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := lexer.New(input)
	for i, tk := range tests {
		tok := l.NextToken()
		if tk.expectedType != tok.Type {
			t.Fatalf(
				"tests[%d] - tokentype wrong. expected=%q, got=%q",
				i,
				tk.expectedType,
				tok.Type,
			)
		}

		if tk.expectedLiteral != tok.Literal {
			t.Fatalf(
				"tests[%d] - tokentype wrong. expected=%q, got=%q",
				i,
				tk.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
