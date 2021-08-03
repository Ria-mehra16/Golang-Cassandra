package storage

import (
	"log"

	"github.com/Ria-mehra16/Golang-Cassandra/pkg/deleting"
	"github.com/google/uuid"
)

func (s *Storage) DeleteCandy(c deleting.Candy) (string, error) {
	id := uuid.New().String()
	if err := s.db.Query(`DELETE FROM candies WHERE id=(candy_id) VALUES (?) IF EXISTS`,
		id).Exec(); err != nil {
		log.Println("Error while trying to delete from DB: ", err)
		return "", err
	}
	return id, nil
}
