package list

import "errors"

var ErrListId error = errors.New("list id invalid")
var ErrListInvalid error = errors.New("list is invalid")
var ErrListInternal error = errors.New("list internal error")
