package maps

type Dictionary map[string]string

const (
	keyNotFoundError    = DictionaryErr("key not found")
	ErrWordExists       = DictionaryErr("key already exists")
	ErrWordDoesNotExist = DictionaryErr("key not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, err := d[word]
	if !err {
		return "", keyNotFoundError
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case keyNotFoundError:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case keyNotFoundError:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	delete(d, word)
}
