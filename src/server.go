package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Server struct {
	db   *gorm.DB
	form *Hash
}

type Hash struct {
	Origin string `json:"origin"`
	Value  string `json:"value"`
}

func (s *Server) setup() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to test.db")
	}
	s.db = db
}

func (s *Server) teardown() {
	s.db.Close()
}

func (s *Server) createForm(r *http.Request){
	swtich r.Header["Content-Type"] {
	case "application/x-www-form-urlencoded":
		r.ParseForm()
		
	}
	
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(r.RequestURI) > 1 {
		s.Get(w, r)
	} else {
		s.Post(w, r)
	}
}

func EncodeString(url string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(url))
	neededlength := len(sEnc) / 2
	return sEnc[:neededlength]
}

func (s *Server) Post(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Header["Content-Type"])
	r.ParseForm()
	f := &r.Form

	w.Header().Set("Content-Type", "application/json")

	url := (*f)["url"][0]
	h := Hash{
		Value:  EncodeString(url),
		Origin: url,
	}
	s.db.Create(&Url{Address: url, Hash: h.Value})
	response, _ := json.MarshalIndent(
		h,
		"",
		"  ",
	)
	w.Write(response)
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	hashvalue := r.RequestURI[1:]
	url := Url{}

	s.db.Where("hash = ?", hashvalue).First(&url)
	response, _ := json.MarshalIndent(
		&Hash{
			Value:  url.Hash,
			Origin: url.Address,
		},
		"",
		"  ",
	)
	w.Write(response)
}
