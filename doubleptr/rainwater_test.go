package doubleptr

import "testing"

func Test_trapRainWater(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name      string
		args      args
		wantWater int
	}{
		{
			name:      "trapRainWater",
			args:      args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			wantWater: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWater := trapRainWater(tt.args.height); gotWater != tt.wantWater {
				t.Errorf("trapRainWater() = %v, want %v", gotWater, tt.wantWater)
			}
		})
	}
}
