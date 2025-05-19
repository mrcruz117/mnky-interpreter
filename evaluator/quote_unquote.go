package evaluator

import (
	"github.com/mrcruz117/mnky-interpreter/ast"
	"github.com/mrcruz117/mnky-interpreter/object"
)

func quote(node ast.Node) object.Object {
	if node == nil {
		return nil
	}
	return &object.Quote{Node: node}
}
