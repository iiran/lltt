package core

const (
	TIME_FORMAT = "1970-01-01 01:01:01"
	C_SPACE     = ' '
	C_LINE      = '\n'
	S_SPACE     = " "
	S_LINE      = "\n"
	TRUE        = "1"
	FALSE       = "0"
)

const (
	COOKIE_NAME_SESSION = "iiran_cookie_name_session"
	DEFAULT_LOGIN_COOKIE_AGE_HOUR = 24
	MIN_USERID = 1
)

const (
	PARAM_PAGE     = "page"
	PARAM_USERNAME = "username"
	PARAM_POST_ID  = "post_id"
	PARAM_REPLY_ID = "reply_id"
)

const (
	QUERY_PAGE      = "page"
	QUERY_PAGE_SIZE = "page_size"
	QUERY_USERNAME  = "username"
	QUERY_REPLYID   = "reply_id"
	QUERY_POSTID    = "post_id"
)

const (
	STORE_PAGE                       = "page"
	STORE_PAGE_SIZE                  = "page_size"
	STORE_OPERAND_USERNAME           = "operand_username"
	STORE_OPERATOR_USERID            = "operator_user_id"
	STORE_OPERAND_USERID             = "operand_user_id"
	STORE_REPLYID                    = "reply_id"
	STORE_POSTID                     = "post_id"
	STORE_USER_DATA                  = "store_user_session_data"
	STORE_USER_SESSION_ID            = "store_user_session_id"
	STORE_MARK_COOKIE_EXPIRED        = "store_mark_cookie_expired"
	STORE_MARK_SHOULD_CREATE_SESSION = "store_mark_should_create_session"
)
