package dbs

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

const dbFile = "bolt.db"
const keystore = "keystore"

var dbptr *bolt.DB

func init() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	dbptr = db
	fmt.Println("init db ok:", dbptr.Info().PageSize)
}

// 判断db是否已经存在
func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}

func AddData(key []byte, value []byte) {
	fmt.Println("add:", dbptr.String())
	dbptr.Update(func(tx *bolt.Tx) error {
		//2.1 获取bucket
		buck := tx.Bucket([]byte(keystore))
		if buck == nil {

			bucket, _ := tx.CreateBucket([]byte(keystore))
			bucket.Put(key, value)

		} else {
			buck.Put(key, value)
		}
		return nil
	})
}

func Query(key []byte) []byte {
	var val []byte
	fmt.Println("Query:", dbptr.String())
	dbptr.View(func(tx *bolt.Tx) error {
		//2.1 获取bucket
		buck := tx.Bucket([]byte(keystore))
		if buck == nil {
			return nil
		}
		val = buck.Get(key)
		return nil
	})

	return val
}
