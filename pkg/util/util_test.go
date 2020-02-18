package util

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestConvertStringSliceToIntSlice(t *testing.T) {
	type args struct {
		stringSlice []string
	}
	tests := []struct {
		name         string
		args         args
		wantIntSlice []int
		wantErr      bool
	}{
		{
			name: "can convert string slice to int slice",
			args: args{
				stringSlice: []string{"1", "2", "3"},
			},
			wantIntSlice: []int{1, 2, 3},
			wantErr:      false,
		},
		{
			name: "will be error if string can not be convert to number",
			args: args{
				stringSlice: []string{"1", "2", "a"},
			},
			wantIntSlice: nil,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIntSlice, err := ConvertStringSliceToIntSlice(tt.args.stringSlice)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertStringSliceToIntSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIntSlice, tt.wantIntSlice) {
				t.Errorf("ConvertStringSliceToIntSlice() = %v, want %v", gotIntSlice, tt.wantIntSlice)
			}
		})
	}
}

func TestPrettyPrintError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				err: errors.New("sample error"),
			},
			want: fmt.Sprintln("Error: sample error"),
		},
		{
			name: "",
			args: args{
				err: fmt.Errorf("a: %w", errors.New("b")),
			},
			want: fmt.Sprintln("Error: a") + fmt.Sprintln("  b"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrettyPrintError(tt.args.err); got != tt.want {
				t.Errorf("PrettyPrintError() = %v, want %v", got, tt.want)
			}
		})
	}
}
