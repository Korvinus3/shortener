package storage

import (
	bolt "go.etcd.io/bbolt"
	"log"
	"ShortenerService/config"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

var DBInstance *Store

//Store type
type Store struct {
	Name         string
	client *bolt.DB
}

//NewBoltStorage create new DB instance
func NewBoltStorage() (*Store, error) {
	storage := new(Store)
	err := storage.open()
	return storage, err
}

func (st *Store) open() error {

	db, err := bolt.Open(config.Config.DataBase.Name, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	st.client = db

	return err

}

//Close current connection
func (st *Store) Close() {
	err := st.client.Close()
	if err != nil {
		log.Fatal(err)
	}
}

//Set value to DB
func (st *Store) Set(value string) (string, error) {

	var entryID string

	return entryID, st.client.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(config.Config.DataBase.URLsBucketName))

		id, _ := b.NextSequence()

		key := getHashKey(value, id)

		entryID = key

		return b.Put([]byte(entryID), []byte(value))
	})

}

//Get value from DB
func (st *Store) Get(key string) string {

	var val string

	st.client.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(config.Config.DataBase.URLsBucketName))

		val = string(b.Get([]byte(key)))

		return nil

	})

	return val

}

//NewBucket create new bucket in DB
func (st *Store) NewBucket(bucketName string) error{

	return st.client.Update(func(tx *bolt.Tx) error {

		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))

		return err
	})

}

func getHashKey(url string, key uint64) string {

	stringForHash := strconv.FormatUint(key, 10) + url

	hasher := md5.New()
	hasher.Write([]byte(stringForHash))
	hashedString := hex.EncodeToString(hasher.Sum(nil))

	return hashedString[:6]

}

//InitDB perform DB initializing
func InitDB() error {

	var err error

	DBInstance, err = NewBoltStorage()

	if err != nil {

		return err

	}

	DBInstance.NewBucket(config.Config.DataBase.URLsBucketName)

	return nil

}

//CloseDBConn close DB connection
func CloseDBConn()  {
	DBInstance.Close()
}