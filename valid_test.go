package semver

import (
	"testing"
)

func TestValid(t *testing.T) {
	type args struct {
		sv SemVer
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "01", args: args{sv: SemVer{Major: -1}}, want: false},
		{name: "02", args: args{sv: SemVer{Minor: -1}}, want: false},
		{name: "03", args: args{sv: SemVer{Patch: -1}}, want: false},
		{name: "04", args: args{sv: SemVer{Major: 1, PreRelease: "."}}, want: false},
		{name: "05", args: args{sv: SemVer{Minor: 1, Build: "."}}, want: false},
		{name: "06", args: args{sv: SemVer{Major: 1, PreRelease: "01"}}, want: false},
		{name: "07", args: args{sv: SemVer{Minor: 1, PreRelease: "0.03.7"}}, want: false},
		{name: "08", args: args{sv: SemVer{Major: 1, PreRelease: "_"}}, want: false},
		{name: "1", args: args{sv: SemVer{}}, want: true},
		{name: "2", args: args{sv: SemVer{Minor: 1, PreRelease: "0"}}, want: true},
		{name: "3", args: args{sv: SemVer{Major: 1, PreRelease: "0.3.7"}}, want: true},
		{name: "4", args: args{sv: SemVer{Minor: 1, Build: "0.3.7"}}, want: true},
		{name: "5", args: args{sv: SemVer{Patch: 1, Build: "0.3.07"}}, want: true},
		{name: "6", args: args{sv: SemVer{Minor: 1, Build: "01"}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.args.sv); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
