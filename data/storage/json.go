package storage

import (
	"encoding/json"
	"messaging/data"
	"messaging/data/entities"
	"os"
)

type StorageJSON struct {
	UsersFile         string
	MessagesFile      string
	ConversationsFile string
}
type userList struct {
	Entries []entities.User
}

func (j *StorageJSON) AddUser(u *entities.User) error {
	return nil
}

func (j *StorageJSON) RemoveUser(u *entities.User) error {
	return nil
}

func (j *StorageJSON) FindUser(s *data.SearchableUser) (*entities.User, error) {
	users, err := j.readUsers()
	if err != nil {
		return nil, err
	}
	for _, u := range users.Entries {
		if u.Email == s.Email || u.ID == s.ID || u.Username == s.Username {
			return &u, nil
		}
	}
	return nil, ErrNotFound
}

func (j *StorageJSON) Test() error {
	_, err := j.readUsers()
	if err != nil {
		return err
	}
	_, err = j.readMessages()
	if err != nil {
		return err
	}
	_, err = j.readConversations()
	return err
}

func (j *StorageJSON) readUsers() (*userList, error) {
	data, err := os.ReadFile(j.UsersFile)
	if err != nil {
		return nil, err
	}
	var list userList
	err = json.Unmarshal(data, &list)
	return &list, err
}

func (j *StorageJSON) readMessages() (*userList, error) {
	data, err := os.ReadFile(j.UsersFile)
	if err != nil {
		return nil, err
	}
	var list userList
	err = json.Unmarshal(data, &list)
	return &list, err
}

func (j *StorageJSON) readConversations() (*userList, error) {
	data, err := os.ReadFile(j.UsersFile)
	if err != nil {
		return nil, err
	}
	var list userList
	err = json.Unmarshal(data, &list)
	return &list, err
}
