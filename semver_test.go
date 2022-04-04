package semver

import (
	"testing"
)

func TestSemVer_String(t *testing.T) {
	tests := []struct {
		name string
		v    SemVer
		want string
	}{
		{name: "1", v: SemVer{}, want: "0.0.0"},
		{name: "2", v: SemVer{Minor: 2, Patch: 3}, want: "0.2.3"},
		{name: "3", v: SemVer{Major: 1, Patch: 3}, want: "1.0.3"},
		{name: "4", v: SemVer{Major: 1, Minor: 2}, want: "1.2.0"},
		{name: "5", v: SemVer{Major: 1, Minor: 2, Patch: 3}, want: "1.2.3"},
		{name: "6", v: SemVer{Major: 1, Minor: 2, Patch: 3, PreRelease: "preRelease"}, want: "1.2.3-preRelease"},
		{name: "7", v: SemVer{Major: 1, Minor: 2, Patch: 3, Build: "build"}, want: "1.2.3+build"},
		{name: "8", v: SemVer{Major: 1, Minor: 2, Patch: 3, PreRelease: "preRelease", Build: "build"}, want: "1.2.3-preRelease+build"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("SemVer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSemVer_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		v       *SemVer
		args    args
		wantErr bool
	}{
		{name: "01", args: args{}, v: &SemVer{}, wantErr: true},
		{name: "02", args: args{text: nil}, v: &SemVer{}, wantErr: true},
		{name: "03", args: args{text: []byte{}}, v: &SemVer{}, wantErr: true},
		{name: "04", args: args{text: []byte("")}, v: &SemVer{}, wantErr: true},
		{name: "05", args: args{text: []byte(" ")}, v: &SemVer{}, wantErr: true},
		{name: "06", args: args{text: []byte(".")}, v: &SemVer{}, wantErr: true},
		{name: "07", args: args{text: []byte("..")}, v: &SemVer{}, wantErr: true},
		{name: "08", args: args{text: []byte(".2.3")}, v: &SemVer{}, wantErr: true},
		{name: "09", args: args{text: []byte("1..3")}, v: &SemVer{}, wantErr: true},
		{name: "10", args: args{text: []byte("1.2.")}, v: &SemVer{}, wantErr: true},
		{name: "11", args: args{text: []byte("1.2.3-")}, v: &SemVer{}, wantErr: true},
		{name: "12", args: args{text: []byte("1.2.3+")}, v: &SemVer{}, wantErr: true},
		{name: "13", args: args{text: []byte("1.2.3-+")}, v: &SemVer{}, wantErr: true},
		{name: "1", args: args{text: []byte("1.2.3")}, v: &SemVer{}},
		{name: "2", args: args{text: []byte("1.2.3-alpha")}, v: &SemVer{}},
		{name: "3", args: args{text: []byte("1.2.3+-")}, v: &SemVer{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.v.UnmarshalText([]byte(tt.args.text))
			if (err != nil) != tt.wantErr {
				t.Errorf("SemVer.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
