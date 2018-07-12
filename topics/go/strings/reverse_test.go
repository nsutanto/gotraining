package strings

import "testing"

type reverseTest struct {
	name string
	in string
	want string
}

func TestReverse(t *testing.T) {
	tests := []reverseTest{
		{"odd", "jacob", "bocaj"},
		{"evan", "kell", "llek"},
	}

	for _, test := range tests {
		fn := func(t *testing.T) {
			got := Reverse(test.in)

			if got != test.want {
				t.Errorf("Reverse(%q) = %q, wanted %q", test.in, got, test.want)
			}
		}
		t.Run(test.name, fn)
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("abcefghjijklmnopqrstuvwxyz")
	}
}