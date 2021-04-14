package delivery

import (
	"bytes"
	"challenge.haraj.com.sa/kraicklist/services"
	"encoding/json"
	"log"
	"net/http"
)

type RecordHandler struct {
	servicesRecord services.RecordServices
}

func NewRecordHandler (services services.RecordServices) {
	handler := &RecordHandler{
		servicesRecord: services,
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/search", handler.GetSearch)
}

func (v RecordHandler) GetSearch(w http.ResponseWriter, r *http.Request) {
	// fetch query string from query params
	q := r.URL.Query().Get("q")
	log.Println(q)
	if len(q) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("missing search query in query params"))
		if err != nil {
			log.Println(err)
		}
		return
	}

	records, err := v.servicesRecord.Search(q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			log.Fatal(err2)
		}
		log.Fatal(err)
	}

	// output success response
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	err = encoder.Encode(records)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}
