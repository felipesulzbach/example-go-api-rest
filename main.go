package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"testing"
	"time"

	"github.com/_dev/exemplo-api-rest/model"
	"github.com/_dev/exemplo-api-rest/model/entity"
	"github.com/_dev/exemplo-api-rest/service"

	"github.com/gorilla/mux"
)

func main() {
	serverConfigure()
	//exampleDecoderJSON()
}

func serverConfigure() {
	log.Println("Server Start!")
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "he duration for which the server normally expects existing connections to end - e.g. 15s ou 1m")
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
	subrouterCurso := router.PathPrefix("/course").Subrouter()
	subrouterCurso.Path("").HandlerFunc(service.ListarCurso).Methods("GET")
	subrouterCurso.Path("/").Queries("id", "{id}").HandlerFunc(service.BuscarCurso).Methods("GET").Name("BuscarCurso")
	subrouterCurso.Path("/").Queries("name", "{name}", "description", "{description}", "registrationDate", "{registrationDate}").HandlerFunc(service.InserirCurso).Methods("POST").Name("InserirCurso")
	subrouterCurso.Path("/").Queries("id", "{id}", "name", "{name}", "description", "{description}", "registrationDate", "{registrationDate}").HandlerFunc(service.AtualizarCurso).Methods("PUT").Name("AtualizarCurso")
	subrouterCurso.Path("/").Queries("id", "{id}").HandlerFunc(service.RemoverCurso).Methods("DELETE").Name("RemoverCurso")

	subrouterTurma := router.PathPrefix("/class").Subrouter()
	subrouterTurma.Path("").HandlerFunc(service.ListarTurma).Methods("GET")
	subrouterTurma.Path("/").Queries("id", "{id}").HandlerFunc(service.BuscarTurma).Methods("GET").Name("BuscarTurma")
	subrouterTurma.Path("/").Queries("courseID", "{courseID}", "startDate", "{startDate}", "endDate", "{endDate}", "registrationDate", "{registrationDate}").HandlerFunc(service.InserirTurma).Methods("POST").Name("InserirTurma")
	subrouterTurma.Path("/").Queries("id", "{id}", "courseID", "{courseID}", "startDate", "{startDate}", "endDate", "{endDate}", "registrationDate", "{registrationDate}").HandlerFunc(service.AtualizarTurma).Methods("PUT").Name("AtualizarTurma")
	subrouterTurma.Path("/").Queries("id", "{id}").HandlerFunc(service.RemoverTurma).Methods("DELETE").Name("RemoverTurma")

	subrouterPessoa := router.PathPrefix("/person").Subrouter()
	subrouterPessoa.Path("").HandlerFunc(service.ListarPessoa).Methods("GET")

	subrouterAluno := router.PathPrefix("/student").Subrouter()
	subrouterAluno.Path("").HandlerFunc(service.ListarAluno).Methods("GET")
	subrouterAluno.Path("/").Queries("personID", "{personID}").HandlerFunc(service.BuscarAluno).Methods("GET").Name("BuscarAluno")
	subrouterAluno.Path("/").Queries("classID", "{classID}", "name", "{name}", "cpf", "{cpf}", "cellPhone", "{cellPhone}", "city", "{city}", "zipCode", "{zipCode}", "address", "{address}", "registrationDate", "{registrationDate}").HandlerFunc(service.InserirAluno).Methods("POST").Name("InserirAluno")
	subrouterAluno.Path("/").Queries("personID", "{personID}", "classID", "{classID}", "name", "{name}", "cpf", "{cpf}", "cellPhone", "{cellPhone}", "city", "{city}", "zipCode", "{zipCode}", "address", "{address}", "registrationDate", "{registrationDate}").HandlerFunc(service.AtualizarAluno).Methods("PUT").Name("AtualizarAluno")
	subrouterAluno.Path("/").Queries("personID", "{personID}").HandlerFunc(service.RemoverAluno).Methods("DELETE").Name("RemoverAluno")

	subrouterProfessor := router.PathPrefix("/teacher").Subrouter()
	subrouterProfessor.Path("").HandlerFunc(service.ListarProfessor).Methods("GET")
	subrouterProfessor.Path("/").Queries("personID", "{personID}").HandlerFunc(service.BuscarProfessor).Methods("GET").Name("BuscarProfessor")
	subrouterProfessor.Path("/").Queries("courseID", "{courseID}", "name", "{name}", "cpf", "{cpf}", "cellPhone", "{cellPhone}", "city", "{city}", "zipCode", "{zipCode}", "address", "{address}", "registrationDate", "{registrationDate}").HandlerFunc(service.InserirProfessor).Methods("POST").Name("InserirProfessor")
	subrouterProfessor.Path("/").Queries("personID", "{personID}", "courseID", "{courseID}", "name", "{name}", "cpf", "{cpf}", "cellPhone", "{cellPhone}", "city", "{city}", "zipCode", "{zipCode}", "address", "{address}", "registrationDate", "{registrationDate}").HandlerFunc(service.AtualizarProfessor).Methods("PUT").Name("AtualizarProfessor")
	subrouterProfessor.Path("/").Queries("personID", "{personID}").HandlerFunc(service.RemoverProfessor).Methods("DELETE").Name("RemoverProfessor")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8000", router))

	return router
}

func exampleDecoderJSON() {
	jsonStream := []byte(`{"ID":3,"Name":"Name 3","Description":"Description 3","RegistrationDate":"2019-01-21T10:07:16.543807Z"}`)

	var course *entity.Course
	err := json.Unmarshal(jsonStream, &course)
	if err != nil {
		fmt.Println("Error:", err)
		log.Panic(err)
	}
	log.Println(course.ToString())
}

/// >> test performance string, buffer, builder ///
func testeConcat(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str += "x"
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func testeBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
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
		b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
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
		b.Errorf("unexpected result; got=%s, want=%s", strBuilder.String(), s)
	}
}

/// << test performance string, buffer, builder ///

/// >> example create table ///
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

	if _, err := stmt.Exec("course"); err != nil {
		log.Fatal(err)
	}
}

/// << example create table ///
