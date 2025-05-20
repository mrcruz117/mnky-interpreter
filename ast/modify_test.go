package ast

import (
	"reflect"
	"testing"
)

func TestModify(t *testing.T) {
	one := func() Expression { return &IntegerLiteral{Value: 1} }
	two := func() Expression { return &IntegerLiteral{Value: 2} }

	turnOneIntoTwo := func(node Node) Node {
		integer, ok := node.(*IntegerLiteral)
		if !ok {
			return node
		}
		if integer.Value != 1 {
			return node
		}
		integer.Value = 2
		return integer
	}

	tests := []struct {
		input    Node
		expected Node
	}{
		{
			input:    one(),
			expected: two(),
		},
		{
			input: &Program{
				Statements: []Statement{
					&ExpressionStatement{
						Expression: one(),
					},
				},
			},
			expected: &Program{
				Statements: []Statement{
					&ExpressionStatement{
						Expression: two(),
					},
				},
			},
		},
		{
			&InfixExpression{Left: one(), Operator: "+", Right: two()},
			&InfixExpression{Left: two(), Operator: "+", Right: two()},
		},
		{
			&InfixExpression{Left: two(), Operator: "+", Right: one()},
			&InfixExpression{Left: two(), Operator: "+", Right: two()},
		},
		{
			&PrefixExpression{Operator: "-", Right: one()},
			&PrefixExpression{Operator: "-", Right: two()},
		},
	}
	for _, tt := range tests {
		modified := Modify(tt.input, turnOneIntoTwo)

		equal := reflect.DeepEqual(modified, tt.expected)
		if !equal {
			t.Errorf("Modify(%#v) = %#v, expected %#v", tt.input, modified, tt.expected)
		}
	}
}
