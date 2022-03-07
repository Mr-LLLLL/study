package builder

import (
	"maze/base"
	"maze/door"
	"maze/maze"
	"maze/room"
	"maze/wall"
)

// builder
type IMazeBuilder interface {
	BuildMaze()
	BuildRoom(int)
	BuildDoor(roomFrom, roomTo int)
	GetMaze() maze.IMaze
}

type StandardMazeBuilder struct {
	_currentMaze maze.IMaze
}

func (b *StandardMazeBuilder) BuildMaze() {
	b._currentMaze = new(maze.NormalMaze)
}

func (b *StandardMazeBuilder) BuildRoom(n int) {
	if b._currentMaze.RoomNo(n) == nil {
		room := new(room.NormalRoom)
		b._currentMaze.AddRoom(room)

		room.SetSide(base.North, new(wall.NormalWall))
		room.SetSide(base.South, new(wall.NormalWall))
		room.SetSide(base.East, new(wall.NormalWall))
		room.SetSide(base.West, new(wall.NormalWall))
	}
}

func (b *StandardMazeBuilder) BuildDoor(roomFrom, roomTo int) {
	r1 := b._currentMaze.RoomNo(roomFrom)
	r2 := b._currentMaze.RoomNo(roomTo)
	d := door.NewNormalDoor(r1.(*room.NormalRoom), r2.(*room.NormalRoom))

	r1.SetSide(b.commonWall(r1, r2), d)
	r2.SetSide(b.commonWall(r1, r2), d)
}

func (b *StandardMazeBuilder) GetMaze() maze.IMaze {
	return b._currentMaze
}

func (b *StandardMazeBuilder) commonWall(r1, r2 room.IRoom) base.Direction {
	// TODO:no implement
	return base.North
}

func NewStandardMazeBuilder() *StandardMazeBuilder {
	return &StandardMazeBuilder{
		_currentMaze: nil,
	}
}

type CountingMazeBuilder struct {
	_doors int
	_rooms int
}

func (b *CountingMazeBuilder) BuildMaze() {}

func (b *CountingMazeBuilder) BuildRoom(n int) {
	b._rooms++
}

func (b *CountingMazeBuilder) BuildDoor(roomFrom, roomTo int) {
	b._doors++
}

func (b *CountingMazeBuilder) GetMaze() maze.IMaze {
	return nil
}

func (b *CountingMazeBuilder) GetCounts(rooms, doors *int) {
	*rooms = b._rooms
	*doors = b._doors
}

func NewCountingMazeBuilder() *CountingMazeBuilder {
	return &CountingMazeBuilder{
		_rooms: 0,
		_doors: 0,
	}
}
