package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"
)

//article structure
// type Article struct {
//     Title string `json:"Title"`
//     Desc string `json:"desc"`
//     Content string `json:"content"`
// }

type Meeting struct {
	Id string `josn: "Id" required:"true"`
	Title string `json: Title" required:"true"`
	Participants string `json: Participants" required:"true"`
	Start_Time time.Time `json: Start_Time" required:"true"` 
	End_Time time.Time `json: End_Time" required:"true"`
	Creation_Timestamp time.Time `json: Created" required:"true"`

}

type Participant struct{
	Name  string `json:"name" validate:"required,min=2,max=100" required:"true"`
	Email string `json:"email" validate:"required,email" required:"true"`
	RSVP string `jsom:"RSVP" required:"true"`

}
// array of articles
// type Articles []Article

type Meetings []Meeting


// func allArticles(w http.ResponseWriter, r *http.Request){
// 	articles := Articles{
// 		Article{Title: "Test Title", Desc: "Test Description", Content: "Article Content"},
// 	}


// 	fmt.Println("Endpoint Hit: All articles Endpoint")
// 	json.NewEncoder(w).Encode(articles)
// }


func allMeetings(w http.ResponseWriter, r *http.Request){
	meetings := Meetings{
		Meeting{
			Id: "1", 
			Title: "title 1", 
			Participants: "2",
			Start_Time: time.Date(2020, 10, 14, 10, 45, 0, 0, time.UTC), 
			End_Time: time.Date(2020, 10, 14, 11, 00, 16, 0, time.UTC),
			Creation_Timestamp: time.Date(2020, 10, 14, 10, 45, 0, 0, time.UTC)},
		Meeting{
			Id: "2", 
			Title: "title 2", 
			Participants: "2", 
			Start_Time: time.Date(2020, 10, 14, 11, 45, 0, 0, time.UTC), 
			End_Time: time.Date(2020, 10, 14, 12, 15, 16, 0, time.UTC), 
			Creation_Timestamp: time.Date(2020, 10, 14, 10, 45, 0, 0, time.UTC)},
		Meeting{
			Id: "3", 
			Title: "title 3", 
			Participants: "2", 
			Start_Time: time.Date(2020, 10, 14, 9, 15, 0, 0, time.UTC), 
			End_Time: time.Date(2020, 10, 14, 9, 45, 16, 0, time.UTC), 
			Creation_Timestamp: time.Date(2020, 10, 14, 10, 45, 0, 0, time.UTC)},
	}


	fmt.Println("Endpoint Hit: All meetings Endpoint")
	json.NewEncoder(w).Encode(meetings)
}

// func (t Time) Clock() (hour, min, sec int) {
// 	return absClock(t.abs())
// }

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests(){
	http.HandleFunc("/",homePage)
	//register func here
	http.HandleFunc("/meetings", allMeetings)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

//entry point to our application
func main(){
	handleRequests()
}