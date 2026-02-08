package main

import (
	"testing"

)

//tolerance for floating point comparisons
const tolerance = 0.0001

func almostEqual(a, b float64) bool {
	return (a-b) < tolerance && (b-a) < tolerance
}


// Rectangle Tests


func TestRectangleArea(t *testing.T) {
	r := Rectangle{Width: 4, Height: 5}
	expected := 20.0

	if !almostEqual(r.Area(), expected) {
		t.Errorf("expected %.2f, got %.2f", expected, r.Area())
	}
}

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{Width: 4, Height: 5}
	expected := 18.0

	if !almostEqual(r.Perimeter(), expected) {
		t.Errorf("expected %.2f, got %.2f", expected, r.Perimeter())
	}
}

func TestRectangleScale(t *testing.T) {
	r := Rectangle{Width: 2, Height: 3}
	r.Scale(2)

	if r.Width != 4 || r.Height != 6 {
		t.Errorf("expected width=4 height=6, got width=%.2f height=%.2f", r.Width, r.Height)
	}
}


// Circle Tests

func TestCircleArea(t *testing.T) {
	c := Circle{Radius: 2}
	expected := 12.56637

	if !almostEqual(c.Area(), expected) {
		t.Errorf("expected %.5f, got %.5f", expected, c.Area())
	}
}

func TestCirclePerimeter(t *testing.T) {
	c := Circle{Radius: 2}
	expected := 12.56637

	if !almostEqual(c.Perimeter(), expected) {
		t.Errorf("expected %.5f, got %.5f", expected, c.Perimeter())
	}
}

func TestCircleScale(t *testing.T) {
	c := Circle{Radius: 3}
	c.Scale(3)

	if c.Radius != 9 {
		t.Errorf("expected radius=9, got %.2f", c.Radius)
	}
}


// Triangle Tests

func TestTriangleArea(t *testing.T) {
	tr := Triangle{Base: 6, Height: 4}
	expected := 12.0

	if !almostEqual(tr.Area(), expected) {
		t.Errorf("expected %.2f, got %.2f", expected, tr.Area())
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tr := Triangle{Base: 6}
	expected := 18.0

	if !almostEqual(tr.Perimeter(), expected) {
		t.Errorf("expected %.2f, got %.2f", expected, tr.Perimeter())
	}
}

func TestTriangleScale(t *testing.T) {
	tr := Triangle{Base: 3, Height: 4}
	tr.Scale(2)

	if tr.Base != 6 || tr.Height != 8 {
		t.Errorf("expected base=6 height=8, got base=%.2f height=%.2f", tr.Base, tr.Height)
	}
}
