package sum_test

import (
	"fmt"
	"github.com/mpppk/cli-template/pkg/sum"
	"github.com/mpppk/cli-template/pkg/util"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func ExampleSum() {
	numbers := []int{1, -2, 3}
	fmt.Println(sum.Sum(numbers))
	// Output:
	// 2
}

func TestSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name: "return sum of numbers",
			args: args{
				numbers: []int{1, 2, 3},
			},
			wantSum: 6,
		},
		{
			name: "return sum of numbers",
			args: args{
				numbers: []int{1, -2, 3},
			},
			wantSum: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := sum.Sum(tt.args.numbers); gotSum != tt.wantSum {
				t.Errorf("Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestSumFromFile(t *testing.T) {
	type args struct {
		numbers []int
	}

	type tCase struct {
		name    string
		args    args
		wantSum int
	}

	f, err := os.Open("../../testdata/sum.txt")
	if err != nil {
		t.Fatal(err)
	}
	contentBytes, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	contents := string(contentBytes)
	lines := strings.Split(contents, "\n")
	var tests []tCase
	for i, line := range lines {
		strRow := strings.Split(line, " ")
		row, err := util.ConvertStringSliceToIntSlice(strRow)
		if err != nil {
			t.Fatal(err)
		}
		want := row[len(row)-1]
		nums := row[:len(row)-1]
		tc := tCase{
			name: fmt.Sprintf("case%d\n", i),
			args: args{
				numbers: nums,
			},
			wantSum: want,
		}
		tests = append(tests, tc)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := sum.Sum(tt.args.numbers); gotSum != tt.wantSum {
				t.Errorf("Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestSumFromString(t *testing.T) {
	type args struct {
		stringNumbers []string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
		wantErr bool
	}{
		{
			name: "return sum of numbers",
			args: args{
				stringNumbers: []string{"1", "2", "3"},
			},
			wantSum: 6,
			wantErr: false,
		},
		{
			name: "will be error if args includes not number string",
			args: args{
				stringNumbers: []string{"1", "2", "a"},
			},
			wantSum: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, err := sum.SumFromString(tt.args.stringNumbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSum != tt.wantSum {
				t.Errorf("SumFromString() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func ExampleL1Norm() {
	numbers := []int{1, -2, 3}
	fmt.Println(sum.L1Norm(numbers))
	// Output:
	// 6
}

func TestL1Norm(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name: "return sum of numbers",
			args: args{
				numbers: []int{1, 2, 3},
			},
			wantSum: 6,
		},
		{
			name: "return sum of numbers",
			args: args{
				numbers: []int{1, -2, 3},
			},
			wantSum: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := sum.L1Norm(tt.args.numbers); gotSum != tt.wantSum {
				t.Errorf("Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
