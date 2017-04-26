package evaluator

import (
	"github.com/jamesroutley/monkey/lexer"
	"github.com/jamesroutley/monkey/object"
	"github.com/jamesroutley/monkey/parser"
	"log"
	"io/ioutil"
	"testing"
)

func init() {
	// Silence logs when testing
	log.SetOutput(ioutil.Discard)
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct{
		input string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not an Integer. got %T", obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value, expected: %d, got: %d", expected,
			result.Value)
		return false
	}
	return true
}
