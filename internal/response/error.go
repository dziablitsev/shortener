package response

import "net/http"

func BadRequest(res http.ResponseWriter) {
	http.Error(res, "Request is not allowed!", http.StatusBadRequest)
}
