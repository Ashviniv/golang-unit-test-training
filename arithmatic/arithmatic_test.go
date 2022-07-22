package arithmatic

import "testing"

//func TestDivide(t *testing.T)  {
//	out, err := div(8, 4)
//
//	if out != 2 {
//		t.Errorf("Failed to divide 8/4")
//	}
//
//	out = div(0, 4)
//	if out != 0 {
//		t.Errorf("Failed to divide 0/4")
//	}
//
//	out = div(3, 0)
//	if out != 0 {
//		t.Errorf("Failed to divide 3/0")
//	}
//}

func TestTableDrivenDivide(t *testing.T)  {
	tests := []struct{
		name string
		x float64
		y float64
		expected float64
		wantErr bool
	}{
		{ "divide 4 by 8", 4, 8, 0.5, false},
		{ "divide 8 by 4", 8, 4, 2, false},
		{"divide 8 by 0", 8, 0, 0, true},
		{"divide 0 by 8", 0, 8, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
		   	result, err := div(tt.x, tt.y)
			if err == nil && tt.wantErr {
				t.Error("wanted error, got no error")
			}

			if result != tt.expected {
				   t.Errorf("division for %f by %f failed, got: %f, expected: %f", tt.x, tt.y, result, tt.expected)
			}
		})
	}
}
