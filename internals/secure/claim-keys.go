package secure

// JWT constant
const JWT_KEY_AUD string = "aud"
const JWT_KEY_ISS string = "iss"
const JWT_KEY_UUID string = "uuid"
const JWT_KEY_USER_NAME string = "sub"
const JWT_KEY_ROLES string = "roles"
const JWT_KEY_COMPANY_UUID string = "companyUuid"
const JWT_KEY_FACILITY_UUID string = "facilityUuid"
const JWT_KEY_TYPE string = "type"
const JWT_KEY_STUDY_UID string = "studyUid"
const JWT_KEY_PACS_PATIENT_ID string = "pacsPatientId"
const JWT_TOKEN_TYPE string = "bearer"

const JWT_TYPE_PATIENT string = "PATIENT"
const JWT_TYPE_STAFF string = "STAFF"
const JWT_TYPE_PACS_VIEWER string = "PACS-VIEWER"
const JWT_TYPE_INTEGRATION string = "ITG"
const JWT_TYPE_TELERAD_PARTNER string = "TELERAD-PARTNER"

const JWT_AUD_TELERAD string = "TELERAD"

const JWT_ISS_TELERAD string = "TELERAD"

// Header constant
const HEADER_FACILITY_UUID string = "Facility-Uuid"

const HEADER_IS_MOBILE_PLATFORM string = "Sec-Ch-Ua-Mobile"
const HEADER_MOBILE_PLATFORM string = "Sec-Ch-Ua-Platform"
const HEADER_ROLE string = "role"
