package service

import (
  "net/http"
  "strconv"

  "github.com/_dev/exemplo-api-rest/model"

  "github.com/gorilla/mux"
)

// DeletePO - Parameters usage in delete.
type deletePO struct {
  entity     string
  column     string
  parametter string
}

// Delete - Removes a record from the base.
func (po deletePO) Delete(w http.ResponseWriter, r *http.Request, db *model.DB) error {
  params := mux.Vars(r)
  id, err := strconv.ParseInt(params[po.parametter], 10, 64)
  if err != nil {
    return err
  }

  if err = db.Delete(po.entity, po.column, id); err != nil {
    return err
  }
  return nil
}
