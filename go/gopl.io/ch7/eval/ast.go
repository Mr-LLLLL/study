package eval

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

type Var string

type literal float64

type Env map[Var]float64

type call struct {
	fn   string
	args []Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type unary struct {
	op rune
	x  Expr
}
