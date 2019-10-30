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

// FindAllTeacher - Returns total list of registered teachers.
func FindAllTeacher(w http.ResponseWriter, r *http.Request) {
  db, err := model.NewDB()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    log.Panic(err)
    return
  }

  list, err := db.FindAllTeacher()
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

// FindByIDTeacher - Returns a specific teacher by ID.
func FindByIDTeacher(w http.ResponseWriter, r *http.Request) {
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

  entityy, err := db.FindByIDTeacher(id)
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

// InsertTeacher - Inserts a new class record in the data base.
func InsertTeacher(w http.ResponseWriter, r *http.Request) {
  db, err := model.NewDB()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    log.Panic(err)
    return
  }

  var entityy entity.Teacher
  _ = json.NewDecoder(r.Body).Decode(&entityy)

  id, err := db.NextIDPerson()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    log.Panic(err)
    db.CloseDB()
    return
  }
  entityy.Person.ID = id

  entityyCourse, err := db.FindByIDCourse(entityy.Course.ID)
  switch {
  case err == sql.ErrNoRows:
    idCourse, err := db.InsertCourse(entityy.Course)
    if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      log.Panic(err)
      db.CloseDB()
      return
    }
    entityy.Course.ID = idCourse
  case err != nil:
    log.Panic(err)
    db.CloseDB()
    return
  default:
    entityy.Course.ID = entityyCourse.ID
  }

  idReturned, err := db.InsertTeacher(entityy)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    log.Panic(err)
    db.CloseDB()
    return
  }
  db.CloseDB()
  json.NewEncoder(w).Encode(idReturned)
}

// UpdateTeacher - Updates a base teacher record.
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
  db, err := model.NewDB()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    log.Panic(err)
    return
  }

  var entityy entity.Teacher
  _ = json.NewDecoder(r.Body).Decode(&entityy)

  if err = db.UpdateTeacher(entityy); err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    log.Panic(err)
    db.CloseDB()
    return
  }
  db.CloseDB()
}

// DeleteTeacher - Removes a record from the base.
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
  db, err := model.NewDB()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    log.Panic(err)
    return
  }

  delPO := deletePO{"teacher", "id_person", "id"}
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
