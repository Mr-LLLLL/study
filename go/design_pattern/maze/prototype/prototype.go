package prototype

import (
	"maze/door"
	"maze/maze"
	"maze/room"
	"maze/wall"
)

type MazePrototypeFactory struct {
	_prototypeMaze maze.IMaze
	_prototypeRoom room.IRoom
	_prototypeWall wall.IWall
	_prototypeDoor door.IDoor
}

func (p *MazePrototypeFactory) MakeMaze() maze.IMaze {
	return maze.NewNormalMaze()
}

func (p *MazePrototypeFactory) MakeWall() wall.IWall {
	return p._prototypeWall.Clone()
}

func (p *MazePrototypeFactory) MakeRoom(n int) room.IRoom {
	return room.NewNormalRoom(n)
}

func (p *MazePrototypeFactory) MakeDoor(r1, r2 room.IRoom) door.IDoor {
	door := p._prototypeDoor.Clone()
	door.Initialize(r1, r2)
	return door
}

func NewMazePrototypeFactory(m maze.IMaze, w wall.IWall, r room.IRoom, d door.IDoor) *MazePrototypeFactory {
	return &MazePrototypeFactory{
		_prototypeMaze: m,
		_prototypeWall: w,
		_prototypeRoom: r,
		_prototypeDoor: d,
	}
}
