package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"exemplo-api-rest/model"
	"exemplo-api-rest/model/entity"
	"exemplo-api-rest/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ListarTurma Retorna lista total turmas registradas.
func ListarTurma(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	//	env := &Env{db}

	//	lista, err := env.db.BuscarCursos()
	lista, err := model.ListarTurma(db)
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

// BuscarTurma Retorna um curso especifico.
func BuscarTurma(w http.ResponseWriter, r *http.Request) {
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

	entidade, err := model.BuscarTurma(db, id)
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

// InserirTurma Insere um novo registro curso na base.
func InserirTurma(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	id, err := model.NextIDTurma(db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	params := mux.Vars(r)
	idcurso, err := strconv.ParseInt(params["idcurso"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	datainicio := util.StringToTime(params["datainicio"])
	datafim := util.StringToTime(params["datafim"])
	datacadastro := util.StringToTime(params["datacadastro"])

	var entidade entity.Turma
	entidade.New(id, idcurso, datainicio, datafim, datacadastro)

	idRetorno, err := model.InserirTurma(db, entidade)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
	json.NewEncoder(w).Encode(idRetorno)
}

// AtualizarTurma Atualiza um registro turma da base.
func AtualizarTurma(w http.ResponseWriter, r *http.Request) {
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
	idcurso, err := strconv.ParseInt(params["idcurso"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	datainicio := util.StringToTime(params["datainicio"])
	datafim := util.StringToTime(params["datafim"])
	datacadastro := util.StringToTime(params["datacadastro"])

	var entidade entity.Turma
	entidade.New(id, idcurso, datainicio, datafim, datacadastro)

	if err = model.AtualizarTurma(db, entidade); err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
}

// RemoverTurma Remove um registro turma da base.
func RemoverTurma(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	if err := Remover(w, r, db, "turma", "id", "id"); err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
}
