package graphdb

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type RequestRepository interface {
	GetRequests() ([]string, error)
}

type RequestNeo4jRepository struct {
	Driver neo4j.Driver
}

func NewRequestRepository(uri, username, password string) (*RequestNeo4jRepository, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	return &RequestNeo4jRepository{
		Driver: driver,
	}, nil
}

func (r *RequestNeo4jRepository) GetRequests() (requests []string, err error) {
	session := r.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer func() {
		err = session.Close()
	}()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return getAllRequests(tx)
	})
	if result == nil {
		return nil, err
	}
	requests = result.([]string)
	return requests, err
}

func getAllRequests(tx neo4j.Transaction) ([]string, error) {
	result, err := tx.Run("Match (r:Request) RETURN r.projectId AS id", nil)

	if err != nil {
		return nil, err
	}
	var ids []string
	for result.Next() {
		record := result.Record()
		if id, ok := record.Get("id"); ok {
			ids = append(ids, id.(string))
		}
	}
	return ids, nil
}
