package golo

import "errors"

func Map[I, O any](input []I, fn func(elem I) O) []O {
	output := make([]O, len(input))

	for i := range input {
		output[i] = fn(input[i])
	}

	return output
}

func Delete[T any](values *[]T, index int) error {
	if values == nil || *&values == nil {
		return errors.New("slice is nil.")
	} else if index < 0 || len(*values)-1 < index {
		return errors.New("index does not exist.")
	}

	if len(*values)-1 > index {
		copy((*values)[index:], (*values)[index+1:])
	}

	var t T
	(*values)[len(*values)-1] = t

	*values = (*values)[:len(*values)-1]

	return nil
}
