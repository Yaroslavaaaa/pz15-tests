package stringsx

import "testing"

func TestClip(t *testing.T) {
	tests := []struct {
		name string
		s    string
		max  int
		want string
	}{
		{
			name: "empty string",
			s:    "",
			max:  5,
			want: "",
		},
		{
			name: "max zero",
			s:    "hello",
			max:  0,
			want: "",
		},
		{
			name: "max negative",
			s:    "hello",
			max:  -3,
			want: "",
		},
		{
			name: "max equals length",
			s:    "hello",
			max:  5,
			want: "hello",
		},
		{
			name: "max greater than length",
			s:    "hello",
			max:  10,
			want: "hello",
		},
		{
			name: "max less than length",
			s:    "hello",
			max:  3,
			want: "hel",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Clip(tt.s, tt.max)
			if got != tt.want {
				t.Fatalf(
					"Clip(%q, %d) = %q; want %q",
					tt.s, tt.max, got, tt.want,
				)
			}
		})
	}
}
