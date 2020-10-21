package repository

import (
	"bytes"

)

// Delete - Removes a record from the base.
func (db *DB) Delete(entity string, column string, value int64) error {
	var sqlStatement bytes.Buffer
	sqlStatement.WriteString("DELETE FROM fs_auto.")
	sqlStatement.WriteString(entity)
	sqlStatement.WriteString(" WHERE ")
	sqlStatement.WriteString(column)
	sqlStatement.WriteString("=$1")
	_, err := db.Exec(sqlStatement.String(), value)
	if err != nil {
		return err
	}
	return nil
}
