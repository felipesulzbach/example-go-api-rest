package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"exemplo-api-rest/model"
	"exemplo-api-rest/model/entity"
	"exemplo-api-rest/util"

	"github.com/gorilla/mux"
)

// ListarCurso Retorna lista total cursos registrados.
func ListarCurso(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	//	env := &Env{db}

	//	lista, err := env.db.ListarCurso()
	lista, err := model.ListarCurso(db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	for _, item := range lista {
		log.Println(item.ToString())
	}

	model.CloseDB(db)
	json.NewEncoder(w).Encode(lista)
}

// BuscarCurso Retorna um curso especifico.
func BuscarCurso(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	params := mux.Vars(r)
	//id := r.FormValue("id")
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	entidade, err := model.BuscarCurso(db, id)
	switch {
	case err == sql.ErrNoRows:
		var descerro bytes.Buffer
		descerro.WriteString("ERROR: Nenhum registro encontrado para id=")
		descerro.WriteString(strconv.FormatInt(id, 10))
		log.Println(descerro.String())
		json.NewEncoder(w).Encode(descerro.String())
		model.CloseDB(db)
		return
	case err != nil:
		log.Panic(err)
		model.CloseDB(db)
		return
	default:
	}

	log.Println(entidade.ToString())
	model.CloseDB(db)
	json.NewEncoder(w).Encode(entidade)
}

// InserirCurso Insere um novo registro curso na base.
func InserirCurso(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	id, err := model.NextIDCurso(db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	params := mux.Vars(r)
	nome := params["nome"]
	descricao := params["descricao"]
	dataCadastro := util.StringToTime(params["datacadastro"])

	var entidade entity.Curso
	entidade.New(id, nome, descricao, dataCadastro)

	idRetorno, err := model.InserirCurso(db, entidade)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
	json.NewEncoder(w).Encode(idRetorno)
}

// AtualizarCurso Atualiza um registro curso da base.
func AtualizarCurso(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	nome := params["nome"]
	descricao := params["descricao"]
	dataCadastro := util.StringToTime(params["datacadastro"])

	var entidade entity.Curso
	entidade.New(id, nome, descricao, dataCadastro)

	if err = model.AtualizarCurso(db, entidade); err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
}

// RemoverCurso Remove um registro curso da base.
func RemoverCurso(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	if err := Remover(w, r, db, "curso", "id", "id"); err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
}
