package evaluator

import (
	"testing"

	"github.com/mrcruz117/mnky-interpreter/ast"
	"github.com/mrcruz117/mnky-interpreter/lexer"
	"github.com/mrcruz117/mnky-interpreter/object"
	"github.com/mrcruz117/mnky-interpreter/parser"
)

func TestDefineMacros(t *testing.T) {
	input := `
	let number = 1;
	let function = fn(x, y) { x + y };
	let mymacro = macro(x, y) { x + y; };
	`

	env := object.NewEnvironment()
	program := testParseProgram(input)

	DefineMacros(program, env)

	if len(program.Statements) != 2 {
		t.Fatalf("Wrong number of statements. got=%d", len(program.Statements))
	}

	_, ok := env.Get("number")
	if ok {
		t.Fatalf("number should not be defined")
	}

	_, ok = env.Get("function")
	if ok {
		t.Fatalf("function should not be defined")
	}

	obj, ok := env.Get("mymacro")
	if !ok {
		t.Fatalf("macro not in env")
	}

	macro, ok := obj.(*object.Macro)
	if !ok {
		t.Fatalf("object is not a macro. got=%T", obj)
	}

	if len(macro.Parameters) != 2 {
		t.Fatalf("wrong number of parameters. got=%d", len(macro.Parameters))
	}
	if macro.Parameters[0].String() != "x" {
		t.Fatalf("wrong parameter. got=%s", macro.Parameters[0])
	}
	if macro.Parameters[1].String() != "y" {
		t.Fatalf("wrong parameter. got=%s", macro.Parameters[1])
	}

	expectedBody := "(x + y)"
	if macro.Body.String() != expectedBody {
		t.Fatalf("wrong body. got=%s", macro.Body.String())
	}
}

func testParseProgram(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
