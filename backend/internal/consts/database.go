package consts

const (
	// Tables

	TableUser = "users"

	// User Fields

	FieldUserID      = "id"
	FieldName        = "full_name"
	FieldCpf         = "cpf"
	FieldEmail       = "email"
	FieldAddress     = "address"
	FieldPhoneNumber = "phone_number"
	FieldGender      = "gender"
	FieldPassword    = "password"
	FieldBirthDate   = "birth_date"
	FieldCreatedAt   = "created_at"
	FieldUpdatedAt   = "updated_at"

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
