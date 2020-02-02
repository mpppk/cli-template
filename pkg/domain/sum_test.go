package domain

import (
	"reflect"
	"testing"
)

func TestNewNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want Numbers
	}{
		{
			name: "",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNumbersFromStringSlice(t *testing.T) {
	type args struct {
		strNumbers []string
	}
	tests := []struct {
		name    string
		args    args
		want    Numbers
		wantErr bool
	}{
		{
			args: args{
				strNumbers: []string{"1"},
			},
			want:    []int{1},
			wantErr: false,
		},
		{
			args: args{
				strNumbers: []string{"1", "2"},
			},
			want:    []int{1, 2},
			wantErr: false,
		},
		{
			args: args{
				strNumbers: []string{"-1", "2"},
			},
			want:    []int{-1, 2},
			wantErr: false,
		},
		{
			args: args{
				strNumbers: []string{"1", "a"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNumbersFromStringSlice(tt.args.strNumbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewNumbersFromStringSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNumbersFromStringSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbers_CalcL1Norm(t *testing.T) {
	tests := []struct {
		name string
		n    Numbers
		want int
	}{
		{
			n:    []int{1, 2},
			want: 3,
		},
		{
			n:    []int{-1, 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.CalcL1Norm(); got != tt.want {
				t.Errorf("CalcL1Norm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbers_CalcSum(t *testing.T) {
	tests := []struct {
		name string
		n    Numbers
		want int
	}{
		{
			n:    []int{1, 2},
			want: 3,
		},
		{
			n:    []int{-1, 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.CalcSum(); got != tt.want {
				t.Errorf("CalcSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
