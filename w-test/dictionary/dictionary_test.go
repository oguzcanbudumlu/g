package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func Search(dictionary map[string]string, s string) string {
	vl, _ := dictionary[s]
	return vl
}
