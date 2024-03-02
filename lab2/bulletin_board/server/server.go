// bulletin_board
// commands:
// LIST - return all ads
// ; - add ad
//
// requirements:
// save to external file or database and then return client message "Message added: “S”"
// if client enter empty string - close connection

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	simplejsondb "github.com/pnkj-kmr/simple-json-db"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type bulletinServer struct {
	serveMux http.ServeMux
}

type Answer struct {
	AnswerType string
	Value      interface{}
}

func newBulletinServer() *bulletinServer {
	cs := &bulletinServer{}
	cs.serveMux.Handle("/", http.FileServer(http.Dir(".")))
	cs.serveMux.HandleFunc("/bulletin_board", cs.BulletinHandler)

	return cs
}

func (cs *bulletinServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cs.serveMux.ServeHTTP(w, r)
}

func (cs *bulletinServer) BulletinHandler(w http.ResponseWriter, r *http.Request) {
	opts := websocket.AcceptOptions{
		OriginPatterns: []string{"localhost:5175"},
	}

	conn, err := websocket.Accept(w, r, &opts)
	log.Println("connection accepted")
	if err != nil {
		panic(err)
	}

	defer conn.Close(http.StatusInternalServerError, "the sky is falling")

	var item string
	var answer Answer
	for {
		err = wsjson.Read(context.Background(), conn, &item)
		if err != nil {
			panic(err)
		}

		if item == "LIST" {
			var items_collection simplejsondb.Collection
			items_collection, err = db.Collection("items")
			if err != nil {
				log.Println("table err", err)
			}
			items_string := getAllItems(items_collection)

			var answer Answer
			answer.AnswerType = "LIST"
			answer.Value = items_string
			wsjson.Write(r.Context(), conn, answer)

			continue
		} else if item == "" {
			// close connection

			answer.AnswerType = "MESSAGE"
			answer.Value = "Connection closed"
			wsjson.Write(r.Context(), conn, answer)
			return
		} else {
			var items_collection simplejsondb.Collection
			items_collection, err = db.Collection("items")
			if err != nil {
				log.Println("table err", err)
			}
			addItem(items_collection, uuid.NewString(), item)

			answer.AnswerType = "MESSAGE"
			answer.Value = "Message added: " + item
			wsjson.Write(r.Context(), conn, answer)

			continue
		}
	}
}
