package handlers

import (
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
	fmt.Println(r.FormValue("sendedData"))
}
