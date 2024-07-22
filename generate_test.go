package main

import "testing"

func Test_generateSexySlug(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test spaces",
			args: args{
				title: "This is the world's weirdest shit CHECK ME OUT &*#)! (Secretly)",
			},
			want: "this-is-the-worlds-weirdest-shit-check-me-out",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateSexySlug(tt.args.title); got != tt.want {
				t.Errorf("generateSexySlug() = %v, want %v", got, tt.want)
			}
		})
	}
}
