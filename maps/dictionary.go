package maps

var ErrNotFound = DictionaryErr("could not find that word")
var ErrWordExists = DictionaryErr("this word already exists in this dictionary")
var ErrWordDoesNotExist = DictionaryErr("this word does not exist in this dictionary")

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
