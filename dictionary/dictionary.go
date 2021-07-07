package dictionary

const (
	errorNotFound     = DictionaryError("could not find the word you are looking for")
	errorKeyExisting  = DictionaryError("key already existed!")
	errKeyNotExisting = DictionaryError("Key is not existed!")
)

type DictionaryError string

type Dictionary map[string]string

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if ok {
		return definition, nil
	}
	return "", errorNotFound
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case errorNotFound:
		d[key] = value
		return nil
	case nil:
		return errorKeyExisting
	default:
		return err
	}
}
func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = value
		return nil
	case errorNotFound:
		return errKeyNotExisting
	default:
		return err
	}
}
func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		delete(d, key)
		return nil
	case errorNotFound:
		return errKeyNotExisting
	default:
		return err
	}
}

func (d DictionaryError) Error() string {
	return string(d)
}
