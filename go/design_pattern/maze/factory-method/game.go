package factory_method

import (
	mf "maze/abstract-factory"
	"maze/base"
	b "maze/builder"
	"maze/door"
	"maze/maze"
	"maze/room"
	"maze/wall"
	"sync"
)

var (
	lazySingletonMazeGame       *mazeGame
	lazySingletonBombedMazeGame *bombedMazeGame
	singletongEnchantedMazeGame *enchantedMazeGame
	once                        = &sync.Once{}
	bombedOnce                  = &sync.Once{}
)

func init() {
	singletongEnchantedMazeGame = new(enchantedMazeGame)
}

// factory method
type IMazeGame interface {
	CreateMaze() maze.IMaze
	MakeMaze() maze.IMaze
	MakeRoom(int) room.IRoom
	MakeWall() wall.IWall
	MakeDoor(r1, r2 room.IRoom) door.IDoor
}

type mazeGame struct{}

func (m *mazeGame) MakeMaze() maze.IMaze {
	return maze.NewNormalMaze()
}

func (m *mazeGame) MakeWall() wall.IWall {
	return wall.NewNormalWall()
}

func (m *mazeGame) MakeRoom(n int) room.IRoom {
	return room.NewNormalRoom(n)
}

func (m *mazeGame) MakeDoor(r1, r2 room.IRoom) door.IDoor {
	return door.NewNormalDoor(r1.(*room.NormalRoom), r2.(*room.NormalRoom))
}

func (m *mazeGame) CreateMaze() maze.IMaze {
	aMaze := m.MakeMaze()

	r1 := m.MakeRoom(1)
	r2 := m.MakeRoom(2)
	theDoor := m.MakeDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(base.North, m.MakeWall())
	r1.SetSide(base.East, theDoor)
	r1.SetSide(base.South, m.MakeWall())
	r1.SetSide(base.West, m.MakeWall())

	r2.SetSide(base.North, m.MakeWall())
	r2.SetSide(base.East, m.MakeWall())
	r2.SetSide(base.South, m.MakeWall())
	r2.SetSide(base.West, theDoor)

	return aMaze
}

func (m *mazeGame) CreateMazeByFactory(factory mf.IMazeFactory) maze.IMaze {
	aMaze := factory.MakeMaze()
	r1 := factory.MakeRoom(1)
	r2 := factory.MakeRoom(2)
	aDoor := factory.MakeDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(base.North, factory.MakeWall())
	r1.SetSide(base.East, aDoor)
	r1.SetSide(base.South, factory.MakeWall())
	r1.SetSide(base.West, factory.MakeWall())

	r2.SetSide(base.North, factory.MakeWall())
	r2.SetSide(base.East, factory.MakeWall())
	r2.SetSide(base.South, factory.MakeWall())
	r2.SetSide(base.West, aDoor)

	return aMaze
}

func (m *mazeGame) CreateMazeByBuilder(builder b.IMazeBuilder) maze.IMaze {
	builder.BuildMaze()

	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}

// lazy singleton pattern
func NewMazeGame() *mazeGame {
	if lazySingletonMazeGame == nil {
		once.Do(func() {
			lazySingletonMazeGame = new(mazeGame)
		})
	}

	return lazySingletonMazeGame
}

type bombedMazeGame struct{}

func (m *bombedMazeGame) MakeMaze() maze.IMaze {
	return maze.NewNormalMaze()
}

func (m *bombedMazeGame) MakeWall() wall.IWall {
	return wall.NewBombedWall()
}

func (m *bombedMazeGame) MakeRoom(n int) room.IRoom {
	return room.NewRoomWithBomb(n)
}

func (m *bombedMazeGame) MakeDoor(r1, r2 room.IRoom) door.IDoor {
	return door.NewNormalDoor(r1.(*room.NormalRoom), r2.(*room.NormalRoom))
}

func (m *bombedMazeGame) CreateMaze() maze.IMaze {
	aMaze := m.MakeMaze()

	r1 := m.MakeRoom(1)
	r2 := m.MakeRoom(2)
	theDoor := m.MakeDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(base.North, m.MakeWall())
	r1.SetSide(base.East, theDoor)
	r1.SetSide(base.South, m.MakeWall())
	r1.SetSide(base.West, m.MakeWall())

	r2.SetSide(base.North, m.MakeWall())
	r2.SetSide(base.East, m.MakeWall())
	r2.SetSide(base.South, m.MakeWall())
	r2.SetSide(base.West, theDoor)

	return aMaze
}

// lazy singleton patter
func NewBombedMazeGame() *bombedMazeGame {
	if lazySingletonBombedMazeGame == nil {
		bombedOnce.Do(func() {
			lazySingletonBombedMazeGame = new(bombedMazeGame)
		})
	}
	return lazySingletonBombedMazeGame
}

type enchantedMazeGame struct{}

func (m *enchantedMazeGame) MakeMaze() maze.IMaze {
	return maze.NewNormalMaze()
}

func (m *enchantedMazeGame) MakeWall() wall.IWall {
	return wall.NewBombedWall()
}

func (m *enchantedMazeGame) MakeRoom(n int) room.IRoom {
	return room.NewEnchantedRoom(n)
}

func (m *enchantedMazeGame) MakeDoor(r1, r2 room.IRoom) door.IDoor {
	return door.NewDoorNeedingSpell(r1.(*room.EnchantedRoom), r2.(*room.EnchantedRoom))
}

func (m *enchantedMazeGame) CreateMaze() maze.IMaze {
	aMaze := m.MakeMaze()

	r1 := m.MakeRoom(1)
	r2 := m.MakeRoom(2)
	theDoor := m.MakeDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(base.North, m.MakeWall())
	r1.SetSide(base.East, theDoor)
	r1.SetSide(base.South, m.MakeWall())
	r1.SetSide(base.West, m.MakeWall())

	r2.SetSide(base.North, m.MakeWall())
	r2.SetSide(base.East, m.MakeWall())
	r2.SetSide(base.South, m.MakeWall())
	r2.SetSide(base.West, theDoor)

	return aMaze
}

// singleton pattern
func NewEnchantedMazeGame() *enchantedMazeGame {
	return singletongEnchantedMazeGame
}
