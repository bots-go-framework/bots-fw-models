package botsfwdal

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("not found")

func IsNotFoundErr(err error) bool {
	return errors.Is(err, errNotFound)
}

func NotFoundErr(err error) error {
	return fmt.Errorf("%w: %v", errNotFound, err)
}
