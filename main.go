package main

import (
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	bc := NewBlockchain()
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			log.Panic(err)
		}
	}(bc.db)

	cli := CLI{bc}
	cli.Run()
}
