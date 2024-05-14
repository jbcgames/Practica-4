package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Superhero struct {
	Name       string     `json:"name"`
	Biography  Biography  `json:"biography"`
	Powerstats Powerstats `json:"powerstats"`
	Images     Images     `json:"images"`
}

type Biography struct {
	FullName string `json:"fullName"`
}

type Powerstats struct {
	Intelligence int `json:"intelligence"`
	Strength     int `json:"strength"`
	Speed        int `json:"speed"`
	Durability   int `json:"durability"`
	Power        int `json:"power"`
	Combat       int `json:"combat"`
}

type Images struct {
	Xs string `json:"xs"`
	Sm string `json:"sm"`
	Md string `json:"md"`
	Lg string `json:"lg"`
}

var superheroes []Superhero

func main() {
	router := mux.NewRouter()
	superheroes = append(superheroes, Superhero{Name: "Wolverine", Biography: Biography{FullName: "John Logan"}, Powerstats: Powerstats{Intelligence: 63, Strength: 32, Speed: 50, Durability: 100, Power: 89, Combat: 100}, Images: Images{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg"}})
	superheroes = append(superheroes, Superhero{Name: "Spider-Man", Biography: Biography{FullName: "Peter Parker"}, Powerstats: Powerstats{Intelligence: 90, Strength: 55, Speed: 67, Durability: 75, Power: 74, Combat: 85}, Images: Images{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg"}})
	superheroes = append(superheroes, Superhero{Name: "Iron Man", Biography: Biography{FullName: "Tony Stark"}, Powerstats: Powerstats{Intelligence: 100, Strength: 85, Speed: 58, Durability: 85, Power: 100, Combat: 64}, Images: Images{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg"}})
	superheroes = append(superheroes, Superhero{Name: "Black Widow", Biography: Biography{FullName: "Natasha Romanoff"}, Powerstats: Powerstats{Intelligence: 75, Strength: 13, Speed: 33, Durability: 30, Power: 36, Combat: 100}, Images: Images{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/107-black-widow.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/107-black-widow.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/107-black-widow.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/107-black-widow.jpg"}})
	superheroes = append(superheroes, Superhero{Name: "Thor", Biography: Biography{FullName: "Thor Odinson"}, Powerstats: Powerstats{Intelligence: 69, Strength: 100, Speed: 83, Durability: 100, Power: 100, Combat: 100}, Images: Images{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg"}})
	router.HandleFunc("/api/superhero", GetSuperhero).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func GetSuperhero(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	hero := params.Get("hero")
	for _, item := range superheroes {
		if item.Name == hero {
			json.NewEncoder(w).Encode(item)
			print(hero)
			return
		}
	}
	http.Error(w, "Superhero not found", http.StatusNotFound)
	print(hero)
}
