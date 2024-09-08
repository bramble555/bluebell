package global

import (
	"database/sql"
	"webapp/pkg/snowflake"

	"github.com/sirupsen/logrus"
)

var (
	Log  *logrus.Logger
	DB   *sql.DB
	Snflk *snowflake.Snowflake
)
