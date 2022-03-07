package door

import (
	"maze/base"
	"maze/room"
)

type IDoor interface {
	base.MapSite

	OtherSideFrom(room.IRoom) room.IRoom
	Clone() IDoor
	Initialize(r1, r2 room.IRoom)
}

type NormalDoor struct {
	_room1  *room.NormalRoom
	_room2  *room.NormalRoom
	_isOpen bool
}

func (d *NormalDoor) Initialize(r1, r2 room.IRoom) {
	d._room1 = r1.(*room.NormalRoom)
	d._room2 = r2.(*room.NormalRoom)
}

func (d *NormalDoor) Enter() {
	panic("not implemented") // TODO: Implement
}

func (d *NormalDoor) Clone() IDoor {
	return &NormalDoor{
		_room1:  d._room1,
		_room2:  d._room2,
		_isOpen: d._isOpen,
	}
}

func (d *NormalDoor) OtherSideFrom(_ room.IRoom) room.IRoom {
	panic("not implemented") // TODO: Implement
}

func NewNormalDoor(r1, r2 *room.NormalRoom) *NormalDoor {
	return &NormalDoor{
		_room1: r1,
		_room2: r2,
	}
}

type DoorNeedingSpell struct {
	_room1  *room.EnchantedRoom
	_room2  room.IRoom
	_isOpen bool
}

func (d *DoorNeedingSpell) Enter() {
	panic("not implemented")
}

func (d *DoorNeedingSpell) Initialize(r1, r2 room.IRoom) {
	d._room1 = r1.(*room.EnchantedRoom)
	d._room2 = r1.(*room.EnchantedRoom)
}

func (d *DoorNeedingSpell) Clone() IDoor {
	return &DoorNeedingSpell{
		_room1:  d._room1,
		_room2:  d._room2,
		_isOpen: d._isOpen,
	}
}

func (d *DoorNeedingSpell) OtherSideFrom(_ room.IRoom) room.IRoom {
	panic("not implemented")
}

func NewDoorNeedingSpell(r1, r2 *room.EnchantedRoom) *DoorNeedingSpell {
	return &DoorNeedingSpell{
		_room1: r1,
		_room2: r2,
	}
}
