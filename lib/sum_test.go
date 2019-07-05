package lib

import "testing"

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
				numbers: []int{1,2,3},
			},
			wantSum: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := Sum(tt.args.numbers); gotSum != tt.wantSum {
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
			gotSum, err := SumFromString(tt.args.stringNumbers)
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
