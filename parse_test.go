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
		want    *SemVer
		wantErr bool
	}{
		{name: "00", args: args{}, wantErr: true},
		{name: "01", args: args{s: "1.0.0-"}, wantErr: true},
		{name: "1", args: args{s: "1.2.3"}, want: &SemVer{Major: 1, Minor: 2, Patch: 3}},
		{name: "2", args: args{s: "1.0.0-x-y-z.-"}, want: &SemVer{Major: 1, PreRelease: "x-y-z.-"}},
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
