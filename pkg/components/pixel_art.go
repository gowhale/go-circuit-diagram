package components

const (
	pixelEmpt = 0
	pixelFill = 1
)

var ledPixels = [][]int{
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill},
	{pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelFill, pixelFill},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelFill},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
}

var shiftPixels = [][]int{
	{pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt},
}
