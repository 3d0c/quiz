package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Responder interface {
	Encode(v interface{}) []byte
	Write(w http.ResponseWriter)
}

type jsonResponder struct {
	w http.ResponseWriter
}

func newJsonResponder(w http.ResponseWriter) jsonResponder {
	return jsonResponder{w: w}
}

func (j jsonResponder) Encode(v interface{}) ([]byte, error) {
	var (
		b   []byte
		err error
	)

	if b, err = json.MarshalIndent(v, "", "    "); err != nil {
		return nil, err
	}

	return b, nil
}

func (j jsonResponder) Write(v interface{}) {
	var (
		b   []byte
		err error
	)

	if b, err = j.Encode(v); err != nil {
		http.Error(j.w, "", http.StatusInternalServerError)
		return
	}

	if _, err = j.w.Write(b); err != nil {
		log.Printf("Error writing response - %s\n", err)
	}

	return
}

func decode(r io.ReadCloser, into interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	r.Close()

	if err = json.Unmarshal(b, into); err != nil {
		return err
	}

	return nil
}
