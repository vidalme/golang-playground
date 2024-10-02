package main

import (
	"testing"
)

func ExampleMain() {
	main()
	// Output: Hello World
}
func TestGreet(t *testing.T) {

	type testCase struct {
		lang language
		want string
	}

	tests := map[string]testCase{
		"Greek": {
			lang: "el",
			want: "Χαίρετε Κόσμε",
		},
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Hebrew": {
			lang: "he",
			want: "שלום עולם",
		},
		"Urdu": {
			lang: "ur",
			want: "ہیلو",
		},
		"Vietnamese": {
			lang: "vi",
			want: "Xin chào Thế Giới",
		},
		"Portuguese": {
			lang: "pt",
			want: "Olá mundo",
		},
	}

	//range over all the scenarios
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("got: %q wanted: %q ", got, tc.want)
			}
		})
	}
}

// 	want := "Hello World"
// 	got := greet()
// 	if got != want {
// 		t.Errorf("expected: %q, got: %q\n", want, got)
// 	}
// }
