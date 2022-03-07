package room

import (
	"maze/base"
	"maze/spell"
)

type IRoom interface {
	base.MapSite

	GetSide(base.Direction) base.MapSite
	SetSide(base.Direction, base.MapSite)
}
type NormalRoom struct {
	_sides      [4]base.MapSite
	_roomNumber int
}

func (r *NormalRoom) GetSide(_ base.Direction) base.MapSite {
	panic("not implemented") // TODO: Implement
}

func (r *NormalRoom) SetSide(_ base.Direction, _ base.MapSite) {
	panic("not implemented") // TODO: Implement
}

func (r *NormalRoom) Enter() {
	panic("not implemented") // TODO: Implement
}

func NewNormalRoom(n int) *NormalRoom {
	return &NormalRoom{
		_roomNumber: n,
	}
}

type EnchantedRoom struct {
	NormalRoom
}

func (e *EnchantedRoom) Enter() {
	panic("not implemented") // TODO: Implement
}

func (e *EnchantedRoom) GetSide(_ base.Direction) base.MapSite {
	panic("not implemented") // TODO: Implement
}

func (e *EnchantedRoom) SetSide(_ base.Direction, _ base.MapSite) {
	panic("not implemented") // TODO: Implement
}

func (e *EnchantedRoom) CastSpell() spell.ISpell {
	return spell.NewSpell()
}

func NewEnchantedRoom(n int) *EnchantedRoom {
	return &EnchantedRoom{
		NormalRoom: NormalRoom{
			_roomNumber: n,
		},
	}
}

type RoomWithBomb struct {
	NormalRoom
}

func (r *RoomWithBomb) Enter() {
	panic("not implemented") // TODO: Implement
}

func (r *RoomWithBomb) GetSide(_ base.Direction) base.MapSite {
	panic("not implemented") // TODO: Implement
}

func (r *RoomWithBomb) SetSide(_ base.Direction, _ base.MapSite) {
	panic("not implemented") // TODO: Implement
}

func (r *RoomWithBomb) CastSpell() spell.ISpell {
	return spell.NewSpell()
}

func NewRoomWithBomb(n int) *RoomWithBomb {
	return &RoomWithBomb{
		NormalRoom: NormalRoom{
			_roomNumber: n,
		},
	}
}
