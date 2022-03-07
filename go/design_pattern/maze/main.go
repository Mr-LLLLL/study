package main

import (
	af "maze/abstract-factory"
	b "maze/builder"
	"maze/door"
	fm "maze/factory-method"
	"maze/maze"
	"maze/prototype"
	"maze/room"
	"maze/wall"
)

func main() {
	game := fm.NewMazeGame()

	// abstract factory
	{
		factory := af.NewBombedMazeFactory()
		game.CreateMazeByFactory(factory)

		bombedMazeFactory := af.NewBombedMazeFactory()
		game.CreateMazeByFactory(bombedMazeFactory)

		EnchantedMazeFactory := af.NewEnchantedMazeFactory()
		game.CreateMazeByFactory(EnchantedMazeFactory)
	}

	// builder
	{
		builder := b.NewStandardMazeBuilder()
		game.CreateMazeByBuilder(builder)

		countingMazeBuilder := b.NewCountingMazeBuilder()
		game.CreateMazeByBuilder(countingMazeBuilder)
	}

	// factory method
	{
		game.CreateMaze()

		bombedMazeGame := fm.NewBombedMazeGame()
		bombedMazeGame.CreateMaze()

		enchantedMazeGame := fm.NewEnchantedMazeGame()
		enchantedMazeGame.CreateMaze()
	}

	// prototype
	{
		simpleMazeFacotry := prototype.NewMazePrototypeFactory(new(maze.NormalMaze), new(wall.NormalWall), new(room.NormalRoom), new(door.NormalDoor))
		game.CreateMazeByFactory(simpleMazeFacotry)

		bombedMazeFactory := prototype.NewMazePrototypeFactory(new(maze.NormalMaze), new(wall.BombedWall), new(room.RoomWithBomb), new(door.NormalDoor))
		game.CreateMazeByFactory(bombedMazeFactory)
	}
}

func Use(...interface{}) {}

type sfsf struct {
	i int
}
