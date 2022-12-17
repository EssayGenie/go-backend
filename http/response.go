package http

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	b, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error marshalling JSON: %v", err))
	}
	_, err = w.Write(b)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error writing JSON: %v", err))
	}
	return nil
}
