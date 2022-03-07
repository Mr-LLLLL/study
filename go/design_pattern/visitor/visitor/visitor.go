package visitor

import "visitor/base"

type IEquipment interface {
	Accept(IEquipmentVisitor)
	NetPrice() base.Currency
	DiscountPrice() base.Currency
}

type FloppyDisk struct {
	_name string
}

func (f *FloppyDisk) Accept(v IEquipmentVisitor) {
	v.VisitFloppyDisk(f)
}

func (f *FloppyDisk) NetPrice() base.Currency {
	return 0
}

func (f *FloppyDisk) DiscountPrice() base.Currency {
	return 0
}

type Chassis struct {
	_parts []IEquipment
	_name  string
}

func (c *Chassis) Accept(visitor IEquipmentVisitor) {
	for _, v := range c._parts {
		v.Accept(visitor)
	}
	visitor.VisitChassis(c)
}

func (f *Chassis) NetPrice() base.Currency {
	return 0
}

func (f *Chassis) DiscountPrice() base.Currency {
	return 0
}

type IEquipmentVisitor interface {
	VisitFloppyDisk(*FloppyDisk)
	VisitChassis(*Chassis)
}

type PricingVisitor struct {
	_total base.Currency
}

func (p *PricingVisitor) VisitFloppyDisk(f *FloppyDisk) {
	p._total += f.NetPrice()
}

func (p *PricingVisitor) VisitChassis(c *Chassis) {
	p._total += c.DiscountPrice()
}
