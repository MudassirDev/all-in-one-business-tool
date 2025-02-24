package Json

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err.Error())
	}

	type Response struct {
		Msg string `json:"message"`
	}

	RespondWithJson(w, code, Response{
		Msg: msg,
	})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	toSend, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{msg: 'json encoding failed'}"))
		fmt.Printf("json encoding failed while sending back response, err: %v\n", err.Error())
		return
	}

	w.WriteHeader(code)
	w.Write(toSend)
}
