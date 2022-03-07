package abstract_factory

import (
	"maze/door"
	"maze/maze"
	"maze/room"
	"maze/wall"
	"sync"
)

var (
	singletonNormalMazeFactory        *normalMazeFactory
	singletongBombedMazeFactory       *bombedMazeFactory
	once                              = new(sync.Once)
	lazySingletonEnchantedMazeFactory *enchantedMazeFactory
)

func init() {
	singletonNormalMazeFactory = new(normalMazeFactory)
	singletongBombedMazeFactory = new(bombedMazeFactory)
}

// abstract factory
type IMazeFactory interface {
	MakeMaze() maze.IMaze
	MakeWall() wall.IWall
	MakeRoom(int) room.IRoom
	MakeDoor(r1, r2 room.IRoom) door.IDoor
}

type normalMazeFactory struct{}

func (f *normalMazeFactory) MakeMaze() maze.IMaze {
	return maze.NewNormalMaze()
}

func (f *normalMazeFactory) MakeWall() wall.IWall {
	return wall.NewNormalWall()
}

func (f *normalMazeFactory) MakeRoom(n int) room.IRoom {
	return room.NewNormalRoom(n)
}

func (f *normalMazeFactory) MakeDoor(r1, r2 room.IRoom) door.IDoor {
	return door.NewNormalDoor(r1.(*room.NormalRoom), r2.(*room.NormalRoom))
}

// singleton pattern
func NewNormalMazeFactory() *normalMazeFactory {
	return singletonNormalMazeFactory
}

type enchantedMazeFactory struct{}

func (e *enchantedMazeFactory) MakeMaze() maze.IMaze {
	return maze.NewNormalMaze()
}

func (e *enchantedMazeFactory) MakeWall() wall.IWall {
	return wall.NewNormalWall()
}

func (e *enchantedMazeFactory) MakeRoom(n int) room.IRoom {
	return room.NewEnchantedRoom(n)
}

func (e *enchantedMazeFactory) MakeDoor(r1, r2 room.IRoom) door.IDoor {
	return door.NewDoorNeedingSpell(r1.(*room.EnchantedRoom), r2.(*room.EnchantedRoom))
}

// lazy singleton pattern
func NewEnchantedMazeFactory() *enchantedMazeFactory {
	if lazySingletonEnchantedMazeFactory == nil {
		once.Do(func() {
			lazySingletonEnchantedMazeFactory = new(enchantedMazeFactory)
		})
	}

	return lazySingletonEnchantedMazeFactory
}

type bombedMazeFactory struct{}

func (b *bombedMazeFactory) MakeMaze() maze.IMaze {
	return maze.NewNormalMaze()
}

func (b *bombedMazeFactory) MakeWall() wall.IWall {
	return wall.NewBombedWall()
}

func (b *bombedMazeFactory) MakeRoom(n int) room.IRoom {
	return room.NewRoomWithBomb(n)
}

func (b *bombedMazeFactory) MakeDoor(r1, r2 room.IRoom) door.IDoor {
	return door.NewNormalDoor(r1.(*room.NormalRoom), r2.(*room.NormalRoom))
}

// singleton pattern
func NewBombedMazeFactory() *bombedMazeFactory {
	return singletongBombedMazeFactory
}
