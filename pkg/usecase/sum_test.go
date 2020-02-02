package usecase

import "testing"

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
			got, err := CalcL1NormFromStringSlice(tt.args.strNumbers)
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
			got, err := CalcSumFromStringSlice(tt.args.strNumbers)
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
