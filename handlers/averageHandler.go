package handlers

import (
	"encoding/json" //package for serialize and deserialize json to/from object
	//package for writning some message on sreen
	"html/template" //package for parsing html
	"log"
	"net/http" //package for working with http requests
	//package for converting one data type to other data type
)

//AverageHandlerGet - гет запрос с страничке /average, возвращает html клиетну
func AverageHandlerGet(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./static/average.html") //Parsing html template
	tmpl.ExecuteTemplate(w, "average", nil)                 //Send this one to client
}

//MiddleSliceHandler - take a time row from client
func MiddleSliceHandler(w http.ResponseWriter, r *http.Request) {
	//Take data from client (ajax send it in json fromat)
	gettingTimeRow := r.FormValue("sendedData")
	//Create a struct (like a class in java)
	type MiddleSlice struct {
		TimeRow       []string
		SmoothK       float64
		TimeRowLength int
	}
	//Create an object of struct
	oneMessage := MiddleSlice{}
	//Unmarshaling json string to object. return error if there is something wrong
	//Return <nil> if all ok
	err := json.Unmarshal([]byte(gettingTimeRow), &oneMessage)

	//It will print error if error != nil
	if err != nil {
		log.Println(err)
	}

	//w.Write(middleValueAndTimeRowJSON)
}

//GetMiddleValue - return a middle value from array of []float64 (context - time row)
func GetMiddleValue(array []float64) float64 {
	//Create a variable of float64 (it will be middle value)
	var number float64
	//This a loop like foreach(...) in C#
	//renge return index and value of each element from time row
	//here number will had a sum of all element
	for _, v := range array {
		number += v
	}
	//here we will divide a sum on count
	number /= float64(len(array))
	return number
}
