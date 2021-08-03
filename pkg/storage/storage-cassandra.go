package storage

import (
	"log"

	"github.com/Ria-mehra16/Golang-Cassandra/pkg/adding"
	"github.com/Ria-mehra16/Golang-Cassandra/pkg/deleting"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type Storage struct {
	db *gocql.Session
}

func SetupStorage() (*Storage, error) {
	cluster := gocql.NewCluster("172.28.51.166")
	cluster.Keyspace = "candy_shop_db"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return &Storage{}, err
	}
	return &Storage{db: session}, nil
}

func (s *Storage) GetAllCandyNames() ([]string, error) {
	var candy string
	var candies []string
	iter := s.db.Query(`SELECT name FROM candies`).Iter()
	for iter.Scan(&candy) {
		candies = append(candies, candy)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return candies, nil
}

func (s *Storage) AddCandy(c adding.Candy) (string, error) {
	id := uuid.New().String()
	if err := s.db.Query(`INSERT INTO candies (candy_id, category, name, price) VALUES (?, ?, ?, ?)`,
		id, c.Category, c.Name, c.Price).Exec(); err != nil {
		log.Println("Error while trying to save to DB: ", err)
		return "", err
	}
	return id, nil
}

func (s *Storage) DeleteCandy(c deleting.Candy) (string, error) {
	if err := s.db.Query(`DELETE FROM candies WHERE id=? IF EXISTS`,
		c.Id).Exec(); err != nil {
		log.Println("Error while trying to delete from DB: ", err)
		return "", err
	}
	return c.Id, nil
}
