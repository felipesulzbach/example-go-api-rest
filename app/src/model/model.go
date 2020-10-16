package model

import (
	"bytes"

)

// SeparatorEntityFront - Used to separate entity information.
const SeparatorEntityFront = ": '"

// SeparatorEntityBehind - Used to separate entity information.
const SeparatorEntityBehind = "', "

// SeparatorEntityBehind2 - Used to separate entity information.
const SeparatorEntityBehind2 = "'"

// ToString - Returns string with entity information.
func ToString(entityName string, fields map[string]string) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(entityName)
	buffer.WriteString(" = {")

	total := len(fields)
	count := 0
	for field, valueField := range fields {
		buffer.WriteString(field)
		buffer.WriteString(SeparatorEntityFront)
		buffer.WriteString(valueField)
		count++
		if total != count {
			buffer.WriteString(SeparatorEntityBehind)
		} else {
			buffer.WriteString(SeparatorEntityBehind2)
		}
	}
	buffer.WriteString("}]")
	return buffer.String()
}
