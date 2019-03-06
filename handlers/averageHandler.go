package handlers

import (
	"fmt"
	"net/http"
)

//AverageHandlerGet - гет запрос с страничке /average, возвращает html клиетну
func AverageHandlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "qweqweqw")
}
