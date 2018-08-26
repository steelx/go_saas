package engine

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody parses the request body into a struct
func ParseBody(body io.ReadCloser, result interface{}) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(result)
}

// Respond return an object with specific status as JSON
func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) error {
	if e, ok := data.(error); ok {
		var tmp = new(struct {
			Status string `json:"status"`
			Error  string `json:"error"`
		})
		tmp.Status = "error"
		tmp.Error = e.Error()
		data = tmp
	}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	logRequest(r, status)
	return nil
}
