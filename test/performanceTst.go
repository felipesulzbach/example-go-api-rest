package test

import (
  "bytes"
  "log"
  "strings"
  "testing"

  "github.com/_dev/exemplo-api-rest/model"
)

/// >> performance test string, buffer, builder ///
func TesteConcat(b *testing.B) {
  var str string
  for n := 0; n < b.N; n++ {
    str += "x"
  }
  b.StopTimer()

  if s := strings.Repeat("x", b.N); str != s {
    b.Errorf("unexpected result; got=%s, want=%s", str, s)
  }
}

func TesteBuffer(b *testing.B) {
  var buffer bytes.Buffer
  for n := 0; n < b.N; n++ {
    buffer.WriteString("x")
  }
  b.StopTimer()

  if s := strings.Repeat("x", b.N); buffer.String() != s {
    b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
  }
}

func TesteCopy(b *testing.B) {
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

func TesteStringBuilder(b *testing.B) {
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
func TesteCreateTable() {
  db, err := model.NewDB()
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
