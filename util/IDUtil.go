package util

import (
	"github.com/bwmarrin/snowflake"
)

var Node *snowflake.Node

func Init(workerID int64) error {

	node, err := snowflake.NewNode(workerID)

	if err != nil {
		return err
	}

	Node = node

	return nil
}

func NextID() int64 {
	return Node.Generate().Int64()
}
