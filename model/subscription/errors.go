package subscription

import "errors"

var ErrSubscriptionId error = errors.New("subscription id invalid")
var ErrSubscriptionInvalid error = errors.New("subscription is invalid")
var ErrSubscriptionInternal error = errors.New("subscription internal error")
