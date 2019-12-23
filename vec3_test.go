package vec3

import "testing"

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
		got.Add(1, 1, 1)
		want := Vec3{4, 7, 8}

		check(got, want)
	})
}
