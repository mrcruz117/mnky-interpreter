package ast

type ModifierFunc func(Node) Node

func Modify(node Node, modifier ModifierFunc) Node {
	switch node := node.(type) {

	case *Program:
		for i, statement := range node.Statements {
			node.Statements[i], _ = Modify(statement, modifier).(Statement)
		}
	// statements
	case *ExpressionStatement:
		node.Expression = Modify(node.Expression, modifier).(Expression)
	case *InfixExpression:
		node.Left, _ = Modify(node.Left, modifier).(Expression)
		node.Right, _ = Modify(node.Right, modifier).(Expression)
	case *PrefixExpression:
		node.Right, _ = Modify(node.Right, modifier).(Expression)
	case *IndexExpression:
		node.Left, _ = Modify(node.Left, modifier).(Expression)
		node.Index, _ = Modify(node.Index, modifier).(Expression)
	case *IfExpression:
		node.Condition, _ = Modify(node.Condition, modifier).(Expression)
		node.Consequence = Modify(node.Consequence, modifier).(*BlockStatement)
		if node.Alternative != nil {
			node.Alternative = Modify(node.Alternative, modifier).(*BlockStatement)
		}
	case *BlockStatement:
		for i := range node.Statements {
			node.Statements[i], _ = Modify(node.Statements[i], modifier).(Statement)
		}
	case *ReturnStatement:
		node.ReturnValue, _ = Modify(node.ReturnValue, modifier).(Expression)
	case *LetStatement:
		node.Value, _ = Modify(node.Value, modifier).(Expression)
	// literals
	case *FunctionLiteral:
		for i, param := range node.Parameters {
			node.Parameters[i], _ = Modify(param, modifier).(*Identifier)
		}
		node.Body, _ = Modify(node.Body, modifier).(*BlockStatement)
	case *ArrayLiteral:
		for i := range node.Elements {
			node.Elements[i], _ = Modify(node.Elements[i], modifier).(Expression)
		}
	case *HashLiteral:
		newPairs := make(map[Expression]Expression)
		for key, value := range node.Pairs {
			newKey, _ := Modify(key, modifier).(Expression)
			newValue, _ := Modify(value, modifier).(Expression)
			newPairs[newKey] = newValue
		}
		node.Pairs = newPairs
	}

	return modifier(node)
}
