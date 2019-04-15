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

// ListarProfessor Retorna lista total professores registrados.
func ListarProfessor(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	lista, err := model.ListarProfessor(db)
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

// BuscarProfessor Retorna um professor especifico.
func BuscarProfessor(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	params := mux.Vars(r)
	idPessoa, err := strconv.ParseInt(params["idPessoa"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	entidade, err := model.BuscarProfessor(db, idPessoa)
	switch {
	case err == sql.ErrNoRows:
		var descerro bytes.Buffer
		descerro.WriteString("ERROR: Nenhum registro encontrado para idPessoa=")
		descerro.WriteString(strconv.FormatInt(idPessoa, 10))
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

// InserirProfessor Insere um novo registro professor na base.
func InserirProfessor(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	idpessoa, err := model.NextIDPessoa(db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	params := mux.Vars(r)
	nome := params["nome"]
	numerocpf := params["numerocpf"]
	numerocelular := params["numerocelular"]
	cidade := params["cidade"]
	numerocep := params["numerocep"]
	endereco := params["endereco"]
	dataCadastro := util.StringToTime(params["datacadastro"])
	idcurso, err := strconv.ParseInt(params["idcurso"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	entidadeCurso, err := model.BuscarCurso(db, idcurso)
	switch {
	case err == sql.ErrNoRows:
		var descerro bytes.Buffer
		descerro.WriteString("ERROR: Nenhum registro encontrado para idcurso=")
		descerro.WriteString(strconv.FormatInt(idcurso, 10))
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

	var pessoa entity.Pessoa
	pessoa.New(idpessoa, nome, numerocpf, numerocelular, cidade, numerocep, endereco, dataCadastro)
	idpessoaRetorno, err := model.InserirPessoa(db, pessoa)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	var entidade entity.Professor
	entidade.New(idpessoaRetorno, entidadeCurso.ID)
	idprofessorRetorno, err := model.InserirProfessor(db, entidade)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
	json.NewEncoder(w).Encode(idprofessorRetorno)
}

// AtualizarProfessor Atualiza um registro professor da base.
func AtualizarProfessor(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	params := mux.Vars(r)
	idpessoa, err := strconv.ParseInt(params["idpessoa"], 10, 64)
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
	nome := params["nome"]
	numerocpf := params["numerocpf"]
	numerocelular := params["numerocelular"]
	cidade := params["cidade"]
	numerocep := params["numerocep"]
	endereco := params["endereco"]
	datacadastro := util.StringToTime(params["datacadastro"])

	var entidade entity.Professor
	entidade.New(idpessoa, idcurso)
	var entidadepessoa entity.Pessoa
	entidadepessoa.New(idpessoa, nome, numerocpf, numerocelular, cidade, numerocep, endereco, datacadastro)

	if err = model.AtualizarProfessor(db, entidade, entidadepessoa); err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
}

// RemoverProfessor Remove um registro professor da base.
func RemoverProfessor(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	if err := Remover(w, r, db, "professor", "id_pessoa", "idpessoa"); err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	if err := Remover(w, r, db, "pessoa", "id", "idpessoa"); err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
}
