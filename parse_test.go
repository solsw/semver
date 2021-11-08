package semver

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    SemVer
		wantErr bool
	}{
		{name: "01", args: args{}, wantErr: true},
		{name: "02", args: args{s: "1.02.3"}, wantErr: true},
		{name: "03", args: args{s: "a.2.3"}, wantErr: true},
		{name: "04", args: args{s: "1.-2.3"}, wantErr: true},
		{name: "05", args: args{s: "1.2.3-"}, wantErr: true},
		{name: "06", args: args{s: "1.2.3+"}, wantErr: true},
		{name: "07", args: args{s: "1.0.-"}, wantErr: true},
		{name: "08", args: args{s: "1.2.3-p+"}, wantErr: true},
		{name: "09", args: args{s: "1.2.3-p_r"}, wantErr: true},
		{name: "10", args: args{s: "1.2.3+b_"}, wantErr: true},
		{name: "1", args: args{s: "0.0.0"}, want: SemVer{}},
		{name: "2", args: args{s: "1.2.3"}, want: SemVer{Major: 1, Minor: 2, Patch: 3}},
		{name: "3", args: args{s: "1.0.0-x-y-z.-"}, want: SemVer{Major: 1, PreRelease: "x-y-z.-"}},
		{name: "4", args: args{s: "1.0.0-p+b"}, want: SemVer{Major: 1, PreRelease: "p", Build: "b"}},
		{name: "5", args: args{s: "1.0.0-p+-b-"}, want: SemVer{Major: 1, PreRelease: "p", Build: "-b-"}},
		{name: "6", args: args{s: "1.0.0--p-+-b-"}, want: SemVer{Major: 1, PreRelease: "-p-", Build: "-b-"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
