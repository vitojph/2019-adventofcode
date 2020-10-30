package fuel

import "testing"

func TestFuel(t *testing.T) {
	t1 := Fuel(12)
	t2 := Fuel(14)
	t3 := Fuel(1969)
	t4 := Fuel(100756)
	if t1 != 2 {
		t.Errorf("Fuel for 12: %d; want 2", t1)
	}
	if t2 != 2 {
		t.Errorf("Fuel for 14: %d; want 2", t2)
	}
	if t3 != 966 {
		t.Errorf("Fuel for 1969: %d; want 966", t3)
	}
	if t4 != 50346 {
		t.Errorf("Fuel for 100756: %d; want 50346", t4)
	}
}
