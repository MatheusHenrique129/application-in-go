package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/consts"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-sql-driver/mysql"

	// Goqu need this import to support MySQL queries
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

// goquMysql is a MySQL query generator. Used with goqu lib
var goquMysql = goqu.Dialect(config.DriverName)

type BaseRepository struct {
	config *config.Config
	db     *sql.DB
	logger *util.Logger
}

type RepoError interface {
	Code() string
	Error() string
	Reason() string
}

// Compile time error interface check
var _ error = &repositoryError{}

type repositoryError struct {
	err    error
	code   string
	reason string
}

func (e *repositoryError) Error() string {
	return e.err.Error()
}

func (e *repositoryError) Reason() string {
	return e.reason
}

func (e *repositoryError) Code() string {
	return e.code
}

func NewFromDatabaseError(err error) RepoError {
	// If there is no rows from DB, we should always return nil
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}

	var mysqlErr *mysql.MySQLError

	if errors.As(err, &mysqlErr) {
		return newFromMySQLError(mysqlErr)
	}

	return newGenericError(err)
}

func newFromMySQLError(err *mysql.MySQLError) RepoError {
	if err.Number == consts.MySQLErrorAlreadyExists {
		return newError(err, consts.AlreadyExistsCode)
	}

	switch {
	case strings.Contains(err.Message, consts.MySQLUniqueIndexErrorPostfixCpf):
		return newError(err, consts.InternalIDAlreadyExistsCode)
	}

	return newGenericError(err)
}

func newGenericError(err error) RepoError {
	return newError(err, consts.GenericCode)
}

var errorMap = map[string]string{
	consts.GenericCode:    "Error in Database",
	consts.ConnectionCode: "Error Connecting to Database",
}

func newError(err error, code string) RepoError {
	return &repositoryError{
		err:    err,
		reason: errorMap[code],
		code:   code,
	}
}

func NewBaseRepository(conf *config.Config, db *sql.DB, logger *util.Logger) BaseRepository {
	return BaseRepository{
		config: conf,
		db:     db,
		logger: logger,
	}
}
