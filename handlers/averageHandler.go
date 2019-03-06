package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//AverageHandlerGet - гет запрос с страничке /average, возвращает html клиетну
func AverageHandlerGet(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./static/average.html")
	tmpl.ExecuteTemplate(w, "average", nil)
}

//AverageHandlerPost - take a time row from client
func AverageHandlerPost(w http.ResponseWriter, r *http.Request) {
	gettingTimeRow := r.FormValue("sendedData")
	type ObjectMessage struct {
		TimeRow       []string
		TimeRowLength int
	}
	oneMessage := ObjectMessage{}
	err := json.Unmarshal([]byte(gettingTimeRow), &oneMessage)

	if err != nil {
		log.Println(err)
	}

	timeRowNumbers := make([]float64, oneMessage.TimeRowLength)
	for i := 0; i < len(timeRowNumbers); i++ {
		timeRowNumbers[i], _ = strconv.ParseFloat(oneMessage.TimeRow[i], 64)
	}

	fmt.Println(timeRowNumbers, GetMiddleValue(timeRowNumbers))
}

func GetMiddleValue(array []float64) float64 {
	var number float64
	for _, v := range array {
		number += v
	}
	number /= float64(len(array))
	return number
}
