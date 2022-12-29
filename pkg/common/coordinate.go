package common

// Coordinate represents a point on board x,y
type Coordinate struct {
	x int
	y int
}

// NewCord returns a coordinate struct
func NewCord(x, y int) Coordinate {
	return Coordinate{
		x: x,
		y: y,
	}
}

// X gets the x part of the coord
func (c *Coordinate) X() int {
	return c.x
}

// Y gets the y part of the coord
func (c *Coordinate) Y() int {
	return c.y
}
