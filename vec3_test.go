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
