package storage

import (
	"log"

	"github.com/Ria-mehra16/Golang-Cassandra/pkg/deleting"
)

func (s *Storage) DeleteCandy(c deleting.Candy) (string, error) {
	// id := uuid.New().String()
	if err := s.db.Query(`DELETE FROM candies WHERE id=? IF EXISTS`,
		c.Id).Exec(); err != nil {
		log.Println("Error while trying to delete from DB: ", err)
		return "", err
	}
	return c.Id, nil
}
