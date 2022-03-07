package facade

import (
	"facade/base"
	code_generator "facade/code-generator"
	"facade/parser"
	program_node "facade/program-node"
	"facade/scanner"
)

type compiler struct {
}

func (c *compiler) Compile(input *base.IsStream, output *base.BytecodeStream) {
	scanner := scanner.NewScanner(input)
	builder := program_node.NewProgramNodeBuilder()
	parser := parser.NewParser()

	parser.Parse(scanner, builder)

	generator := code_generator.NewRISCCodeGenerator(output)
	parseTree := builder.GetRootNode()
	parseTree.Traverse(generator)

}
