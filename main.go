package main
import (
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "math/rand"
  "strconv"
)
type Username struct {
  ID string `json:"id"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
}

var usernames []Username

func main() {
	router := mux.NewRouter()
	usernames = append(usernames, Username{ID: "1", FirstName: "Nanda", LastName: "Prasetyo"})
	router.HandleFunc("/username-list", getUsernames).Methods("GET")
	router.HandleFunc("/create-username", createUsername).Methods("POST")
	router.HandleFunc("/get-username/{id}", getUsername).Methods("GET")
	router.HandleFunc("/update-username/{id}", updateUsername).Methods("PUT")
	router.HandleFunc("/delete-username/{id}", deleteUsername).Methods("DELETE")
	http.ListenAndServe(":8008", router)
}

func getUsernames(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(usernames)
}

func createUsername(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var post Username
  _ = json.NewDecoder(r.Body).Decode(&post)
  
  post.ID = strconv.Itoa(rand.Intn(1000000))
  usernames = append(usernames, post)
  json.NewEncoder(w).Encode(&post)
}

func getUsername(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for _, item := range usernames {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Username{})
}

func updateUsername(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range usernames {
    if item.ID == params["id"] {
      usernames = append(usernames[:index], usernames[index+1:]...)
      var post Username
      _ = json.NewDecoder(r.Body).Decode(&post)
      post.ID = params["id"]
      usernames = append(usernames, post)
      json.NewEncoder(w).Encode(&post)
      return
    }
  }
  json.NewEncoder(w).Encode(usernames)
}

func deleteUsername(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range usernames {
    if item.ID == params["id"] {
      usernames = append(usernames[:index], usernames[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(usernames)
}