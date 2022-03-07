package interpreter

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	x := NewVariableExp("X")
	y := NewVariableExp("Y")

	expression := NewAndExp(x, y)

	ctx := NewContext()
	ctx.Assign(x, true)
	ctx.Assign(y, true)

	res := expression.Evaluate(ctx)
	fmt.Println(res)
}
