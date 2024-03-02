package main

import (
	"encoding/json"
	"log"

	simplejsondb "github.com/pnkj-kmr/simple-json-db"
)

var db simplejsondb.DB

func createDb(filename string) {
	var err error
	db, err = simplejsondb.New(filename, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("database created")
}

func initCollection(collectionName string) simplejsondb.Collection {
	collection, err := db.Collection(collectionName)
	if err != nil {
		log.Println("table err", err)
		return collection
	}
	log.Println("collection created")
	return collection
}

func addItem(collection simplejsondb.Collection, key string, data string) error {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("marshal err", err)
		return err
	}

	err = collection.Create(key, b)
	if err != nil {
		log.Println("create err", err)
		return err
	}

	log.Println("item added", data)
	return nil
}

func getAllItems(collection simplejsondb.Collection) []string {
	b := collection.GetAll()

	var items []string
	for _, v := range b {
		var item string
		err := json.Unmarshal(v, &item)
		if err != nil {
			log.Println("unmarshal err", err)
			return nil
		}
		items = append(items, item)
	}
	return items
}
