package semver

import (
	"testing"
)

func TestCompare(t *testing.T) {
	type args struct {
		sv1 SemVer
		sv2 SemVer
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "01", args: args{sv1: SemVer{Major: -1}}, wantErr: true},
		{name: "1", args: args{}, want: 0},
		{name: "2", args: args{sv1: SemVer{}, sv2: SemVer{}}, want: 0},
		{name: "3", args: args{sv1: ParseMust("1.2.3"), sv2: ParseMust("1.2.3")}, want: 0},
		{name: "4", args: args{sv1: ParseMust("1.2.3"), sv2: ParseMust("2.1.8")}, want: -1},
		{name: "5", args: args{sv1: ParseMust("1.2.3"), sv2: ParseMust("1.1.8")}, want: 1},
		{name: "6", args: args{sv1: ParseMust("1.2.3"), sv2: ParseMust("1.2.8")}, want: -1},
		{name: "7", args: args{sv1: ParseMust("1.0.0-alpha"), sv2: ParseMust("1.0.0")}, want: -1},
		{name: "8", args: args{sv1: ParseMust("1.0.0+alpha"), sv2: ParseMust("1.0.0")}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Compare(tt.args.sv1, tt.args.sv2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
