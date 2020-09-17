package maps

// Dictionary defines a dictionary.
type Dictionary map[string]string

var (
	// ErrNotFound - search error.
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists - word exists.
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExist - word not exists.
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr - error of dictionary.
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search returns the value of the given word key.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add - add the defination into the dictionary.
func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update updates the defination of the given word.
func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete deletes the defination from the dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
