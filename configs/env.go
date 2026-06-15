package configs

import (
	"fmt"
	"time"

	"github.com/BeeTechHub/go-common/logger"

	commonConfigs "github.com/BeeTechHub/go-common/configs"
)

func EnvTeleradCorePostgresURI() string {
	return commonConfigs.GetEnvFromOS("TELERAD_CORE_POSTGRES_URI")
}

func GetTeleradCorePostgresConnectTimeout() time.Duration {
	result, err := commonConfigs.GetInt64EnvFromOS("TELERAD_CORE_POSTGRES_CONNECT_TIMEOUT")
	if err != nil || result <= 0 {
		return 10 * time.Second
	}
	return time.Duration(result) * time.Second
}

func GetTeleradCorePostgresMaxConns() int32 {
	result, err := commonConfigs.GetInt64EnvFromOS("TELERAD_CORE_POSTGRES_MAX_CONNS")
	if err != nil || result <= 0 {
		return 10
	}
	return int32(result)
}

func GetTeleradCorePostgresMinConns() int32 {
	result, err := commonConfigs.GetInt64EnvFromOS("TELERAD_CORE_POSTGRES_MIN_CONNS")
	if err != nil || result < 0 {
		return 1
	}
	return int32(result)
}

// -------------- LOG CONFIG -----------------------------------------------
func GetLogLevel() string {
	return commonConfigs.GetEnvFromOS("LOG_LEVEL")
}

func GetWriteLogToFile() bool {
	val, _ := commonConfigs.GetBoolEnvFromOS("LOG_WRITE_TO_FILE")
	return val
}

// -------------- Start JWT ----------------------------

// GetJwtPublicKey trả về nội dung PEM public key (lưu thẳng trong env JWT_PUBLIC_KEY,
// dạng 1 dòng có escape \n — godotenv tự expand khi load .env).
func GetJwtPublicKey() string {
	return commonConfigs.GetEnvFromOS("JWT_PUBLIC_KEY")
}

// GetJwtPrivateKey trả về nội dung PEM private key (env JWT_PRIVATE_KEY).
func GetJwtPrivateKey() string {
	return commonConfigs.GetEnvFromOS("JWT_PRIVATE_KEY")
}

func GetJwtPublicKeyPath() string {
	return commonConfigs.GetEnvFromOS("JWT_PUBLIC_KEY_PATH")
}

func GetJwtPrivateKeyPath() string {
	return commonConfigs.GetEnvFromOS("JWT_PRIVATE_KEY_PATH")
}

func GetJwtExpiryTime() int64 {
	val, err := commonConfigs.GetInt64EnvFromOS("JWT_EXPIRY_TIME")
	if err != nil {
		fmt.Println("Error getting JWT_EXPIRY_TIME from env:", err)
		return 604800
	}
	return val
}

// -------------- End JWT ----------------------------

// Start get Product Common Config
func GetProductPort() string {
	return commonConfigs.GetEnvFromOS("PRODUCT_PORT")
}

// GetPacsViewerUrl trả URL pacs-viewer (mở ảnh DICOM theo StudyInstanceUID).
func GetPacsViewerUrl() string {
	return commonConfigs.GetEnvFromOS("PACS_VIEWER_URL")
}

// End get Product Common Config

// Start get cache config

func GetCacheClusterName() string {
	return commonConfigs.GetEnvFromOS("CACHE_CLUSTER_NAME")
}

func GetCacheUrl() string {
	return commonConfigs.GetEnvFromOS("CACHE_URL")
}

func GetCacheExprireTime() time.Duration {
	result, err := commonConfigs.GetInt64EnvFromOS("CACHE_EXPIRATION_TIME")
	if err != nil {
		logger.Warn("GetCacheExprireTime error, please check 'CACHE_EXPIRATION_TIME'")
		return time.Duration(12 * 60 * 60 * 1000 * 1000 * 1000)
	}
	return time.Duration(result) * time.Second
}

// End get cache config

// Start get sqs config

func GetUserAccessApiSqsQueueName() string {
	return commonConfigs.GetEnvFromOS("SQS_USER_ACCESS_API_QUEUE_NAME")
}

// END get sqs config

// Start get CORS config
func GetCorsAllowedOrigins() string {
	return commonConfigs.GetEnvFromOS("CORS_ALLOWED_ORIGINS")
}

// GetCorsAllowMethods returns allowed HTTP methods for CORS. Default: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS".
func GetCorsAllowMethods() string {
	v := commonConfigs.GetEnvFromOS("CORS_ALLOW_METHODS")
	if v == "" {
		return "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS"
	}
	return v
}

// GetCorsAllowHeaders returns allowed request headers for CORS. Default: "Origin,Content-Type,Accept,Authorization,X-Requested-With".
func GetCorsAllowHeaders() string {
	v := commonConfigs.GetEnvFromOS("CORS_ALLOW_HEADERS")
	if v == "" {
		return "Origin,Content-Type,Accept,Authorization,X-Requested-With,Application,Facility-Uuid"
	}
	return v
}

// GetCorsExposeHeaders returns exposed response headers for CORS. Default: "Content-Length,Content-Type".
func GetCorsExposeHeaders() string {
	v := commonConfigs.GetEnvFromOS("CORS_EXPOSE_HEADERS")
	if v == "" {
		return "Content-Length,Content-Type"
	}
	return v
}

// End get CORS config
