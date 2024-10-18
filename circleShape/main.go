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

type Square struct {
	side float64
}
type Cube struct {
	Square
}

type Rectangle struct {
	length, width float64
}

type Cuboid struct {
	Rectangle
	height float64
}

type Traingle struct {
	base, height float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func NewTriangle(base, height float64) *Traingle {
	return &Traingle{base: base, height: height}
}

func NewSquare(side float64) *Square {
	return &Square{side: side}
}

func NewRectangle(length, width float64) *Rectangle {
	return &Rectangle{length: length, width: width}
}

func NewCube(side float64) *Cube {
	return &Cube{
		Square: *NewSquare(side),
	}
}

func NewCuboid(length, width, height float64) *Cuboid {
	return &Cuboid{
		Rectangle: *NewRectangle(length, width),
		height:    height,
	}
}

func (t *Traingle) Area() float64 {
	return (t.base * t.height) / 2
}

func (t *Traingle) Volume() float64 {
	return 0
}

func (t *Traingle) Name() string {
	return "Triangle"
}
func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Volume() float64 {
	return 0
}

func (r *Rectangle) Name() string {
	return "Rectangle"
}

func (sq *Square) Area() float64 {
	return math.Pow(sq.side, 2)
}

func (sq *Square) Volume() float64 {
	return 0
}

func (sq *Square) Name() string {
	return "Square"
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

func (cz *Cube) Area() float64 {
	return (6.0) * cz.Square.Area()
}

func (cx *Cube) Volume() float64 {
	return math.Pow(cx.Square.Area(), 3)
}

func (cb *Cube) Name() string {
	return "Cube"
}

func (cy *Cuboid) Area() float64 {
	return 2 * (cy.Rectangle.Area() + cy.length*cy.height + cy.height*cy.width)
}

func (cy *Cuboid) Volume() float64 {
	return cy.Rectangle.Area() * cy.height
}

func (cy *Cuboid) Name() string {
	return "Cuboid"
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

type Hemisphere struct {
	Circle
}

func NewHemisphere(radius float64) *Sphere {
	return &Sphere{
		Circle: *NewCircle(radius),
	}
}

func (s *Hemisphere) Area() float64 {
	return 3 * s.Circle.Area()
}

func (s *Hemisphere) Volume() float64 {
	return (2.0 / 3.0) * s.Circle.Area() * s.radius
}

func (s *Hemisphere) Name() string {
	return "Hemisphere"
}

func printShapeInfo(s Shape) {
	fmt.Printf("Shape: %s\n", s.Name())
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Volume: %.2f\n\n", s.Volume())
}

func main() {
	// fmt.Println("Happy Coding @Gophers!")

	circle := NewCircle(15)
	traingle := NewTriangle(15, 6)
	square := NewSquare(8)
	rectangle := NewRectangle(12, 8)
	cylinder := NewCylinder(12, 10)
	cone := NewCone(12, 15)
	sphere := NewSphere(5)
	hemispere := NewHemisphere(6)
	cube := NewCube(6)
	cuboid := NewCuboid(12, 15, 10)

	shapes := []Shape{circle, traingle, square, rectangle, cylinder, cone, sphere, hemispere, cube, cuboid}

	for _, shape := range shapes {
		printShapeInfo(shape)
	}
}
