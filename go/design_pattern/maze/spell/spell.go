package spell

type ISpell interface{}

type Spell struct{}

func NewSpell() *Spell {
	return new(Spell)
}
