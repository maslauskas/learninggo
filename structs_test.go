package main

import (
    "fmt"
    "testing"
)

func TestPerimeter(t *testing.T) {
    want := 40.0
    have := Perimeter(Rectangle{10.0, 10.0})

    if want != have {
        t.Errorf("expected %.2f, have %.2f", want, have)
    }
}

func TestArea(t *testing.T) {
    areatests := []struct {
        shape Shape
        want float64
    }{
        {shape: Rectangle{Width: 10.0, Height: 10.0}, want: 100.0},
        {Circle{10.0}, 314.1592653589793},
        {Triangle{12.0, 6.0}, 36.0},
    }

    for i, tt := range areatests {
        t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
            have := tt.shape.Area()

            if tt.want != have {
                t.Errorf("expected %.2f, have %.2f", tt.want, have)
            }
        })
    }
}
