package util

import (
	"log"
	"time"
)

// FormatarDataHora Retorna a data formatada "yyyy/mm/dd hh:mm:ss".
func FormatarDataHora(datahora time.Time) string {
	// Verificar no link 'https://golang.org/src/time/format.go', os formatos possiveis.
	retorno := datahora.Format("2006/01/02 15:04:05")
	return retorno
}

// StringToTime Converte data de String para Time.
func StringToTime(data string) time.Time {
	datatime, err := time.Parse(time.RFC3339, data)
	if err != nil {
		log.Panic(err)
	}
	return datatime
}
