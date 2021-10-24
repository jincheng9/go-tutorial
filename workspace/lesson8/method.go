package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) changeRadius(radius float64) {
	c.radius = radius
}

func (c *Circle) changeRadius2(radius float64) {
	c.radius = radius
}

func (c Circle) addRadius(x float64) float64{
	return c.radius + x
}

func main() {
	var c Circle
	c.radius = 10
	fmt.Println("radius=", c.radius, "area=", c.getArea())		

	c.changeRadius(20)
	fmt.Println("radius=", c.radius, "area=", c.getArea())		

	c.changeRadius2(20)
	fmt.Println("radius=", c.radius, "area=", c.getArea())	

	result := c.addRadius(3.6)
	fmt.Println("radius=", c.radius, "result=", result)
}