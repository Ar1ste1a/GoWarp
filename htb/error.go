package htb

import "errors"

var (
	HTB_ERROR_UNKNOWN              = errors.New("Unknown error")
	HTB_ERROR_INVALID_API_KEY      = errors.New("Invalid API Key")
	HTB_ERROR_USER_ID_MISSING      = errors.New("User ID missing")
	HTB_ERROR_USER_ID_INVALID      = errors.New("User ID invalid")
	HTB_ERROR_MACHINE_ID_MISSING   = errors.New("Machine ID missing")
	HTB_ERROR_MACHINE_ID_INVALID   = errors.New("Machine ID invalid")
	HTB_ERROR_CHALLENGE_ID_MISSING = errors.New("Challenge ID missing")
	HTB_ERROR_CHALLENGE_ID_INVALID = errors.New("Challenge ID invalid")
	HTB_ERROR_ENDGAME_ID_MISSING   = errors.New("Endgame ID missing")
	HTB_ERROR_ENDGAME_ID_INVALID   = errors.New("Endgame ID invalid")
	HTB_ERROR_FORTRESS_ID_MISSING  = errors.New("Fortress ID missing")
	HTB_ERROR_FORTRESS_ID_INVALID  = errors.New("Fortress ID invalid")
	HTB_ERROR_PROLAB_ID_MISSING    = errors.New("Prolabs ID missing")
	HTB_ERROR_PROLAB_ID_INVALID    = errors.New("Prolabs ID invalid")
	HTB_PAGE_NUMBER_INVALID        = errors.New("Page number invalid")
	HTB_ERROR_PAGE_NUMBER_MISSING  = errors.New("Page number missing")
	LOCAL_ERROR_FILE_NOT_FOUND     = errors.New("File not found")
	LOCAL_ERROR_FILE_READ_ERROR    = errors.New("File read error")
	VPN_ERROR_INVALID_CONFIG       = errors.New("Invalid VPN config")
	LOCAL_ERROR_API_KEY_UNSET      = errors.New("API Key not set")
)
