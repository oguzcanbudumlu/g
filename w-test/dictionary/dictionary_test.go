package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatalf("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		def := "this is just a test"
		err := dictionary.Add(word, def)
		assertNoError(t, err)
		assertDefinition(t, dictionary, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		d := Dictionary{word: def}
		err := d.Add(word, def)
		assertError(t, err, ErrAlreadyExists)
		assertDefinition(t, d, word, def)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dictionary := Dictionary{word: def}

		newDef := "new def"

		err := dictionary.Update(word, newDef)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		newDef := "new def"

		err := dictionary.Update(word, newDef)

		assertError(t, err, ErrDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	d := Dictionary{word: "def"}

	d.Delete(word)

	_, err := d.Search(word)
	assertError(t, err, ErrNotFound)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertNoError(t testing.TB, got error) {
	if got != nil {
		t.Errorf("did not want error but got one")
	}
}
