package storage

import (
	"errors"
	"messaging/data"
	"messaging/data/entities"
)

var ErrNotFound = errors.New("query returned no results") // abstract 'not found' error
var METHOD StorageMethod

// abstract interface for StorageMethod type
type StorageMethod interface {
	
	// AddUser stores a provided user
	AddUser(*entities.User) error

	// RemoveUser deletes a provided user from storage by their ID
	RemoveUser(string) error
	
	// FindUser searches for a user and returns the results
	FindUser(*data.SearchableUser) (*entities.User, error)
	
	// Test ensures the storage method can be read from and returns an error on failure
	Test() error

}
