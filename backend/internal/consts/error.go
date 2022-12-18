package consts

const (
	//Error Code

	ValidationErrorCode = "validation_error"
	RepoErrorCode       = "repo_error"

	// Messages

	EmailInvalidMessage            = "Email is invalid."
	FieldIsRequiredMessage         = "Field is required."
	RequestFieldsInvalidTypes      = "Request has fields with invalid type"
	InvalidUserIDMessage           = "User ID is invalid. ID: '%s'"
	InvalidRequestNotParseMessage  = "Invalid Request: Could not parse request"
	ValidationErrorMessage         = "There were some validation errors on the fields"
	InvalidUriValueUserIDMessage   = "Invalid URI values. Please, check that the user_id parameter are correct."
	InvalidRequestJsonMessage      = "Invalid request JSON body."
	UserNotFoundMessage            = "User Not Found :("
	OccurredErrorFindUserMessage   = "An error occurred while searching for the user."
	OccurredErrorCreateUserMessage = "An error occurred while to create a user."
	OccurredErrorUpdateUserMessage = "An error occurred while to update a user."
	OccurredErrorDeleteUserMessage = "An error occurred while to delete a user."
	IDCannotStringMessage          = "id cannot be of type string"
	InternalIDAlreadyExistsMessage = "Internal ID already exists"
)
