package common

// Coordinate represents a point on board x,y
type Coordinate struct {
	X int
	Y int
}

// NewCord returns a coordinate struct
func NewCord(x, y int) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
	}
}

// GetX gets the x part of the coord
func (c *Coordinate) GetX() int {
	return c.X
}

// GetY gets the y part of the coord
func (c *Coordinate) GetY() int {
	return c.Y
}
