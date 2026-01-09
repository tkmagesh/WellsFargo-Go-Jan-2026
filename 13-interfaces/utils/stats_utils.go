package utils

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShape(x ShapeStatsFinder) {
	PrintArea(x)      // x should be AreaFinder
	PrintPerimeter(x) // x should be PerimeterFinder
}
