package service

import (
  "encoding/json"
  "log"
  "net/http"

  "github.com/_dev/exemplo-api-rest/model"
)

// FindAllPerson - Returns total list of registered persons.
func FindAllPerson(w http.ResponseWriter, r *http.Request) {
  db, err := model.NewDB()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    log.Panic(err)
    return
  }

  list, err := db.FindAllPerson()
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
