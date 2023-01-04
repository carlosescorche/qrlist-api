package subscription

import (
	"fmt"
)

func FindById(id string) (*Subscription, error) {
	return findById(id)
}

func FindByListId(id string) ([]Subscription, error) {
	return findByListId(id)
}

func Insert(s *Subscription) error {
	_, err := insert(s)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrSubscriptionInvalid, err)
	}
	return nil
}

func Update(subs Subscription) error {
	err := update(subs)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrSubscriptionInternal, err)
	}
	return nil
}
