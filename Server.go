package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

type Person struct{
   Id        string    `json:"id"`
   FirstName string    `json:"first_name"`
   LastName  string    `json:"last_name"`
   Address   *Address  `json:"address"`
}
type Address struct{
    City string     `json:"city"`
    State string    `json:"state"`
}

type People []Person

var P1 People

func getPeople(w http.ResponseWriter, r *http.Request){

    json.NewEncoder(w).Encode(P1)

}

func addPeople(w http.ResponseWriter,r *http.Request){
    json.NewEncoder(w).Encode(P1)
    URL_map := mux.Vars(r)

    temp,_ := URL_map["Id"]
    Pnum := Person{Id:temp,FirstName:"Karuna",LastName:"Muthu",Address: &Address{City:"Theni",State:"Tamilnadu"}}

     P1 = append(P1,Pnum)

    json.NewEncoder(w).Encode(P1)
}

func delPeople(w http.ResponseWriter,r *http.Request){
    temp_Map := mux.Vars(r)
    var P2 People
    temp,_ := temp_Map["Id"]
    for index,value:=range P1{
       if temp == value.Id{
           P2 = append(P1[:index],P1[index+1:]...)
       }
    }
    json.NewEncoder(w).Encode(P2)
}

func main(){

      router := mux.NewRouter()

      P1 = []Person{
                    {"1","sundar","muni",&Address{"Theni","Tamilnadu"}},
                    {"2","Raji","muni",&Address{"Theni","Tamilnadu"}},
                    {"3","Indra","muni",&Address{"Theni","Tamilnadu"}},
                    {"4","Sanjay","Ramasamy",&Address{"Theni","Tamilnadu"}},
                   }


      router.HandleFunc("/people",getPeople).Methods("GET")
      router.HandleFunc("/person/{id}",addPeople).Methods("POST")
      router.HandleFunc("/person/{id}",delPeople).Methods("DELETE")

      err := http.ListenAndServe(":7070",router)

      if err != nil{
         log.Fatalln(err)
      }else{
         fmt.Println("Server listening on port 7070")
      }
}
