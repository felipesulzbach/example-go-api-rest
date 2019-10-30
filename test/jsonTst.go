package test

import (
  "encoding/json"
  "fmt"
  "log"

  "github.com/_dev/exemplo-api-rest/model/entity"
)

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
