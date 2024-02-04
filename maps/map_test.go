package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "value"}

	got := Search(dict, "test")
	want := "value"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
