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
		{name: "1",
			v:    SemVer{},
			want: "0.0.0",
		},
		{name: "2",
			v:    SemVer{Minor: 2, Patch: 3},
			want: "0.2.3",
		},
		{name: "3",
			v:    SemVer{Major: 1, Patch: 3},
			want: "1.0.3",
		},
		{name: "4",
			v:    SemVer{Major: 1, Minor: 2},
			want: "1.2.0",
		},
		{name: "5",
			v:    SemVer{Major: 1, Minor: 2, Patch: 3},
			want: "1.2.3",
		},
		{name: "6",
			v:    SemVer{Major: 1, Minor: 2, Patch: 3, PreRelease: "preRelease"},
			want: "1.2.3-preRelease",
		},
		{name: "7",
			v:    SemVer{Major: 1, Minor: 2, Patch: 3, Build: "build"},
			want: "1.2.3+build",
		},
		{name: "8",
			v:    SemVer{Major: 1, Minor: 2, Patch: 3, PreRelease: "preRelease", Build: "build"},
			want: "1.2.3-preRelease+build",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("SemVer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSemVer_EqualTo(t *testing.T) {
	tests := []struct {
		name    string
		v       SemVer
		other   SemVer
		want    bool
		wantErr bool
	}{
		{name: "invalid receiver", v: SemVer{Major: -1}, wantErr: true},
		{name: "invalid argument", other: SemVer{Major: -1}, wantErr: true},
		{name: "equal precedence", v: SemVer{Build: "first"}, other: SemVer{Build: "second"}, want: true},
		{name: "different precedence", other: SemVer{Patch: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.EqualTo(tt.other)
			if got != tt.want || (err != nil) != tt.wantErr {
				t.Errorf("EqualTo() = (%v, %v), want equal=%v, wantErr=%v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestSemVer_UnmarshalTextReceiver(t *testing.T) {
	var v SemVer
	want := SemVer{Major: 1, Minor: 2, Patch: 3, PreRelease: "rc.1", Build: "001"}
	if err := v.UnmarshalText([]byte("1.2.3-rc.1+001")); err != nil {
		t.Fatal(err)
	}
	if v != want {
		t.Fatalf("after successful decode got %#v, want %#v", v, want)
	}
	if err := v.UnmarshalText([]byte("4.5.6+")); err == nil {
		t.Fatal("expected an error for empty build metadata")
	}
	if v != want {
		t.Errorf("after failed decode got %#v, want %#v", v, want)
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
		{name: "001",
			v:       &SemVer{},
			args:    args{},
			wantErr: true,
		},
		{name: "002",
			v:       &SemVer{},
			args:    args{text: nil},
			wantErr: true,
		},
		{name: "003",
			v:       &SemVer{},
			args:    args{text: []byte{}},
			wantErr: true,
		},
		{name: "004",
			v:       &SemVer{},
			args:    args{text: []byte("")},
			wantErr: true,
		},
		{name: "005",
			v:       &SemVer{},
			args:    args{text: []byte(" ")},
			wantErr: true,
		},
		{name: "006",
			v:       &SemVer{},
			args:    args{text: []byte(".")},
			wantErr: true,
		},
		{name: "007",
			v:       &SemVer{},
			args:    args{text: []byte("..")},
			wantErr: true,
		},
		{name: "008",
			v:       &SemVer{},
			args:    args{text: []byte(".2.3")},
			wantErr: true,
		},
		{name: "009",
			v:       &SemVer{},
			args:    args{text: []byte("1..3")},
			wantErr: true,
		},
		{name: "010",
			v:       &SemVer{},
			args:    args{text: []byte("1.2.")},
			wantErr: true,
		},
		{name: "011",
			v:       &SemVer{},
			args:    args{text: []byte("1.2.3-")},
			wantErr: true,
		},
		{name: "012",
			v:       &SemVer{},
			args:    args{text: []byte("1.2.3+")},
			wantErr: true,
		},
		{name: "013",
			v:       &SemVer{},
			args:    args{text: []byte("1.2.3-+")},
			wantErr: true,
		},
		{name: "1",
			v:    &SemVer{},
			args: args{text: []byte("1.2.3")},
		},
		{name: "2",
			v:    &SemVer{},
			args: args{text: []byte("1.2.3-alpha")},
		},
		{name: "3",
			v:    &SemVer{},
			args: args{text: []byte("1.2.3+-")},
		},
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
