package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound       = errors.New("Not Found")
	errCantUpdate     = errors.New("Cant update non-existing word")
	errWordExistes    = errors.New("Word Exists")
	errWordNotExistes = errors.New("Word Not Exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound

}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExistes
	}
	return nil

	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExistes
	// }
	// return nil
}

// Update dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errWordNotExistes
	}
	return nil
}
