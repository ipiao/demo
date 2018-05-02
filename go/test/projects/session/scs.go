package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"time"

	scs "github.com/ipiao/session"
	"github.com/ipiao/session/stores/mysqlstore"
)

var sessionManager = scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")

func main() {
	db, err := sql.Open("mysql", "root:1001@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	var store = mysqlstore.New(db, time.Minute*10)
	sessionManager = scs.NewManager(store)
	sessionManager.Option(scs.Persist(true))
	sessionManager.Option(scs.LifeTime(time.Second * 30))

	mux := http.NewServeMux()
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)

	http.ListenAndServe(":4000", mux)
}

func putHandler(w http.ResponseWriter, r *http.Request) {

	session, err := sessionManager.Load(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = session.PutToResponseWriter(w, "message", "Hello world!")
	// sessionManager.Write(session, w)
	// session.WriteToResponseWriter(w)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	log.Println("PUT:", session.GetData())
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessionManager.Load(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	session.WriteToResponseWriter(w)
	sessions := sessionManager.FindSeesion()
	log.Println("GET:", len(sessions))
	sessions1 := sessionManager.FindSeesion(scs.FindByKVEq("message", "Hello world!"))
	log.Println("GET KVEq:", len(sessions1))
	sessions2 := sessionManager.FindSeesion(scs.FindTimeOut())
	log.Println("GET TimeOut:", len(sessions2))
	message, err := session.GetString("message")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	log.Println("GET:", session.GetData())
	io.WriteString(w, message)
}
