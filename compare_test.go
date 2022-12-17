package semver

import (
	"testing"
)

func parseMust(s string) SemVer {
	sv, _ := Parse(s)
	return sv
}

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
		{name: "02", args: args{sv1: SemVer{Minor: -1}, sv2: SemVer{Major: 1}}, wantErr: true},
		{name: "03", args: args{sv1: SemVer{Minor: 1}, sv2: SemVer{Major: -1}}, wantErr: true},

		{name: "1", args: args{}, want: 0},
		{name: "2", args: args{sv1: SemVer{}, sv2: SemVer{}}, want: 0},
		{name: "3", args: args{sv1: SemVer{Major: 1}, sv2: SemVer{Major: 2}}, want: -1},
		{name: "4", args: args{sv1: SemVer{Minor: 2}, sv2: SemVer{Minor: 3}}, want: -1},
		{name: "5", args: args{sv1: SemVer{Patch: 3}, sv2: SemVer{Patch: 2}}, want: 1},
		{name: "6", args: args{sv1: SemVer{PreRelease: "1"}, sv2: SemVer{PreRelease: ""}}, want: -1},
		{name: "7", args: args{sv1: SemVer{PreRelease: ""}, sv2: SemVer{PreRelease: "1"}}, want: 1},
		{name: "8", args: args{sv1: SemVer{PreRelease: "1"}, sv2: SemVer{PreRelease: "1"}}, want: 0},
		{name: "9", args: args{sv1: SemVer{PreRelease: "2"}, sv2: SemVer{PreRelease: "1"}}, want: 1},
		{name: "10", args: args{sv1: SemVer{PreRelease: "1"}, sv2: SemVer{PreRelease: "a"}}, want: -1},
		{name: "11", args: args{sv1: SemVer{PreRelease: "a"}, sv2: SemVer{PreRelease: "1"}}, want: 1},
		{name: "12", args: args{sv1: SemVer{PreRelease: "a"}, sv2: SemVer{PreRelease: "b"}}, want: -1},
		{name: "13", args: args{sv1: SemVer{PreRelease: "a"}, sv2: SemVer{PreRelease: "a.b"}}, want: -1},
		{name: "14", args: args{sv1: SemVer{PreRelease: "1.2.a"}, sv2: SemVer{PreRelease: "1.2"}}, want: 1},
		{name: "15", args: args{sv1: parseMust("1.2.3"), sv2: parseMust("1.2.3")}, want: 0},
		{name: "16", args: args{sv1: parseMust("1.2.3"), sv2: parseMust("2.1.8")}, want: -1},
		{name: "17", args: args{sv1: parseMust("1.2.3"), sv2: parseMust("1.1.8")}, want: 1},
		{name: "18", args: args{sv1: parseMust("1.2.3"), sv2: parseMust("1.2.8")}, want: -1},
		{name: "19", args: args{sv1: parseMust("1.0.0-alpha"), sv2: parseMust("1.0.0")}, want: -1},
		{name: "20", args: args{sv1: parseMust("1.0.0+alpha"), sv2: parseMust("1.0.0")}, want: 0},
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
