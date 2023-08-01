package shared

import (
	"errors"
	"fmt"
)

var ErrEmpty = errors.New(`empty value`)

var _ error = ErrEmptyIdx{}

type ErrEmptyIdx struct {
	pos uint
}

func NewErrEmptyIdx(pos uint) ErrEmptyIdx {
	return ErrEmptyIdx{pos: pos}
}

func (e ErrEmptyIdx) Error() string {
	return fmt.Sprintf(`empty at pos %d`, e.pos)
}
