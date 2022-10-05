package model

import (
	"fmt"

	"github.com/gochitchat/db/db"
	"github.com/gochitchat/main/db"
	uuid "github.com/satori/go.uuid"
)

type Thread struct {
	id         int32
	uuid       uuid.UUID
	topic      string
	user_id    int32
	created_at string
	updated_at string
}

const (
	queryGetThread        = "SELECT * FROM threads;"
	querGetThreadById     = "SELECT * FROM threads WHERE id = ?;"
	querCreateThread      = "INSERT INTO threads(id,uuid,topic,user_id,created_at,updated_at) values (?,?,?,?);"
	querUpdateThreadTopic = "UPDATE threads SET topic = ?;"
	querDeleteThreadById  = "DELETE FROM threads where id = ?;"
)

func (thread *Thread) GetThread() error {
	smtm, err := db.Pdb.Prepare(queryGetThread)
	if err != nil {
		fmt.Println("You connected to your database.")
	}
	result := smtm.Query()
	if getErr := result.Scan(&thread.id, &thread.uuid, &thread.topic, &thread.user_id, &thread.created_at, &thread.updated_at); getErr != nil {
		return getErr
	}
	return nil
}
