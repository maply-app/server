package errors

import "errors"

var ObjectAlreadyExists = errors.New("object already exists")
var Forbidden = errors.New("forbidden")
var Unauthorized = errors.New("unauthorized")
var ObjectDoesNotExists = errors.New("object does not exists")
