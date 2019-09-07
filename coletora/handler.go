package coletora

import (
	"net/http"
	"log"
	"io/ioutil"
	"io"
	"encoding/json"
	"github.com/pkg/errors"
)

type apiServer struct {
	Port string
}


// endpoint para receber json event de compra e salvar em banco 
func (a apiServer) Get(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w,"Falha interna",http.StatusInternalServerError)
		return
	}

	event := Event{}
	if err = json.Unmarshal(body, &event); err != nil{
		http.Error(w,"Json invalido",http.StatusInternalServerError)
		log.Println(err)
		return
	}
	con,err := Conn()
	if err != nil {
		log.Println(errors.Wrap(err,"falha ao conectar ao banco"))
		return
	}

	err = Insert(con,event)
	if err != nil {
		log.Println(errors.Wrap(err,"Falha no insert"))
		return
	}

	w.WriteHeader(200)
	io.WriteString(w,"Salvo")
}


func (a apiServer) autocomplete(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w,"Falha interna",http.StatusInternalServerError)
		return
	}

	event := Event{}
	if err = json.Unmarshal(body, &event); err != nil{
		http.Error(w,"Json invalido",http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	if event.Name == "" {
		empty := make(map[string]string)
		json.NewEncoder(w).Encode(empty)
		return
	}

	con,err := Conn()
	if err != nil {
		log.Println(errors.Wrap(err,"falha ao conectar ao banco"))
		return
	}

	Name, err := Get(con,event.Name)
	if err != nil {
		log.Println(errors.Wrap(err,"Falha ao obter resultado"))
		return
	}
	json.NewEncoder(w).Encode(Name)
	return
}


// InitServer ...
func InitServer(){
	server := apiServer{}
	server.Port = ":8080"


	http.HandleFunc("/autocomplete",server.autocomplete)
	http.HandleFunc("/get",server.Get)

	log.Fatal(http.ListenAndServe(server.Port,nil))
}