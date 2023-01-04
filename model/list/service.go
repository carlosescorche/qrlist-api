package list

import (
	"fmt"
)

func FindById(id string) (*List, error) {
	list, err := findById(id)
	if err != nil {
		return nil, fmt.Errorf("%w:%v", ErrListInternal, err)
	}
	return list, err
}

func FindSubsById(id string) (*List, error) {
	list, err := findSubsById(id)
	if err != nil {
		return nil, fmt.Errorf("%w:%v", ErrListInternal, err)
	}
	return list, err
}

func Insert(list *List) error {
	err := insert(list)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrListInternal, err)
	}
	return nil
}
