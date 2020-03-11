package utils

import (
	"github.com/bwmarrin/snowflake"
	"log"
	"sync"
)

var mux sync.Mutex
var defaultNode *snowflake.Node

func init() {
	var err error
	defaultNode, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}
}

func GenBase32() string {
	mux.Lock()
	defer mux.Unlock()

	id := defaultNode.Generate()
	return id.Base32()

}
