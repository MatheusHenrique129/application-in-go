package consts

const (
	// Tables

	TableUser = "users"

	// User Fields

	FieldUserID = "id"

	// Error Database Code

	GenericCode                 = "C-0000"
	ConnectionCode              = "C-0001"
	AlreadyExistsCode           = "C-0002"
	InternalIDAlreadyExistsCode = "C-0003"

	// Errors General

	MySQLErrorAlreadyExists = 1062

	// UNIQUE index error postfix

	MySQLUniqueIndexErrorPostfixCpf = "cpf_uk'"
)
