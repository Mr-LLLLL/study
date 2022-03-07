package equipment

import "composite/base"

type IEquipment interface {
	Name() string

	Power() base.Watt
	NetPrice() base.Currency
	DiscountPrice() base.Currency

	Add(IEquipment)
	Remove(IEquipment)
}

type Equipment struct {
	_name string
}

func (e *Equipment) Name() string {
	return e._name
}

func (e *Equipment) Power() base.Watt {
	panic("not implemented") // TODO: Implement
}

func (e *Equipment) NetPrice() base.Currency {
	panic("not implemented") // TODO: Implement
}

func (e *Equipment) DiscountPrice() base.Currency {
	panic("not implemented") // TODO: Implement
}

func (e *Equipment) Add(_ IEquipment) {
	panic("not implemented") // TODO: Implement
}

func (e *Equipment) Remove(_ IEquipment) {
	panic("not implemented") // TODO: Implement
}

func NewEquipment(name string) *Equipment {
	return &Equipment{
		_name: name,
	}
}

type FloppyDisk struct {
	_name string
}

func (f *FloppyDisk) Name() string {
	panic("not implemented") // TODO: Implement
}

func (f *FloppyDisk) Power() base.Watt {
	panic("not implemented") // TODO: Implement
}

func (f *FloppyDisk) NetPrice() base.Currency {
	panic("not implemented") // TODO: Implement
}

func (f *FloppyDisk) DiscountPrice() base.Currency {
	panic("not implemented") // TODO: Implement
}

func (f *FloppyDisk) Add(_ IEquipment) {
	panic("not implemented") // TODO: Implement
}

func (f *FloppyDisk) Remove(_ IEquipment) {
	panic("not implemented") // TODO: Implement
}

func NewFloppyDisk(name string) *FloppyDisk {
	return &FloppyDisk{
		_name: name,
	}
}

type CompositeEquipment struct {
	_equipment []IEquipment
	_name      string
}

func (c *CompositeEquipment) Name() string {
	panic("not implemented") // TODO: Implement
}

func (c *CompositeEquipment) Power() base.Watt {
	panic("not implemented") // TODO: Implement
}

func (c *CompositeEquipment) NetPrice() base.Currency {
	total := base.Currency(0)
	for _, v := range c._equipment {
		total += v.NetPrice()
	}
	return total
}

func (c *CompositeEquipment) DiscountPrice() base.Currency {
	panic("not implemented") // TODO: Implement
}

func (c *CompositeEquipment) Add(_ IEquipment) {
	panic("not implemented") // TODO: Implement
}

func (c *CompositeEquipment) Remove(_ IEquipment) {
	panic("not implemented") // TODO: Implement
}

func NewCompositeEquipment(name string) *CompositeEquipment {
	return &CompositeEquipment{
		_name: name,
	}
}
