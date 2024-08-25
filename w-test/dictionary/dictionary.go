package dictionary

type Dictionary map[string]string

var (
	ErrNotFound      = DictionaryErr("could not find")
	ErrAlreadyExists = DictionaryErr("already exists")
	ErrDoesNotExist  = DictionaryErr("cannot update because it does not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word string, newDef string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrDoesNotExist
	case nil:
		d[word] = newDef
	default:
		return err
	}

	return nil
}
