package interpreter

type IContext interface {
	Lookup(string) bool
	Assign(*VariableExp, bool)
}

type Context struct {
	_m map[string]bool
}

func (c *Context) Lookup(name string) bool {
	return c._m[name]
}

func (c *Context) Assign(exp *VariableExp, b bool) {
	if c._m == nil {
		c._m = make(map[string]bool)
	}

	c._m[exp._name] = b
}

func NewContext() *Context {
	return new(Context)
}

type IBooleanExp interface {
	Evaluate(IContext) bool
	Replace(string, IBooleanExp) IBooleanExp
	Copy() IBooleanExp
}

type VariableExp struct {
	_name string
}

func (v *VariableExp) Evaluate(aContext IContext) bool {
	return aContext.Lookup(v._name)
}

func (v *VariableExp) Replace(name string, exp IBooleanExp) IBooleanExp {
	if name == v._name {
		return exp.Copy()
	} else {
		return NewVariableExp(name)
	}
}

func (v *VariableExp) Copy() IBooleanExp {
	return NewVariableExp(v._name)
}

func NewVariableExp(name string) *VariableExp {
	return &VariableExp{
		_name: name,
	}
}

type AndExp struct {
	_operand1 IBooleanExp
	_operand2 IBooleanExp
}

func (e *AndExp) Evaluate(aContext IContext) bool {
	return e._operand1.Evaluate(aContext) && e._operand2.Evaluate(aContext)
}

func (e *AndExp) Replace(name string, exp IBooleanExp) IBooleanExp {
	return NewAndExp(e._operand1.Replace(name, exp), e._operand2.Replace(name, exp))
}

func (e *AndExp) Copy() IBooleanExp {
	return NewAndExp(e._operand1.Copy(), e._operand2.Copy())
}

func NewAndExp(exp1, exp2 IBooleanExp) *AndExp {
	return &AndExp{
		_operand1: exp1,
		_operand2: exp2,
	}
}
