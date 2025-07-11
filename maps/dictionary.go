package main

type Dictionary map[string]string

const (
	ErrorNotFound      = DictionaryErr("could not find the word you were looking for in the dictionary")
	ErrorAlreadyExists = DictionaryErr("this value already exists in the dictionary")
	ErrorDoesNotExist  = DictionaryErr("the word you are trying to update or delete does not exist in the dictionary")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

func (d Dictionary) Search(key string) (string, error) {
	// ok indicates if the item was found sucessfully
	definition, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		return ErrorAlreadyExists
	case ErrorNotFound:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, update string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = update
	case ErrorNotFound:
		return ErrorDoesNotExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	
	switch err {
	case nil:
		// native Go function (dictionary, key)
		delete(d, key)
	case ErrorNotFound:
		return ErrorDoesNotExist
	default:
		return err
	}
	return nil
}
