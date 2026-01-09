package shapes

type Square struct {
	rectangle Rectangle
}

func NewSquare(side float64) *Square {
	return &Square{
		rectangle: Rectangle{
			Height: side,
			Width:  side,
		},
	}
}

func (s Square) Area() float64 {
	return s.rectangle.Area()
}

func (s Square) Perimeter() float64 {
	return s.rectangle.Perimeter()
}
