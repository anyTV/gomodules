package response

const (
	InvalidClient       = "invalid_client"
	InvalidRequest      = "invalid_request"
	InternalServerError = "internal_server_error"
	InvalidResponseType = "invalid_response_type"
	InvalidScope        = "invalid_scope"

	InvalidToken = "invalid_token" // should prompt the user to relogin

	NothingTodoHere = "nothing_todo_here"
)
