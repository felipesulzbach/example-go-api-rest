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
	router := routersConfigure()
	serverConfigure(router)
	//exampleDecoderJSON()
}

func serverConfigure(router *mux.Router) {
	log.Println("Server Start!")
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "he duration for which the server normally expects existing connections to end - e.g. 15s or 1m")
	flag.Parse()

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

func routersConfigure() *mux.Router {
	router := mux.NewRouter()
	subrouterCourse := router.PathPrefix("/course").Subrouter()
	subrouterCourse.Path("").HandlerFunc(service.FindAllCourse).Methods("GET").Name("FindAllCourse")
	subrouterCourse.Path("/").Queries("id", "{id}").HandlerFunc(service.FindByIDCourse).Methods("GET").Name("FindByIDCourse")
	subrouterCourse.Path("/").Queries("name", "{name}", "description", "{description}", "registrationDate", "{registrationDate}").HandlerFunc(service.InsertCourse).Methods("POST").Name("InsertCourse")
	subrouterCourse.Path("/").Queries("id", "{id}", "name", "{name}", "description", "{description}", "registrationDate", "{registrationDate}").HandlerFunc(service.UpdateCourse).Methods("PUT").Name("UpdateCourse")
	subrouterCourse.Path("/").Queries("id", "{id}").HandlerFunc(service.DeleteCourse).Methods("DELETE").Name("DeleteCourse")

	subrouterClass := router.PathPrefix("/class").Subrouter()
	subrouterClass.Path("").HandlerFunc(service.FindAllClass).Methods("GET").Name("FindAllClass")
	subrouterClass.Path("/").Queries("id", "{id}").HandlerFunc(service.FindByIDClass).Methods("GET").Name("FindByIDClass")
	subrouterClass.Path("/").Queries("courseID", "{courseID}", "startDate", "{startDate}", "endDate", "{endDate}", "registrationDate", "{registrationDate}").HandlerFunc(service.InsertClass).Methods("POST").Name("InsertClass")
	subrouterClass.Path("/").Queries("id", "{id}", "courseID", "{courseID}", "startDate", "{startDate}", "endDate", "{endDate}", "registrationDate", "{registrationDate}").HandlerFunc(service.UpdateClass).Methods("PUT").Name("UpdateClass")
	subrouterClass.Path("/").Queries("id", "{id}").HandlerFunc(service.DeleteClass).Methods("DELETE").Name("DeleteClass")

	subrouterPerson := router.PathPrefix("/person").Subrouter()
	subrouterPerson.Path("").HandlerFunc(service.FindAllPerson).Methods("GET").Name("FindAllPerson")

	subrouterStudent := router.PathPrefix("/student").Subrouter()
	subrouterStudent.Path("").HandlerFunc(service.FindAllStudent).Methods("GET").Name("FindAllStudent")
	subrouterStudent.Path("/").Queries("personID", "{personID}").HandlerFunc(service.FindByIDStudent).Methods("GET").Name("FindByIDStudent")
	subrouterStudent.Path("/").Queries("classID", "{classID}", "name", "{name}", "cpf", "{cpf}", "cellPhone", "{cellPhone}", "city", "{city}", "zipCode", "{zipCode}", "address", "{address}", "registrationDate", "{registrationDate}").HandlerFunc(service.InsertStudent).Methods("POST").Name("InsertStudent")
	subrouterStudent.Path("/").Queries("personID", "{personID}", "classID", "{classID}", "name", "{name}", "cpf", "{cpf}", "cellPhone", "{cellPhone}", "city", "{city}", "zipCode", "{zipCode}", "address", "{address}", "registrationDate", "{registrationDate}").HandlerFunc(service.UpdateStudent).Methods("PUT").Name("UpdateStudent")
	subrouterStudent.Path("/").Queries("personID", "{personID}").HandlerFunc(service.DeleteStudent).Methods("DELETE").Name("DeleteStudent")

	subrouterTeacher := router.PathPrefix("/teacher").Subrouter()
	subrouterTeacher.Path("").HandlerFunc(service.FindAllTeacher).Methods("GET").Name("FindAllTeacher")
	subrouterTeacher.Path("/").Queries("personID", "{personID}").HandlerFunc(service.FindByIDTeacher).Methods("GET").Name("FindByIDTeacher")
	subrouterTeacher.Path("/").Queries("courseID", "{courseID}", "name", "{name}", "cpf", "{cpf}", "cellPhone", "{cellPhone}", "city", "{city}", "zipCode", "{zipCode}", "address", "{address}", "registrationDate", "{registrationDate}").HandlerFunc(service.InsertTeacher).Methods("POST").Name("InsertTeacher")
	subrouterTeacher.Path("/").Queries("personID", "{personID}", "courseID", "{courseID}", "name", "{name}", "cpf", "{cpf}", "cellPhone", "{cellPhone}", "city", "{city}", "zipCode", "{zipCode}", "address", "{address}", "registrationDate", "{registrationDate}").HandlerFunc(service.UpdateTeacher).Methods("PUT").Name("UpdateTeacher")
	subrouterTeacher.Path("/").Queries("personID", "{personID}").HandlerFunc(service.DeleteTeacher).Methods("DELETE").Name("DeleteTeacher")

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

/// >> performance test string, buffer, builder ///
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

/// << performance test string, buffer, builder ///

/// >> example create table ///
func testeCreateTable() {
	db, err := model.NewDB(service.DataSourcePostgre)
	if err != nil {
		log.Panic(err)
		return
	}

	stmt, err := db.Prepare("CREATE TABLE GO_TST.$1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec("course"); err != nil {
		log.Fatal(err)
	}
}

/// << example create table ///
