package helper

import (
	"encoding/json"
	"net/http"
)

func JsonDecode(r *http.Request, payload any) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(payload)
	if err != nil {
		return err
	}
	return nil
}

func JsonEncode(w http.ResponseWriter, payload any) error {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(payload)
	if err != nil {
		return err
	}
	return nil
}
