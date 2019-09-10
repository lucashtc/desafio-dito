package autocomplete

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

// APIServer ...
type APIServer struct {
	Port string
}

// Get endpoint para receber json event de compra e salvar em banco
func (a APIServer) Get(w http.ResponseWriter, r *http.Request) {
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

// AutocompleteAPI ...
func (a APIServer) AutocompleteAPI(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Falha interna", http.StatusInternalServerError)
		return
	}

	event := Event{}
	if err = json.Unmarshal(body, &event); err != nil {
		errString := fmt.Sprint(errors.Wrap(err, "falha"))
		http.Error(w, errString, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	// Não faz pesquisa para resultado em branco
	if event.Event == "" {
		empty := make(map[string]string)
		json.NewEncoder(w).Encode(empty)
		return
	}

	con, err := Conn()
	if err != nil {
		errString := fmt.Sprint(errors.Wrap(err, "falha ao conectar ao banco"))
		http.Error(w, errString, http.StatusInternalServerError)
		return
	}

	Name, err := Get(con, event.Event)
	if err != nil {
		errString := fmt.Sprint(errors.Wrap(err, "Falha ao obter resultado"))
		http.Error(w, errString, http.StatusInternalServerError)
		return
	}
	fmt.Println(Name)
	json.NewEncoder(w).Encode(Name)
	return
}

// Autocomplete ...
func (a APIServer) Autocomplete(w http.ResponseWriter, r *http.Request) {
	page, err := ioutil.ReadFile("./pages/autocomplete.html")
	if err != nil {
		http.Error(w, "Falha ao encontrar pagina", http.StatusBadRequest)
	}
	io.WriteString(w, string(page))
}

// Migration ...
func (a APIServer) Migration(w http.ResponseWriter, r *http.Request) {
	con, err := Conn()
	if err != nil {
		errString := fmt.Sprint(errors.Wrap(err, "falha ao conectar ao banco"))
		http.Error(w, errString, http.StatusInternalServerError)
		return
	}

	err = MigrateCreateTable(con)
	if err != nil {
		errString := fmt.Sprint(errors.Wrap(err, "Falha ao criar tabela"))
		http.Error(w, errString, http.StatusInternalServerError)
		return
	}

	err = InsertMake(con)
	if err != nil {
		errString := fmt.Sprint(errors.Wrap(err, "Falha ao alimentar tabela"))
		http.Error(w, errString, http.StatusInternalServerError)
		return
	}
	fmt.Println("Deu Bom")
	return
}
