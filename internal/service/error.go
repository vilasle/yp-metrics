package service

import "errors"

var ErrEmptyName = errors.New("name is empty")
var ErrEmptyKind = errors.New("type is empty")
var ErrEmptyValue = errors.New("value is empty")
var ErrUnknownKind = errors.New("unknown type of metric")
