package constant

// CONFIGURATION
const (
	SERVER_KONVEN      = "SERVER.KONVEN"
	SERVER_SYARIAH     = "SERVER.SYARIAH"
	SERVER_GALERI      = "SERVER.G24"
	SERVER_KONVEN_DRC  = "SERVER.KONVEN.DRC"
	SERVER_SYARIAH_DRC = "SERVER.SYARIAH.DRC"
	SERVER_GALERI_DRC  = "SERVER.G24.DRC"

	PREFIX_BRANCH_KONVEN = "1"
	PREFIX_BRANCH_GALERI = "5"
	PREFIX_BRANCH_SYAR   = "6"

	PREFIX_BRANCH_PUSAT_KONVEN = "00002"
	PREFIX_BRANCH_PUSAT_SYAR   = "00003"
	PREFIX_BRANCH_PUSAT_GALERI = "00004"

	PREFIX_MIGRASI_KONVEN = "90"
	PREFIX_MIGRASI_SYAR   = "91"

	FLAG_KONVEN = "K"
	FLAG_SYAR   = "S"
	FLAG_GALERI = "G"

	GADAI    = "gadai"
	MIKRO    = "mikro"
	TABUNGAN = "tabunganemas"

	REQ_INQUIRY = "INQUIRY"
)

// RESPONSE CODE AND DESCRIPTION
const (
	CODE_HTTP_NOT_FOUND      = "X4"
	CODE_HTTP_INTERNAL_ERROR = "X5"
	CODE_APPROVED            = "00"
	CODE_INVALID_TRANSACTION = "12"
	CODE_INVALID_AMOUNT      = "13"
	CODE_ACCOUNT_NOT_FOUND   = "14"
	CODE_INVALID_ACCOUNT     = "15"
	CODE_WRONG_FORMAT        = "30"
	CODE_INVALID_CLIENT      = "31"
	CODE_TIMEOUT             = "68"
	CODE_INVALID_BILL        = "88"
	CODE_CUT_OFF_TIME        = "90"
	CODE_SYSTEM_MAINTENANCE  = "96"
	CODE_GENERAL_ERROR       = "98"

	DESC_HTTP_NOT_FOUND      = "There is No Resource Path"
	DESC_HTTP_INTERNAL_ERROR = "Service Internal Error"
	DESC_APPROVED            = "Approved"
	DESC_INVALID_TRANSACTION = "Invalid Transaction"
	DESC_INVALID_AMOUNT      = "Invalid Amount"
	DESC_ACCOUNT_NOT_FOUND   = "Account Not Found"
	DESC_INVALID_ACCOUNT     = "Invalid Account"
	DESC_WRONG_FORMAT        = "Wrong Format"
	DESC_INVALID_CLIENT      = "Invalid Client"
	DESC_TIMEOUT             = "Transaction Timeout"
	DESC_INVALID_BILL        = "Invalid Bill"
	DESC_CUT_OFF_TIME        = "Cut Off Time"
	DESC_SYSTEM_MAINTENANCE  = "System Maintenance"
	DESC_GENERAL_ERROR       = "General Error"
)

// ERROR MESSAGE
const (
	ERR_DATABASE     = "Error Database"
	ERR_ROWS_PARSING = "Error Parsing Rows Results"

	MSG_ERR_REQ_BODY                   = "Error IO Read Request Body"
	MSG_ERR_RES_BODY                   = "Error IO Read Response Body"
	MSG_ERR_JSON_PARSING_REQ           = "Error Parsing JSON Request Body"
	MSG_ERR_JSON_PARSING_RES           = "Error Parsing JSON Response Body"
	MSG_ERR_POST_HTTP                  = "Error Post Http"
	MSG_ERR_CLIENT_ROLE_UNAUTHORIZE    = "Client Role UnAuthorize"
	MSG_ERR_CLIENT_PRODUCT_UNAUTHORIZE = "Client Product UnAuthorize"
	MSG_ERR_PARAM_NOT_FOUND            = "Parameter Not Found"
	MSG_ERR_URL_NOT_FOUND              = "Rest URL Not Found"
	MSG_ERR_WRONG_FORMAT               = "Wrong Format"
	MSG_ERR_DATA_PARSING               = "Error Data Parsing"
)
