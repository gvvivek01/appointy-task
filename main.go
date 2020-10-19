package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Meeting structure
type Meeting struct {
	ID                 string    `josn: "Id"`
	Title              string    `json: Title"`
	Participants       string    `json: Participants"`
	Start_Time         time.Time `json: Start_Time"`
	End_Time           time.Time `json: End_Time"`
	Creation_Timestamp time.Time `json: Created"`
}

type Participant struct {
	Name  string `json:"name" validate:"required,min=2,max=100" required:"true"`
	Email string `json:"email" validate:"required,email" required:"true"`
	RSVP  string `jsom:"RSVP" required:"true"`
}

// array of meetings
type Meetings []Meeting

func allMeetings(w http.ResponseWriter, r *http.Request) {
	meetings := Meetings{
		Meeting{
			ID:                 "1",
			Title:              "title 1",
			Participants:       "2",
			Start_Time:         time.Date(2020, 10, 14, 10, 45, 0, 0, time.UTC),
			End_Time:           time.Date(2020, 10, 14, 11, 00, 16, 0, time.UTC),
			Creation_Timestamp: time.Date(2020, 10, 14, 10, 45, 0, 0, time.UTC)},
		Meeting{
			ID:                 "2",
			Title:              "title 2",
			Participants:       "2",
			Start_Time:         time.Date(2020, 10, 14, 11, 45, 0, 0, time.UTC),
			End_Time:           time.Date(2020, 10, 14, 12, 15, 16, 0, time.UTC),
			Creation_Timestamp: time.Date(2020, 10, 14, 10, 45, 0, 0, time.UTC)},
		Meeting{
			ID:                 "3",
			Title:              "title 3",
			Participants:       "2",
			Start_Time:         time.Date(2020, 10, 14, 9, 15, 0, 0, time.UTC),
			End_Time:           time.Date(2020, 10, 14, 9, 45, 16, 0, time.UTC),
			Creation_Timestamp: time.Date(2020, 10, 14, 10, 45, 0, 0, time.UTC)},
	}

	fmt.Println("Endpoint Hit: All meetings Endpoint")
	json.NewEncoder(w).Encode(meetings)
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	// Set the return Content-Type as JSON
	w.Header().Set("Content-Type", "application/json")

	// Change the response depending on the method being requested
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET method requested"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	//http.HandleFunc("/", homePage)
	//register func here
	http.HandleFunc("/meetings", allMeetings)

	http.HandleFunc("/", apiResponse)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

//entry point to our application
func main() {

	// Replace the uri string with your MongoDB deployment's connection string.
	uri := <your uri> 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//To close the connection at the end
	defer cancel()
	//We need to set up a client first
	//It takes the URI of your database
	client, error := mongo.NewClient(options.Client().ApplyURI(uri))
	if error != nil {
		log.Fatal(error)
	}
	//Call the connect function of client
	error = client.Connect(ctx)
	//Checking the connection
	error = client.Ping(context.TODO(), nil)
	fmt.Println("Database connected")

	handleRequests()

}
