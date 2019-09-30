package main

import (
	"bytes"
	"context"
	"encoding/json"
	"exemplo-api-rest/model"
	"exemplo-api-rest/model/entity"
	"exemplo-api-rest/service"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	configurarServidor()
	//exemploDecoderJSON()
}

func configurarServidor() {
	log.Println("Server Start!")
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "a duracao para a qual o servidor normalmente espera que as conexoes existentes terminem - e.g. 15s ou 1m")
	flag.Parse()

	router := configurarRotas()

	srv := &http.Server{
		Addr:         "localhost:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Server Shutting Down!")
	os.Exit(0)
}

func configurarRotas() *mux.Router {
	router := mux.NewRouter()
	subrouterCurso := router.PathPrefix("/curso").Subrouter()
	subrouterCurso.Path("").HandlerFunc(service.ListarCurso).Methods("GET")
	subrouterCurso.Path("/").Queries("id", "{id}").HandlerFunc(service.BuscarCurso).Methods("GET").Name("BuscarCurso")
	subrouterCurso.Path("/").Queries("nome", "{nome}", "descricao", "{descricao}", "datacadastro", "{datacadastro}").HandlerFunc(service.InserirCurso).Methods("POST").Name("InserirCurso")
	subrouterCurso.Path("/").Queries("id", "{id}", "nome", "{nome}", "descricao", "{descricao}", "datacadastro", "{datacadastro}").HandlerFunc(service.AtualizarCurso).Methods("PUT").Name("AtualizarCurso")
	subrouterCurso.Path("/").Queries("id", "{id}").HandlerFunc(service.RemoverCurso).Methods("DELETE").Name("RemoverCurso")

	subrouterTurma := router.PathPrefix("/turma").Subrouter()
	subrouterTurma.Path("").HandlerFunc(service.ListarTurma).Methods("GET")
	subrouterTurma.Path("/").Queries("id", "{id}").HandlerFunc(service.BuscarTurma).Methods("GET").Name("BuscarTurma")
	subrouterTurma.Path("/").Queries("idcurso", "{idcurso}", "datainicio", "{datainicio}", "datafim", "{datafim}", "datacadastro", "{datacadastro}").HandlerFunc(service.InserirTurma).Methods("POST").Name("InserirTurma")
	subrouterTurma.Path("/").Queries("id", "{id}", "idcurso", "{idcurso}", "datainicio", "{datainicio}", "datafim", "{datafim}", "datacadastro", "{datacadastro}").HandlerFunc(service.AtualizarTurma).Methods("PUT").Name("AtualizarTurma")
	subrouterTurma.Path("/").Queries("id", "{id}").HandlerFunc(service.RemoverTurma).Methods("DELETE").Name("RemoverTurma")

	subrouterPessoa := router.PathPrefix("/pessoa").Subrouter()
	subrouterPessoa.Path("").HandlerFunc(service.ListarPessoa).Methods("GET")

	subrouterAluno := router.PathPrefix("/aluno").Subrouter()
	subrouterAluno.Path("").HandlerFunc(service.ListarAluno).Methods("GET")
	subrouterAluno.Path("/").Queries("idpessoa", "{idpessoa}").HandlerFunc(service.BuscarAluno).Methods("GET").Name("BuscarAluno")
	subrouterAluno.Path("/").Queries("idturma", "{idturma}", "nome", "{nome}", "numerocpf", "{numerocpf}", "numerocelular", "{numerocelular}", "cidade", "{cidade}", "numerocep", "{numerocep}", "endereco", "{endereco}", "datacadastro", "{datacadastro}").HandlerFunc(service.InserirAluno).Methods("POST").Name("InserirAluno")
	subrouterAluno.Path("/").Queries("idpessoa", "{idpessoa}", "idturma", "{idturma}", "nome", "{nome}", "numerocpf", "{numerocpf}", "numerocelular", "{numerocelular}", "cidade", "{cidade}", "numerocep", "{numerocep}", "endereco", "{endereco}", "datacadastro", "{datacadastro}").HandlerFunc(service.AtualizarAluno).Methods("PUT").Name("AtualizarAluno")
	subrouterAluno.Path("/").Queries("idpessoa", "{idpessoa}").HandlerFunc(service.RemoverAluno).Methods("DELETE").Name("RemoverAluno")

	subrouterProfessor := router.PathPrefix("/professor").Subrouter()
	subrouterProfessor.Path("").HandlerFunc(service.ListarProfessor).Methods("GET")
	subrouterProfessor.Path("/").Queries("idpessoa", "{idpessoa}").HandlerFunc(service.BuscarProfessor).Methods("GET").Name("BuscarProfessor")
	subrouterProfessor.Path("/").Queries("idcurso", "{idcurso}", "nome", "{nome}", "numerocpf", "{numerocpf}", "numerocelular", "{numerocelular}", "cidade", "{cidade}", "numerocep", "{numerocep}", "endereco", "{endereco}", "datacadastro", "{datacadastro}").HandlerFunc(service.InserirProfessor).Methods("POST").Name("InserirProfessor")
	subrouterProfessor.Path("/").Queries("idpessoa", "{idpessoa}", "idcurso", "{idcurso}", "nome", "{nome}", "numerocpf", "{numerocpf}", "numerocelular", "{numerocelular}", "cidade", "{cidade}", "numerocep", "{numerocep}", "endereco", "{endereco}", "datacadastro", "{datacadastro}").HandlerFunc(service.AtualizarProfessor).Methods("PUT").Name("AtualizarProfessor")
	subrouterProfessor.Path("/").Queries("idpessoa", "{idpessoa}").HandlerFunc(service.RemoverProfessor).Methods("DELETE").Name("RemoverProfessor")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8000", router))

	return router
}

func exemploDecoderJSON() {
	jsonStream := []byte(`{"ID":3,"Nome":"Nome 3","Descricao":"Descricao 3","DataCadastro":"2019-01-21T10:07:16.543807Z"}`)

	var curso *entity.Curso
	err := json.Unmarshal(jsonStream, &curso)
	if err != nil {
		fmt.Println("Erro:", err)
		log.Panic(err)
	}
	log.Println(curso.ToString())
}

/// >> teste performance string, buffer, builder ///
func testeConcat(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str += "x"
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("resultado inesperado; got=%s, want=%s", str, s)
	}
}

func testeBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("resultado inesperado; got=%s, want=%s", buffer.String(), s)
	}
}

func testeCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], "x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("resultado inesperado; got=%s, want=%s", string(bs), s)
	}
}

func testeStringBuilder(b *testing.B) {
	var strBuilder strings.Builder

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		strBuilder.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); strBuilder.String() != s {
		b.Errorf("resultado inesperado; got=%s, want=%s", strBuilder.String(), s)
	}
}

/// << teste performance string, buffer, builder ///

func testeCreateTable() {
	db, err := model.NewDB(service.DataSourcePostgre)
	if err != nil {
		log.Panic(err)
		return
	}

	stmt, err := db.Prepare("CREATE TABLE $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec("curso"); err != nil {
		log.Fatal(err)
	}
}
