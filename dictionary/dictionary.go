package dictionary

func Search(dict map[string]string, keyword string) string {
	return dict[keyword]
}

type Dictionary map[string]string
type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

const (
	ErrNoResults          = DictionaryError("No results")
	ErrDefinitionExists   = DictionaryError("Definition already exists")
	ErrDefinitionNotFound = DictionaryError("Definition not found")
)

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]

	if !ok {
		return "", ErrNoResults
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNoResults:
		d[key] = value
	case nil:
		return ErrDefinitionExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key string, newValue string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNoResults:
		return ErrDefinitionNotFound
	case nil:
		d[key] = newValue
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNoResults:
		return ErrDefinitionNotFound
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}
