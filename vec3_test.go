package vec3

import (
	"math"
	"testing"
)

func TestToString(t *testing.T) {
	v := Vec3{}
	got := v.ToString()
	want := "Vec3 struct: 0.00, 0.00, 0.00"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	check := func(v1, v2 Vec3) {
		if v1 != v2 {
			t.Errorf("got %v want %v", v1, v2)
		}
	}

	t.Run("add Vec3", func(t *testing.T) {
		got := Vec3{2, 4, 5}
		got.Add(Vec3{X: 3, Y: 1})
		want := Vec3{5, 5, 5}

		check(got, want)
	})

	t.Run("add float64", func(t *testing.T) {
		got := Vec3{3, 6, 7}
		got.Add(4, 2, 1)
		want := Vec3{7, 8, 8}

		check(got, want)
	})
}

func TestSub(t *testing.T) {
	check := func(v1, v2 Vec3) {
		if v1 != v2 {
			t.Errorf("got %v want %v", v1, v2)
		}
	}

	t.Run("sub Vec3", func(t *testing.T) {
		got := Vec3{2, 4, 5}
		got.Sub(Vec3{1, 3, 4})
		want := Vec3{1, 1, 1}

		check(got, want)
	})

	t.Run("sub float64", func(t *testing.T) {
		got := Vec3{3, 6, 7}
		got.Sub(2, 5, 6)
		want := Vec3{1, 1, 1}

		check(got, want)
	})
}

func TestMult(t *testing.T) {
	got := Vec3{2, 5, 8}
	got.Mult(3)
	want := Vec3{6, 15, 24}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDiv(t *testing.T) {
	got := Vec3{4, 5, 8}
	check := func(v1, v2 Vec3) {
		if v1 != v2 {
			t.Errorf("got %v want %v", v1, v2)
		}
	}

	t.Run("div by 0", func(t *testing.T) {
		got.Div(0)
		want := Vec3{4, 5, 8}

		check(got, want)
	})

	t.Run("div", func(t *testing.T) {
		got.Div(2)
		want := Vec3{2, 2.5, 4}

		check(got, want)
	})
}

func TestMagnitude(t *testing.T) {
	v := Vec3{3, 4, 9}
	mag := v.Magnitude()
	got := math.Round(mag/0.01) * 0.01
	want := 10.30

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNormalize(t *testing.T) {
	got := Vec3{10, 20, 2}
	got.Normalize()
	want := Vec3{0.44543540318737396, 0.8908708063747479, 0.0890870806374748}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
