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
  "github.com/gorilla/mux"
)

// FindAllStudent - Returns total list of registered students.
func FindAllStudent(w http.ResponseWriter, r *http.Request) {
  db, err := model.NewDB()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    log.Panic(err)
    return
  }

  list, err := db.FindAllStudent()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    log.Panic(err)
    db.CloseDB()
    return
  }

  for _, item := range list {
    log.Println(item.ToString())
  }

  db.CloseDB()
  json.NewEncoder(w).Encode(list)
}

// FindByIDStudent - Returns a specific student by ID.
func FindByIDStudent(w http.ResponseWriter, r *http.Request) {
  db, err := model.NewDB()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    log.Panic(err)
    return
  }

  params := mux.Vars(r)
  id, err := strconv.ParseInt(params["id"], 10, 64)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    log.Panic(err)
    db.CloseDB()
    return
  }

  entityy, err := db.FindByIDStudent(id)
  switch {
  case err == sql.ErrNoRows:
    var errorDesc bytes.Buffer
    errorDesc.WriteString("ERROR: No records found for id=")
    errorDesc.WriteString(strconv.FormatInt(id, 10))
    log.Println(errorDesc.String())
    json.NewEncoder(w).Encode(errorDesc.String())
    db.CloseDB()
    return
  case err != nil:
    log.Panic(err)
    db.CloseDB()
    return
  default:
  }

  log.Println(entityy.ToString())
  db.CloseDB()
  json.NewEncoder(w).Encode(entityy)
}

// InsertStudent - Inserts a new student record in the data base.
func InsertStudent(w http.ResponseWriter, r *http.Request) {
  db, err := model.NewDB()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    log.Panic(err)
    return
  }

  var entityy entity.Student
  _ = json.NewDecoder(r.Body).Decode(&entityy)

  id, err := db.InsertPerson(entityy.Person)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    log.Panic(err)
    db.CloseDB()
    return
  }
  entityy.Person.ID = id

  entityyClass, err := db.FindByIDCourse(entityy.Class.ID)
  switch {
  case err == sql.ErrNoRows:
    idClass, err := db.InsertClass(entityy.Class)
    if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      log.Panic(err)
      db.CloseDB()
      return
    }
    entityy.Class.ID = idClass
  case err != nil:
    log.Panic(err)
    db.CloseDB()
    return
  default:
    entityy.Class.ID = entityyClass.ID
  }

  idReturned, err := db.InsertStudent(entityy)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    log.Panic(err)
    db.CloseDB()
    return
  }
  db.CloseDB()
  json.NewEncoder(w).Encode(idReturned)
}

// UpdateStudent - Updates a base student record.
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
  db, err := model.NewDB()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    log.Panic(err)
    return
  }

  var entityy entity.Student
  _ = json.NewDecoder(r.Body).Decode(&entityy)

  if err = db.UpdateStudent(entityy); err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    log.Panic(err)
    db.CloseDB()
    return
  }
  db.CloseDB()
}

// DeleteStudent - Removes a record from the base.
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
  db, err := model.NewDB()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    log.Panic(err)
    return
  }

  delPO := deletePO{"student", "id_person", "id"}
  if err := delPO.Delete(w, r, db); err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    log.Panic(err)
    db.CloseDB()
    return
  }

  delPO = deletePO{"person", "id", "id"}
  if err := delPO.Delete(w, r, db); err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    log.Panic(err)
    db.CloseDB()
    return
  }
  db.CloseDB()
}
