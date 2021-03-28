package solid
type (
	calculator struct {

	}
)

func (c calculator) areaSum(shapes ...shape) float64 {
	var sum float64
	for _, shape := range shapes{
		sum += shape.area()
	}
	return sum
}