package storage

import (
	"errors"
	"messaging/data"
	"messaging/data/entities"
)

var ErrNotFound = errors.New("query returned no results")
var METHOD StorageMethod

type StorageMethod interface {
	AddUser(*entities.User) error
	RemoveUser(*entities.User) error
	FindUser(*data.SearchableUser) (*entities.User, error)
	Test() error
}
