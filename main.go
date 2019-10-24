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

const domain = "localhost:8080"

func main() {
	addBundle()
	router := routersConfigure()
	serverConfigure(router)
	//exampleDecoderJSON()
}

func addBundle() {
	//https://phrase.com/blog/posts/internationalisation-in-go-with-go-i18n/
	//https://github.com/nicksnyder/go-i18n
	/*bundle := &i18n.Bundle{DefaultLanguage: language.English}

	loc := i18n.NewLocalizer(bundle, language.English.String())

	messages := &i18n.Message{
		ID:          "Emails",
		Description: "The number of unread emails a user has",
		One:         "{{.Name}} has {{.Count}} email.",
		Other:       "{{.Name}} has {{.Count}} emails.",
	}*/
}

func serverConfigure(router *mux.Router) {
	log.Println("Server Starting!")
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "he duration for which the server normally expects existing connections to end - e.g. 15s or 1m")
	flag.Parse()

	log.Printf("Add routers on: http://%s/", domain)

	srv := &http.Server{
		Addr:         domain,
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
	subrouterCourse.Path("/{id}").HandlerFunc(service.FindByIDCourse).Methods("GET").Name("FindByIDCourse")
	subrouterCourse.Path("").HandlerFunc(service.InsertCourse).Methods("POST").Name("InsertCourse")
	subrouterCourse.Path("").HandlerFunc(service.UpdateCourse).Methods("PUT").Name("UpdateCourse")
	subrouterCourse.Path("/{id}").HandlerFunc(service.DeleteCourse).Methods("DELETE").Name("DeleteCourse")

	subrouterClass := router.PathPrefix("/class").Subrouter()
	subrouterClass.Path("").HandlerFunc(service.FindAllClass).Methods("GET").Name("FindAllClass")
	subrouterClass.Path("/{id}").HandlerFunc(service.FindByIDClass).Methods("GET").Name("FindByIDClass")
	subrouterClass.Path("").HandlerFunc(service.InsertClass).Methods("POST").Name("InsertClass")
	subrouterClass.Path("").HandlerFunc(service.UpdateClass).Methods("PUT").Name("UpdateClass")
	subrouterClass.Path("/{id}").HandlerFunc(service.DeleteClass).Methods("DELETE").Name("DeleteClass")

	subrouterPerson := router.PathPrefix("/person").Subrouter()
	subrouterPerson.Path("").HandlerFunc(service.FindAllPerson).Methods("GET").Name("FindAllPerson")

	subrouterStudent := router.PathPrefix("/student").Subrouter()
	subrouterStudent.Path("").HandlerFunc(service.FindAllStudent).Methods("GET").Name("FindAllStudent")
	subrouterStudent.Path("/{id}").HandlerFunc(service.FindByIDStudent).Methods("GET").Name("FindByIDStudent")
	subrouterStudent.Path("").HandlerFunc(service.InsertStudent).Methods("POST").Name("InsertStudent")
	subrouterStudent.Path("").HandlerFunc(service.UpdateStudent).Methods("PUT").Name("UpdateStudent")
	subrouterStudent.Path("/{id}").HandlerFunc(service.DeleteStudent).Methods("DELETE").Name("DeleteStudent")

	subrouterTeacher := router.PathPrefix("/teacher").Subrouter()
	subrouterTeacher.Path("").HandlerFunc(service.FindAllTeacher).Methods("GET").Name("FindAllTeacher")
	subrouterTeacher.Path("/{id}").HandlerFunc(service.FindByIDTeacher).Methods("GET").Name("FindByIDTeacher")
	subrouterTeacher.Path("").HandlerFunc(service.InsertTeacher).Methods("POST").Name("InsertTeacher")
	subrouterTeacher.Path("").HandlerFunc(service.UpdateTeacher).Methods("PUT").Name("UpdateTeacher")
	subrouterTeacher.Path("/{id}").HandlerFunc(service.DeleteTeacher).Methods("DELETE").Name("DeleteTeacher")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(domain, router))

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
