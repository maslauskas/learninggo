package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		want := "this is just a test"
		got, _ := dictionary.Search("test")

		if got != want {
			t.Errorf("expected search result %q with keyword %q, got %q", want, "test", got)
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err != ErrNoResults {
			t.Errorf("expected error %q, got %q", ErrNoResults, err)
		}
	})

	t.Run("add word to dictionary", func(t *testing.T) {
		dictionary.Add("foo", "bar")
		want := "bar"
		got, _ := dictionary.Search("foo")

		if got != want {
			t.Errorf("expected search result %q with keyword %q, got %q", want, "foo", got)
		}
	})

	t.Run("adding existing word should throw error", func(t *testing.T) {
		err := dictionary.Add("test", "this value already exists")

		if err != ErrDefinitionExists {
			t.Errorf("expected error %q, got %q", ErrDefinitionExists, err)
		}
	})

	t.Run("update", func(t *testing.T) {
		dictionary.Update("test", "new value")
		want := "new value"
		got, _ := dictionary.Search("test")

		if got != want {
			t.Errorf("expected value %q, got %q", want, got)
		}
	})

	t.Run("will throw error when updating non existent key", func(t *testing.T) {
		err := dictionary.Update("unknown", "updated value")

		if err != ErrDefinitionNotFound {
			t.Errorf("expected error %q, got %q", ErrDefinitionNotFound, err)
		}
	})

	t.Run("delete", func(t *testing.T) {
		dictionary.Add("delete", "this")
		dictionary.Delete("delete")

		_, err := dictionary.Search("delete")

		if err != ErrNoResults {
			t.Errorf("expected %q definition to be deleted", "delete")
		}
	})

	t.Run("will throw error when deleting unknown definition", func(t *testing.T) {
		err := dictionary.Delete("unknown")

		if err != ErrDefinitionNotFound {
			t.Errorf("expected error %q, got %q", ErrDefinitionNotFound, err)
		}
	})
}
