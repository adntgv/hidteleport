package types

type Screen struct {
	Width  int
	Height int
}

func (s *Screen) ComputePositionAt(x, y float64) (int, int) {
	width := float64(s.Width) * x
	height := float64(s.Height) * y
	return int(width), int(height)
}
