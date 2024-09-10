package gogeo

import (
	"testing"
)

func TestDistance(t *testing.T) {
	d := NewPoint(36.256909, 46.265876).DistanceInMeter(NewPoint(36.26729, 46.3783142))
	if d != 10157.930664 {
		t.Fatalf("failed to calc")
	}
}

func TestCircle(t *testing.T) {
	contain := NewPolygon([][]float32{{51.59366908460324, 31.40037506998533}}).CircleContains(NewPoint(36.256909, 46.265876), 123.123)
	if contain {
		t.Fatalf("TestCircle failed")
	}

	contain = NewPolygon([][]float32{{51.59366908460324, 31.40037506998533}}).CircleContains(NewPoint(31.40028516, 51.5928680), 123.123)
	if !contain {
		t.Fatalf("TestCircle failed")
	}
}

func TestRectangle(t *testing.T) {
	contain := NewPolygon([][]float32{{51.718521, 32.708254}, {51.718521, 32.708796}, {51.719561, 32.708796}, {51.719561, 32.708254}, {51.718521, 32.708254}}).RectangleContains(NewPoint(32.708254, 51.708254))
	if contain {
		t.Fatalf("TestRectangle failed")
	}
	contain = NewPolygon([][]float32{{51.718521, 32.708254}, {51.718521, 32.708796}, {51.719561, 32.708796}, {51.719561, 32.708254}, {51.718521, 32.708254}}).RectangleContains(NewPoint(32.708254, 51.718520))
	if !contain {
		t.Fatalf("TestRectangle failed")
	}
}

func TestPolygon(t *testing.T) {
	contain := NewPolygon([][]float32{{51.676791, 32.671763}, {51.676152, 32.671736}, {51.675557, 32.671745}, {51.675407, 32.671546}, {51.675246, 32.671257}, {51.675289, 32.671045}, {51.675546, 32.670774}, {51.675766, 32.670435}, {51.676673, 32.670286}, {51.676925, 32.670516}, {51.676952, 32.670909}, {51.676919, 32.671221}, {51.676791, 32.671763}}).RectangleContains(NewPoint(32.708254, 51.718520))
	if contain {
		t.Fatalf("TestPolygon failed")
	}

	contain = NewPolygon([][]float32{{51.676791, 32.671763}, {51.676152, 32.671736}, {51.675557, 32.671745}, {51.675407, 32.671546}, {51.675246, 32.671257}, {51.675289, 32.671045}, {51.675546, 32.670774}, {51.675766, 32.670435}, {51.676673, 32.670286}, {51.676925, 32.670516}, {51.676952, 32.670909}, {51.676919, 32.671221}, {51.676791, 32.671763}}).RectangleContains(NewPoint(32.671763,51.676791))
	if !contain {
		t.Fatalf("TestPolygon failed")
	}
}
