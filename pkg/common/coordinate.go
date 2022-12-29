package common

type Coordinate struct {
	x int
	y int
}

func NewCord(x, y int) Coordinate {
	return Coordinate{
		x: x,
		y: y,
	}
}

func (c *Coordinate) X() int {
	return c.x
}

func (c *Coordinate) Y() int {
	return c.y
}
