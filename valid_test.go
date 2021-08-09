package semver

import "testing"

func TestValid(t *testing.T) {
	type args struct {
		sv *SemVer
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "01", args: args{sv: &SemVer{Major: -1}}},
		{name: "02", args: args{sv: &SemVer{Minor: -1}}},
		{name: "03", args: args{sv: &SemVer{Patch: -1}}},
		{name: "1", args: args{sv: &SemVer{}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.args.sv); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
