package memento

type InputText struct {
	content string
}

func (in *InputText) Append(content string) {
	in.content += content
}

func (in *InputText) GetText() string {
	return in.content
}

func (in *InputText) Snapshot() *Snapshot {
	return &Snapshot{
		content: in.content,
	}
}

func (in *InputText) Restore(s *Snapshot) {
	in.content = s.GetText()
}

type Snapshot struct {
	content string
}

func (s *Snapshot) GetText() string {
	return s.content
}
