package banana

import "errors"

var (
	UserConflict = errors.New("user already exist")
	SignUpFailed = errors.New("sign up Failed ")
	UserNotFound = errors.New("user not found")
)
