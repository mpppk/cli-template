package usecase_test

import (
	"testing"

	"github.com/mpppk/cli-template/pkg/usecase"
)

func TestCalcL1NormFromStringSlice(t *testing.T) {
	type args struct {
		strNumbers []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				strNumbers: []string{"1", "2"},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				strNumbers: []string{"-1", "2"},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				strNumbers: []string{"1", "a"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := usecase.CalcL1NormFromStringSlice(tt.args.strNumbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcL1NormFromStringSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalcL1NormFromStringSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcSumFromStringSlice(t *testing.T) {
	type args struct {
		strNumbers []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				strNumbers: []string{"1", "2"},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				strNumbers: []string{"-1", "2"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				strNumbers: []string{"1", "a"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := usecase.CalcSumFromStringSlice(tt.args.strNumbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcSumFromStringSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalcSumFromStringSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcSum(t *testing.T) {
	type args struct {
		strNumbers []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				strNumbers: []int{1, 2},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				strNumbers: []int{-1, 2},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := usecase.CalcSum(tt.args.strNumbers)
			if got != tt.want {
				t.Errorf("CalcSumFromStringSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcL1Norm(t *testing.T) {
	type args struct {
		strNumbers []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				strNumbers: []int{1, 2},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				strNumbers: []int{-1, 2},
			},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := usecase.CalcL1Norm(tt.args.strNumbers)
			if got != tt.want {
				t.Errorf("CalcSumFromStringSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
