package main

import (
	"messaging/data/storage"

	"github.com/sophed/lg"
)

const STORAGE_DIR = "storage/"

func storageType(st string) error {
	if st == "json" {
		// defining JSON storage method with file paths
		storage.METHOD = &storage.StorageJSON{
			UsersFile:         STORAGE_DIR + "users.json",
			MessagesFile:      STORAGE_DIR + "messages.json",
			ConversationsFile: STORAGE_DIR + "conversations.json",
		}
	} else {
		lg.Fatl("invalid storage type selected")
	}
	lg.Info("connecting to " + st + "...")
	return storage.METHOD.Test()
}
