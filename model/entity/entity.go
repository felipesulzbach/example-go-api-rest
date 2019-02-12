package entity

import (
	"bytes"
)

// SeparatorEntityFront Usado para separar informacoes entidades.
const SeparatorEntityFront = ": '"

// SeparatorEntityBehind Usado para separar informacoes entidades.
const SeparatorEntityBehind = "', "

// SeparatorEntityBehind2 Usado para separar informacoes entidades.
const SeparatorEntityBehind2 = "'"

// ToString Retorna string com as informacoes da entidade.
func ToString(entidade string, campos map[string]string) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(entidade)
	buffer.WriteString(" = {")

	qtd := len(campos)
	count := 0
	for campo, valorCampo := range campos {
		buffer.WriteString(campo)
		buffer.WriteString(SeparatorEntityFront)
		buffer.WriteString(valorCampo)
		count++
		if qtd != count {
			buffer.WriteString(SeparatorEntityBehind)
		} else {
			buffer.WriteString(SeparatorEntityBehind2)
		}
	}
	buffer.WriteString("}]")
	return buffer.String()
}
