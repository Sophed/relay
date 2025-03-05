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

// userList exists solely to parse the defined JSON files
type userList struct {
	Entries []entities.User `json:"users"`
}

// messageList exists solely to parse the defined JSON files
type messageList struct {
	Entries []entities.Message `json:"messages"`
}

// conversationList exists solely to parse the defined JSON files
type conversationList struct {
	Entries []entities.Conversation `json:"conversation"`
}

// AddUser implements StorageMethod AddUser()
func (j *StorageJSON) AddUser(u *entities.User) error {
	users, err := j.readUsers()
	if err != nil {
		return err
	}
	// append user to the list and write it back to the file
	users.Entries = append(users.Entries, *u)
	err = j.writeUsers(users)
	if err != nil {
		return err
	}
	return nil
}

// RemoveUser implements StorageMethod RemoveUser()
func (j *StorageJSON) RemoveUser(string) error {
	return nil
}

// FindUser implements StorageMethod FindUser()
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

// Test implements StorageMethod Test()
func (j *StorageJSON) Test() error {
	_, err := j.readUsers()
	if err != nil {
		return err // fail if users file cannot be read
	}
	/* TODO: impl
	_, err = j.readMessages()
	if err != nil {
		return err // fail if messages file cannot be read
	}
	_, err = j.readConversations()
	*/
	return err // fail if conversations file cannot be read
}

// readUsers returns a list of entities.User objects
func (j *StorageJSON) readUsers() (*userList, error) {
	data, err := os.ReadFile(j.UsersFile) // read users file
	if err != nil {
		return nil, err // return error on fail
	}
	var list userList
	err = json.Unmarshal(data, &list) // parse JSON as userList type
	return &list, err                 // return list and error result
}

// writeUsers writes a userList to the users file
func (j *StorageJSON) writeUsers(users *userList) error {
	data, err := json.MarshalIndent(users, "", "\t") // convert struct to JSON
	if err != nil {
		return err // error on fail
	}
	err = os.WriteFile(j.UsersFile, data, 0644) // write to the users file
	return err                                  // return error on fail
}

func (j *StorageJSON) readMessages() (*messageList, error) {
	data, err := os.ReadFile(j.MessagesFile) // read messages file
	if err != nil {
		return nil, err // return error on fail
	}
	var list messageList
	err = json.Unmarshal(data, &list) // parse JSON as messageList type
	return &list, err                 // return list and error result
}

func (j *StorageJSON) readConversations() (*conversationList, error) {
	data, err := os.ReadFile(j.ConversationsFile) // read conversations file
	if err != nil {
		return nil, err // return error on fail
	}
	var list conversationList
	err = json.Unmarshal(data, &list) // parse JSON as messageList type
	return &list, err                 // return list and error result
}
