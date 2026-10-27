package main

import "testing"
import "math"

func TestShapePerimeter(t *testing.T){
	
	t.Run("Shape tests", func(t *testing.T)  {
		perimeterTests := []struct{
			shape Shape
			want float64
		}{
			{Rectangle{15.0,20.0},70.0},
			{Circle{10.0},2 * 10.0 * math.Pi},
		}

		for _,tt := range perimeterTests {
			assertFloatEquals(t,tt.shape.Perimeter(),tt.want)
		}
	})
	t.Run("Rectangle perimeter", func(t *testing.T)  {
		rectangle := Rectangle{15.0,20.0}
		got := rectangle.Perimeter()
		want := 70.0
		assertFloatEquals(t,got,want)
	})

	t.Run("circle perimeter", func(t *testing.T)  {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := 2 * 10.0 * math.Pi

		assertFloatEquals(t, got, want)
	})
}

func assertFloatEquals(t testing.TB, got float64, want float64){
	t.Helper()

	if got != want {
		t.Errorf("got %f want %f",got, want)
	}
}