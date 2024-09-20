package global

import (
	"bluebell/pkg/snowflake"
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var (
	Log   *logrus.Logger
	DB    *sql.DB
	Snflk *snowflake.Snowflake
	RDB   *redis.Client
)
