package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "Wassap yo yo",
	}
	t.Run("Valid key string", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "Wassap yo yo"
		assertStrings(t, got, want)
	})
	t.Run("Invalid key string", func(t *testing.T) {
		_, got := dictionary.Search("testadadadad")
		if got == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, got, keyNotFoundError)
	})
}
func TestAdd(t *testing.T) {
	t.Run("Add new key-value pair", func(t *testing.T) {
		dictionary := Dictionary{}
		//key
		word := "test"
		//value
		definition := "testinginMoAko"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("key/word already exists", func(t *testing.T) {
		word := "test"
		definition := "testinginMoAko"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "New test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing key", func(t *testing.T) {
		word := "test"
		definition := "testinginMoAko"
		dictionary := Dictionary{word: definition}
		newDefinition := "New test"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, keyNotFoundError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete existing key", func(t *testing.T) {
		word := "test"
		definition := "testinginMoAko"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		assertError(t, err, keyNotFoundError)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("Should find added word", err)
	}
	assertStrings(t, got, value)
}
