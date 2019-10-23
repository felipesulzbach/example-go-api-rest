package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/_dev/exemplo-api-rest/model"
	"github.com/_dev/exemplo-api-rest/model/entity"
	"github.com/_dev/exemplo-api-rest/util"

	"github.com/gorilla/mux"
)

// ListarAluno Retorna list total alunos registrados.
func ListarAluno(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	list, err := model.ListarAluno(db)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	for _, item := range list {
		log.Println(item.ToString())
	}

	model.CloseDB(db)
	json.NewEncoder(w).Encode(list)
}

// BuscarAluno Retorna um aluno especifico.
func BuscarAluno(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	params := mux.Vars(r)
	idpessoa, err := strconv.ParseInt(params["idpessoa"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	entityy, err := model.BuscarAluno(db, idpessoa)
	switch {
	case err == sql.ErrNoRows:
		var errorDesc bytes.Buffer
		errorDesc.WriteString("ERROR: No records found for id=")
		errorDesc.WriteString(strconv.FormatInt(idpessoa, 10))
		log.Println(errorDesc.String())
		json.NewEncoder(w).Encode(errorDesc.String())
		model.CloseDB(db)
		return
	case err != nil:
		log.Panic(err)
		model.CloseDB(db)
		return
	default:
	}

	log.Println(entityy.ToString())
	model.CloseDB(db)
	json.NewEncoder(w).Encode(entityy)
}

// InserirAluno Insere um novo registro aluno na base.
func InserirAluno(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	idpessoa, err := model.NextIDPessoa(db)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	entityyTurma, err := model.BuscarTurma(db, idturma)
	switch {
	case err == sql.ErrNoRows:
		var errorDesc bytes.Buffer
		errorDesc.WriteString("ERROR: No records found for idturma=")
		errorDesc.WriteString(strconv.FormatInt(idturma, 10))
		log.Println(errorDesc.String())
		json.NewEncoder(w).Encode(errorDesc.String())
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
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	var entityy entity.Aluno
	entityy.New(idpessoaRetorno, entityyTurma.ID)
	idalunoRetorno, err := model.InserirAluno(db, entityy)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	params := mux.Vars(r)
	idpessoa, err := strconv.ParseInt(params["idpessoa"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	idturma, err := strconv.ParseInt(params["idturma"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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

	var entityy entity.Aluno
	entityy.New(idpessoa, idturma)
	var entityypessoa entity.Pessoa
	entityypessoa.New(idpessoa, nome, numerocpf, numerocelular, cidade, numerocep, endereco, datacadastro)

	if err = model.AtualizarAluno(db, entityy, entityypessoa); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	if err := Remover(w, r, db, "aluno", "id_pessoa", "idpessoa"); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	if err := Remover(w, r, db, "pessoa", "id", "idpessoa"); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
}
