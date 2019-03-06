package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//AverageHandlerGet - гет запрос с страничке /average, возвращает html клиетну
func AverageHandlerGet(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./static/average.html")
	tmpl.ExecuteTemplate(w, "average", nil)
}

//AverageHandlerPost - take a time row from client
func AverageHandlerPost(w http.ResponseWriter, r *http.Request) {
	gettingTimeRow := r.FormValue("sendedData")
	fmt.Println(gettingTimeRow)
	type ObjectMessage struct {
		TimeRow       []string
		TimeRowLength string
	}

	oneMessage := ObjectMessage{}

	err := json.Unmarshal([]byte(gettingTimeRow), &oneMessage)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(oneMessage)
}
