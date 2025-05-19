package evaluator

import (
	"testing"

	"github.com/mrcruz117/mnky-interpreter/object"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"quote(5)", "5"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		quote, ok := evaluated.(*object.Quote)
		if !ok {
			t.Fatalf("expercted *object.Quote. got=%T (%+v)", evaluated, evaluated)
		}
		if quote.Node == nil {
			t.Fatalf("quote.Node is nil.")
		}
		if quote.Node.String() != tt.expected {
			t.Errorf(
				"quote.Node.String() wrong. got=%q, expected=%q",
				quote.Node.String(), tt.expected,
			)
		}
	}
}
