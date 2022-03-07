package state

import "fmt"

type Machine struct {
	state IState
}

func (m *Machine) SetState(state IState) {
	m.state = state
}

func (m *Machine) GetStateName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

type IState interface {
	GetName() string
	Approval(*Machine)
	Reject(*Machine)
}

type leaderApproveState struct{}

func (leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader Approval")
	m.SetState(GetFinanceApproveState())
}

func (leaderApproveState) GetName() string {
	return "LeaderApproveState"
}

func (leaderApproveState) Reject(m *Machine) {}

func GetLeaderApproveState() *leaderApproveState {
	return new(leaderApproveState)
}

type financeApproveState struct{}

func (f *financeApproveState) Approval(m *Machine) {
	fmt.Println("finacial approved")
	fmt.Println("send money succeed")
}

func (f *financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

func (f *financeApproveState) GetName() string {
	return "FinanceApproveState"
}

func GetFinanceApproveState() *financeApproveState {
	return new(financeApproveState)
}
