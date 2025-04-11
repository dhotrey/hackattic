package main

import "testing"

func Test_checkLeadingBits(t *testing.T) {
	type args struct {
		hash string
		d    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "t1", args: args{
			hash: "[00001111]",
			d:    4,
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkLeadingBits(tt.args.hash, tt.args.d); got != tt.want {
				t.Errorf("checkLeadingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
