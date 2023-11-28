package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	want := 100.0
	if got != want {
		t.Errorf("got %2.f want %2.f", got, want)
	}

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := Area(circle)
		want := 3.14
		if got != want {
			t.Errorf("got %2.g want %2.g", got, want)
		}
	})
}
