package program_node

import code_generator "facade/code-generator"

type ProgramNodeBuilder struct {
	_node code_generator.IProgramNode
}

func (p *ProgramNodeBuilder) NewVariable(name string) code_generator.IProgramNode {
	return nil
}

func (p *ProgramNodeBuilder) NewAssignment(variable, expression code_generator.IProgramNode) code_generator.IProgramNode {
	return nil
}

func (p *ProgramNodeBuilder) NewReturnStatement(value code_generator.IProgramNode) code_generator.IProgramNode {
	return nil
}

func (p *ProgramNodeBuilder) NewCondition(condition, truePart, falsePart code_generator.IProgramNode) code_generator.IProgramNode {
	return nil
}

func (p *ProgramNodeBuilder) GetRootNode() code_generator.IProgramNode {
	return p._node
}

func NewProgramNodeBuilder() *ProgramNodeBuilder {
	return new(ProgramNodeBuilder)
}
