package Controllers

import "net/http"

type BaseController struct {
}

func (this *BaseController) returnServerError(err error, response http.ResponseWriter) {
	response.WriteHeader(http.StatusInternalServerError)
	_, _ = response.Write([]byte(`{"message": "` + err.Error() + `"}`))
}

func (this *BaseController) returnNotFound(response http.ResponseWriter) {
	response.WriteHeader(http.StatusNotFound)
	_, _ = response.Write([]byte(`{"message": "Item not found in database"}`))
}

func (this *BaseController) returnBadRequest(error string, response http.ResponseWriter) {
	response.WriteHeader(http.StatusBadRequest)
	_, _ = response.Write([]byte(`{"message": "` + error + `"}`))
}

func (this *BaseController) setResponseHeaders(response http.ResponseWriter) {
	response.Header().Add("content-type", "application/json")
}
