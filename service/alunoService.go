package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"rest-api/model"
	"rest-api/model/entity"
	"rest-api/util"

	"github.com/gorilla/mux"
)

// ListarAluno Retorna lista total alunos registrados.
func ListarAluno(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	lista, err := model.ListarAluno(db)
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

// BuscarAluno Retorna um aluno especifico.
func BuscarAluno(w http.ResponseWriter, r *http.Request) {
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

	entidade, err := model.BuscarAluno(db, idpessoa)
	switch {
	case err == sql.ErrNoRows:
		var descerro bytes.Buffer
		descerro.WriteString("ERROR: Nenhum registro encontrado para idpessoa=")
		descerro.WriteString(strconv.FormatInt(idpessoa, 10))
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

// InserirAluno Insere um novo registro aluno na base.
func InserirAluno(w http.ResponseWriter, r *http.Request) {
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
	idturma, err := strconv.ParseInt(params["idturma"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	entidadeTurma, err := model.BuscarTurma(db, idturma)
	switch {
	case err == sql.ErrNoRows:
		var descerro bytes.Buffer
		descerro.WriteString("ERROR: Nenhum registro encontrado para idturma=")
		descerro.WriteString(strconv.FormatInt(idturma, 10))
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

	var entidade entity.Aluno
	entidade.New(idpessoaRetorno, entidadeTurma.ID)
	idalunoRetorno, err := model.InserirAluno(db, entidade)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
	json.NewEncoder(w).Encode(idalunoRetorno)
}

// AtualizarAluno Atualiza um registro aluno da base.
func AtualizarAluno(w http.ResponseWriter, r *http.Request) {
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
	idturma, err := strconv.ParseInt(params["idturma"], 10, 64)
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

	var entidade entity.Aluno
	entidade.New(idpessoa, idturma)
	var entidadepessoa entity.Pessoa
	entidadepessoa.New(idpessoa, nome, numerocpf, numerocelular, cidade, numerocep, endereco, datacadastro)

	if err = model.AtualizarAluno(db, entidade, entidadepessoa); err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
}

// RemoverAluno Remove um registro aluno da base.
func RemoverAluno(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	if err := Remover(w, r, db, "aluno", "id_pessoa", "idpessoa"); err != nil {
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
