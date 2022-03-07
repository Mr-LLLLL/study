package memento

import "memento/base"

type MoveCommand struct {
	_state  *ConstraintSolverMemento
	_delta  base.Point
	_target *base.Graphic
}

func (m *MoveCommand) Execute() {
	solver := NewConstraintSolver()
	m._state = solver.CreateMemento()
	m._target.Move(m._delta)
	solver.Solve()
}

func (m *MoveCommand) Unexecute() {
	solver := NewConstraintSolver()
	m._target.Back(m._delta)
	solver.SetMemento(m._state)
	solver.Solve()
}

type ConstraintSolver struct{}

func (c *ConstraintSolver) Solve() {}

func (c *ConstraintSolver) AddConstraint(startConnection, endConnection *base.Graphic) {}

func (c *ConstraintSolver) RemoveConstraint(startConnection, endConnection *base.Graphic) {}

func (c *ConstraintSolver) SetMemento(*ConstraintSolverMemento) {}

func (c *ConstraintSolver) CreateMemento() *ConstraintSolverMemento {
	return new(ConstraintSolverMemento)
}

var solver *ConstraintSolver

func NewConstraintSolver() *ConstraintSolver {
	if solver == nil {
		solver = new(ConstraintSolver)
	}
	return solver
}

type ConstraintSolverMemento struct{}
