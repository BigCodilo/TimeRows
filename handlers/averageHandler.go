package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	type objectMessage struct {
		timeRow       []float64
		timeRowLength int
	}

	oneMessage := objectMessage{}

	json.Unmarshal([]byte(gettingTimeRow), &oneMessage)

	fmt.Println(oneMessage)
}
