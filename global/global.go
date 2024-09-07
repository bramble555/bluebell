package global

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
	DB  *sql.DB
)
