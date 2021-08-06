package semver

import (
	"reflect"
	"testing"
)

func TestSemVer_String(t *testing.T) {
	tests := []struct {
		name string
		v    *SemVer
		want string
	}{
		{name: "1", v: &SemVer{}, want: "0.0.0"},
		{name: "2", v: &SemVer{Major: 1, Minor: 2, Patch: 3}, want: "1.2.3"},
		{name: "3", v: &SemVer{Major: 1, Minor: 2, Patch: 3, PreRelease: "preRelease"}, want: "1.2.3-preRelease"},
		{name: "4", v: &SemVer{Major: 1, Minor: 2, Patch: 3, Build: "build"}, want: "1.2.3+build"},
		{name: "5", v: &SemVer{Major: 1, Minor: 2, Patch: 3, PreRelease: "preRelease", Build: "build"}, want: "1.2.3-preRelease+build"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("SemVer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSemVer_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		v       *SemVer
		want    string
		wantErr bool
	}{
		{name: "1", v: &SemVer{}, want: "0.0.0"},
		{name: "2", v: &SemVer{Major: 1, Minor: 2, Patch: 3}, want: "1.2.3"},
		{name: "3", v: &SemVer{Major: 1, Minor: 2, Patch: 3, PreRelease: "preRelease"}, want: "1.2.3-preRelease"},
		{name: "4", v: &SemVer{Major: 1, Minor: 2, Patch: 3, Build: "build"}, want: "1.2.3+build"},
		{name: "5", v: &SemVer{Major: 1, Minor: 2, Patch: 3, PreRelease: "preRelease", Build: "build"}, want: "1.2.3-preRelease+build"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("SemVer.MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sgot := string(got)
			if !reflect.DeepEqual(sgot, tt.want) {
				t.Errorf("SemVer.MarshalText() = %s, want %s", sgot, tt.want)
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
		{name: "001", args: args{}, v: &SemVer{}, wantErr: true},
		{name: "002", args: args{text: nil}, v: &SemVer{}, wantErr: true},
		{name: "003", args: args{text: []byte{}}, v: &SemVer{}, wantErr: true},
		{name: "004", args: args{text: []byte("")}, v: &SemVer{}, wantErr: true},
		{name: "005", args: args{text: []byte(" ")}, v: &SemVer{}, wantErr: true},
		{name: "006", args: args{text: []byte(".")}, v: &SemVer{}, wantErr: true},
		{name: "007", args: args{text: []byte("..")}, v: &SemVer{}, wantErr: true},
		{name: "008", args: args{text: []byte("1..3")}, v: &SemVer{}, wantErr: true},
		{name: "009", args: args{text: []byte("1.2.3-")}, v: &SemVer{}, wantErr: true},
		{name: "010", args: args{text: []byte("1.2.3+")}, v: &SemVer{}, wantErr: true},
		{name: "011", args: args{text: []byte("1.2.3-+")}, v: &SemVer{}, wantErr: true},
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
			if err != nil && tt.wantErr {
				t.Logf("SemVer.UnmarshalText() error = %v", err)
			}
		})
	}
}
