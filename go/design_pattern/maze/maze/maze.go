package maze

import (
	"maze/room"
)

type IMaze interface {
	AddRoom(room.IRoom)
	RoomNo(int) room.IRoom
}

type NormalMaze struct{}

func (m *NormalMaze) AddRoom(_ room.IRoom) {
	panic("not implemented") // TODO: Implement
}

func (m *NormalMaze) RoomNo(_ int) room.IRoom {
	panic("not implemented") // TODO: Implement
}

func NewNormalMaze() *NormalMaze {
	return new(NormalMaze)
}
