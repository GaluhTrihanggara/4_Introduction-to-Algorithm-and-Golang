package print_calculate //nolint:gomnd

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive values",
			args: args{a: 2, b: 3},
			want: 5,
		},
		{
			name: "negative values",
			args: args{a: -5, b: -8},
			want: -13,
		},
		{
			name: "zero values",
			args: args{a: 0, b: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Addition(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Addition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtraction(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive values",
			args: args{a: 5, b: 2},
			want: 3,
		},
		{
			name: "negative values",
			args: args{a: -5, b: -8},
			want: 3,
		},
		{
			name: "zero values",
			args: args{a: 0, b: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtraction(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Subtraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplying(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive values",
			args: args{a: 5, b: 2},
			want: 10,
		},
		{
			name: "negative values",
			args: args{a: -5, b: -8},
			want: 40,
		},
		{
			name: "zero values",
			args: args{a: 0, b: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplying(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiplying() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "division by zero",
			args:    args{a: 5, b: 0},
			want:    0,
			wantErr: true,
		},
		{
			name:    "positive values",
			args:    args{a: 6, b: 2},
			want:    3,
			wantErr: false,
		},
		{
			name:    "negative values",
			args:    args{a: -5, b: 2},
			want:    -2,
			wantErr: false,
		},
		{
			name:    "zero values",
			args:    args{a: 0, b: 5},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func main() {
	fmt.Println("2 + 3 =", Addition(2, 3))
	fmt.Println("5 - 2 =", Subtraction(5, 2))
	fmt.Println("2 * 3 =", Multiplying(2, 3))
	result, err := Divide(6, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("6 / 2 =", result)
	}
}
