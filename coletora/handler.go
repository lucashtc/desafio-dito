package coletora

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type apiServer struct {
	Port string
}

// endpoint para receber json event de compra e salvar em banco
func (a apiServer) Get(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Falha interna", http.StatusInternalServerError)
		return
	}

	event := Event{}
	if err = json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Json invalido", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	con, err := Conn()
	if err != nil {
		log.Println(errors.Wrap(err, "falha ao conectar ao banco"))
		return
	}

	err = Insert(con, event)
	if err != nil {
		log.Println(errors.Wrap(err, "Falha no insert"))
		return
	}

	w.WriteHeader(200)
	io.WriteString(w, "Salvo")
}

func (a apiServer) autocompleteApi(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Falha interna", http.StatusInternalServerError)
		return
	}

	event := Event{}
	if err = json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Json invalido", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	// Não faz pesquisa para resultao em branco
	if event.Event == "" {
		empty := make(map[string]string)
		json.NewEncoder(w).Encode(empty)
		return
	}

	con, err := Conn()
	if err != nil {
		log.Println(errors.Wrap(err, "falha ao conectar ao banco"))
		return
	}

	Name, err := Get(con, event.Event)
	if err != nil {
		log.Println(errors.Wrap(err, "Falha ao obter resultado"))
		return
	}
	fmt.Println(Name)
	json.NewEncoder(w).Encode(Name)
	return
}

func (a apiServer) autocomplete(w http.ResponseWriter, r *http.Request) {
	page, err := ioutil.ReadFile("./pages/autocomplete.html")
	if err != nil {
		http.Error(w, "Falha ao encontrar pagina", http.StatusBadRequest)
	}
	io.WriteString(w, string(page))
}

// InitServer ...
func InitServer() {
	server := apiServer{}
	server.Port = ":8080"

	http.HandleFunc("/autocomplete_api", server.autocompleteApi)
	http.HandleFunc("/autocomplete", server.autocomplete)
	http.HandleFunc("/get", server.Get)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./pages"))))

	log.Fatal(http.ListenAndServe(server.Port, nil))
}
