package base

type Direction int8

const (
	North Direction = iota
	South
	East
	West
)

type MapSite interface {
	Enter()
}

