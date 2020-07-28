package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/gorilla/mux"
)

type Service struct {
	database *firestore.Client
	pubsub   *pubsub.Client
}

func Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	ctx := context.Background()
	project := os.Getenv("GOOGLE_PROJECT")
	client, err := firestore.NewClient(ctx, project)
	if err != nil {
		log.Fatal(err)
	}
	pubsubClient, err := pubsub.NewClient(ctx, project)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	service := &Service{database: client, pubsub: pubsubClient}

	r := mux.NewRouter()
	r.HandleFunc("/", service.HomeHandler)
	r.HandleFunc("/request", service.RequestHandler)
	r.HandleFunc("/event", service.EventHandler)

	http.Handle("/", r)

	log.Println("Listening on port " + Port())
	log.Fatal(http.ListenAndServe(":"+Port(), nil))
}

type Request struct {
	URI string `json:"uri"`
}

func (s *Service) HomeHandler(w http.ResponseWriter, r *http.Request) {
	iter := s.database.
		Collection("requests").
		Documents(r.Context())

	snaps, err := iter.GetAll()
	if err != nil {
		sendErr(w, err)
		return
	}

	defer iter.Stop()
	var requests []Request
	for _, ds := range snaps {
		var r Request
		ds.DataTo(&r)
		requests = append(requests, r)
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(requests)
	if err != nil {
		sendErr(w, err)
		return
	}
}

func (s *Service) RequestHandler(w http.ResponseWriter, r *http.Request) {
	request := Request{URI: "www.nos.nl"}
	doc, _, err := s.database.
		Collection("requests").
		Add(context.Background(), request)
	if err != nil {
		sendErr(w, err)
	}
	json, err := json.Marshal(doc)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	s.pubsub.Topic("requests").Publish(r.Context(), &pubsub.Message{Data: json})
	log.Printf("Created request %s", doc.ID)

}

func (s *Service) EventHandler(w http.ResponseWriter, r *http.Request) {
	request, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, string(request))
		log.Print(string(request))
	}
}

func sendErr(w http.ResponseWriter, err error) {
	http.Error(w, "Error", http.StatusInternalServerError)
	fmt.Println(err)
}
