package semver

import (
	"testing"
)

func TestValid(t *testing.T) {
	type args struct {
		sv *SemVer
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "01", args: args{sv: &SemVer{Major: -1}}, want: false},
		{name: "02", args: args{sv: &SemVer{Minor: -1}}, want: false},
		{name: "03", args: args{sv: &SemVer{Patch: -1}}, want: false},
		{name: "04", args: args{sv: &SemVer{Minor: 1, PreRelease: "0.03.7"}}, want: false},
		{name: "1", args: args{sv: &SemVer{}}, want: true},
		{name: "2", args: args{sv: &SemVer{Major: 1, PreRelease: "0.3.7"}}, want: true},
		{name: "3", args: args{sv: &SemVer{Patch: 1, Build: "0.3.07"}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Valid(tt.args.sv)
			if (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
