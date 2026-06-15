package constants

import (
	"time"
)

var GMT7_TIMEZONE = time.FixedZone("ICT", 7*60*60)

const AUTHORIZATION string = "Authorization"

const TIME_OUT_EXECUTE time.Duration = 10    //sec
const TIME_OUT_CONNECTION time.Duration = 10 //sec

// JWT constant
const JWT_KEY_USER_ID string = "userId"
const JWT_KEY_USER_NAME string = "sub"
const JWT_KEY_PHONE string = "phone"
const JWT_KEY_EMAIL_VALIDATED string = "emailValidated"
const JWT_KEY_FIRST_LOGIN string = "firstLogin"
const JWT_TOKEN_TYPE string = "bearer"
const JWT_KEY_ROLE string = "roles"
const JWT_KEY_JWT_MD5 string = "jwtMd5" // thêm vào để lưu cache thôi

const HOSPITAL_ID_JWT_KEY string = "hisId"
const HOSPITAL_CODE_JWT_KEY string = "hospitalCode"
const USER_ID_JWT_KEY string = "userID"
const ROLE_JWT_KEY string = "roles"

// Staff roles
const ROLE_ADMIN string = "ADMIN"
const ROLE_DOCTOR string = "DOCTOR"

// Warning Type
const WARNING_TYPE_TELEGRAM = 1

const VIETNAM_COUNTRY_ID = 1
