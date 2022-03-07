package code_generator

import (
	"facade/base"
)

type ICodeGenerator interface {
	VisitStatement(IProgramNode)
	VisitExpression(IProgramNode)
}

type RISCCodeGenerator struct {
	_output *base.BytecodeStream
}

func (c *RISCCodeGenerator) VisitStatement(IProgramNode) {

}

func (c *RISCCodeGenerator) VisitExpression(IProgramNode) {

}

func NewRISCCodeGenerator(output *base.BytecodeStream) *RISCCodeGenerator {
	return &RISCCodeGenerator{
		_output: output,
	}
}

type IProgramNode interface {
	Add(IProgramNode)
	Remove(IProgramNode)
	Traverse(ICodeGenerator)
}

type StatementNode struct {
	_nextnode IProgramNode
}

func (p *StatementNode) Add(IProgramNode) {

}

func (p *StatementNode) Remove(IProgramNode) {
}

func (p *StatementNode) Traverse(ICodeGenerator) {}

type ExpressionNode struct {
	_nextnode IProgramNode
}

func (p *ExpressionNode) Add(IProgramNode) {

}

func (p *ExpressionNode) Remove(IProgramNode) {
}

func (p *ExpressionNode) Traverse(cg ICodeGenerator) {
	cg.VisitExpression(p)
	node := p._nextnode
	for node != nil {
		cg.VisitExpression(node)
		switch v := node.(type) {
		case *ExpressionNode:
			cg.VisitExpression(v)
			node = v._nextnode
		case *StatementNode:
			cg.VisitStatement(v)
			node = v._nextnode
		}
	}
}
