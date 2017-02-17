package model

import (
	"log"
)

type Node struct {
	NodeId int
	ApiKeyId int
}

func (db *DB) FindNodeByApiKeyId(apiKeyId int)(Node, error) {
	result, err := db.Query(`select * from nodeId where apiKeyId = ?`,
		apiKeyId)

	if err != nil {
		return nil, err
	}

	var node Node

	defer result.Close()

	if result.Next() {
		result.Scan(&node)

		log.Println("find: %+v", node)
	}

	return node,nil
}