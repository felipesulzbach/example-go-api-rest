package migrate

import (
	"time"

)

type file struct {
	fileName string
	filePath string
}

type migrate struct {
	MigrateName      string
	RegistrationDate time.Time
}
