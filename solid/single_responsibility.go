package solid

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type (
	circle struct {
		radius float64
	}

	circleAugmented struct {
		circle
		colorValue string
	}

	square struct {
		length float64
	}

	shapeOutput struct {}
)

type (
	targetPrint interface {
		area() float64
		name() string
	}

	shape interface {
		area() float64
		name() string
	}

	shapeAugmented interface {
		shape
		color() string
	}
)

func (ca circleAugmented) color()  string {
	return ca.colorValue
}

func (ca circleAugmented) name()  string {
	return "Circle Augmented"
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) name() string {
	return "Circle"
}

func (s square) area() float64{
	return s.length * s.length
}

func (s square) name() string {
	return "Square"
}

func (o shapeOutput) Text (sh targetPrint) string {
	return fmt.Sprintf("area of the %s is %f", sh.name(), sh.area())
}

func (o shapeOutput) shapeToJSONString(sh shape) string {
	obj := struct {
		Name    string  `json:"shape"`
		Area    float64 `json:"area"`
	}{
		Name: sh.name(),
		Area: sh.area(),
	}

	bs, err :=  json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}

	return string(bs)
}

func (o shapeOutput) shapeAugmentedToJSONString(sh shapeAugmented) string {
	obj := struct {
		Name        string  `json:"shape"`
		Area        float64 `json:"area"`
		Color       string  `json:"color,omitempty"`
	}{
		Name: sh.name(),
		Area: sh.area(),
		Color: sh.color(),
	}

	bs, err :=  json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}

	return string(bs)
}


func singleResponsibility() {

	fmt.Println("Single Responsibility")

	o := &shapeOutput{}


	c := &circle{
		radius: 20,
	}

	fmt.Println(o.shapeToJSONString(*c))

	s := &square{
		length: 10,
	}

	fmt.Println(o.shapeToJSONString(*s))

	ca := &circleAugmented{
		circle : circle{
			radius: 5,
		},
		colorValue: "red",
	}

	fmt.Println(o.shapeAugmentedToJSONString(*ca))

	fmt.Println(o.Text(*c))
	fmt.Println(o.Text(*s))
	fmt.Println(o.Text(*ca))

	fmt.Printf("\n")
}