package tool

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func GenID() int64 {
	return node.Generate().Int64()
}
