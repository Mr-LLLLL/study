package flyweight

var units = map[int]*ChessPieceUnit{
	1: {
		Id:    1,
		Name:  "车",
		Color: "red",
	},
	2: {
		Id:    2,
		Name:  "马",
		Color: "red",
	},
	3: {
		Id:    3,
		Name:  "炮",
		Color: "red",
	},
	// ....
}

type ChessPieceUnit struct {
	Id    uint
	Name  string
	Color string
}

func NewChessPieceUni(id int) *ChessPieceUnit {
	return units[id]
}

type ChessPiece struct {
	Unit *ChessPieceUnit
	X, Y int
}

type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

func NewChessBoard() *ChessBoard {
	board := &ChessBoard{
		chessPieces: map[int]*ChessPiece{},
	}

	for id := range units {
		board.chessPieces[id] = &ChessPiece{
			Unit: NewChessPieceUni(id),
			X:    0,
			Y:    0,
		}
	}
	return board
}

func (c *ChessBoard) Move(id, x, y int) {
	c.chessPieces[id].X = x
	c.chessPieces[id].Y = y
}
