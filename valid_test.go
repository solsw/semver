package semver

import (
	"testing"
)

func TestValid(t *testing.T) {
	type args struct {
		sv SemVer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "01", args: args{sv: SemVer{Major: -1}}, wantErr: true},
		{name: "02", args: args{sv: SemVer{Minor: -1}}, wantErr: true},
		{name: "03", args: args{sv: SemVer{Patch: -1}}, wantErr: true},
		{name: "04", args: args{sv: SemVer{Major: 1, PreRelease: "."}}, wantErr: true},
		{name: "05", args: args{sv: SemVer{Minor: 1, Build: "."}}, wantErr: true},
		{name: "06", args: args{sv: SemVer{Major: 1, PreRelease: "01"}}, wantErr: true},
		{name: "07", args: args{sv: SemVer{Minor: 1, PreRelease: "0.03.7"}}, wantErr: true},
		{name: "08", args: args{sv: SemVer{Major: 1, PreRelease: "_"}}, wantErr: true},
		{name: "1", args: args{sv: SemVer{}}, wantErr: false},
		{name: "2", args: args{sv: SemVer{Minor: 1, PreRelease: "0"}}, wantErr: false},
		{name: "3", args: args{sv: SemVer{Major: 1, PreRelease: "0.3.7"}}, wantErr: false},
		{name: "4", args: args{sv: SemVer{Minor: 1, Build: "0.3.7"}}, wantErr: false},
		{name: "5", args: args{sv: SemVer{Patch: 1, Build: "0.3.07"}}, wantErr: false},
		{name: "6", args: args{sv: SemVer{Minor: 1, Build: "01"}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Valid(tt.args.sv); (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
