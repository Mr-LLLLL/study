package parser

import (
	program_node "facade/program-node"
	"facade/scanner"
)

type IParser interface {
	Parse(scanner.IScanner, *program_node.ProgramNodeBuilder)
}

type Parser struct{}

func (p *Parser) Parse(scanner scanner.IScanner, builder *program_node.ProgramNodeBuilder) {

}

func NewParser() *Parser {
	return new(Parser)
}
