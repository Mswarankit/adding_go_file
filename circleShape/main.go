package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Name() string
	Area() float64
	Volume() float64
}

type Circle struct {
	radius float64
}

type Cylinder struct {
	Circle
	height float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) Volume() float64 {
	return 0
}

func (c *Circle) Name() string {
	return "Circle"
}

func NewCylinder(radius, height float64) *Cylinder {
	return &Cylinder{
		Circle: Circle{radius: radius},
		height: height,
	}
}

func (cy *Cylinder) Area() float64 {
	return 2*cy.Circle.Area() + 2*math.Pi*cy.radius*cy.height
}

func (cy *Cylinder) Volume() float64 {
	return cy.Circle.Area() * cy.height
}

func (cy *Cylinder) Name() string {
	return "Cylinder"
}

type Cone struct {
	Cylinder
}

func NewCone(radius, height float64) *Cone {
	return &Cone{
		Cylinder: *NewCylinder(radius, height),
	}
}

func (cone *Cone) Area() float64 {
	t := math.Pow
	lateral := t((t(cone.radius, 2) + t(cone.height, 2)), 0.5)
	return math.Pi*t(cone.radius, 2) + math.Pi*cone.radius*lateral
}

func (cone *Cone) Name() string {
	return "Cone"
}

func (cone *Cone) Volume() float64 {
	return (1.0 / 3.0) * cone.Cylinder.Volume()
}

type Sphere struct {
	Circle
}

func NewSphere(radius float64) *Sphere {
	return &Sphere{
		Circle: *NewCircle(radius),
	}
}

func (s *Sphere) Area() float64 {
	return 4 * s.Circle.Area()
}

func (s *Sphere) Volume() float64 {
	return (4.0 / 3.0) * s.Circle.Area() * s.radius
}

func (s *Sphere) Name() string {
	return "Sphere"
}

func printShapeInfo(s Shape) {
	fmt.Printf("Shape: %s\n", s.Name())
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Volume: %.2f\n\n", s.Volume())
}

func main() {
	// fmt.Println("Happy Coding @Gophers!")

	c := NewCircle(15)
	cy := NewCylinder(12, 10)
	cone := NewCone(12, 15)
	s := NewSphere(5)

	shapes := []Shape{c, cy, cone, s}

	for _, shape := range shapes {
		printShapeInfo(shape)
	}
}
