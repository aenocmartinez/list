package app

import (
	"os"
)

func KeySecret() string {
	return os.Getenv("JWT_SECRET_KEY")
}

func AppName() string {
	return os.Getenv("APP_NAME")
}

func AppVersion() string {
	return os.Getenv("APP_VERSION")
}

func AppPath() string {
	return AppName() + "/" + AppVersion()
}

func RoleAdmin() string {
	return os.Getenv("ROLE_ADMIN")
}

func UrlAwsS3() string {
	return os.Getenv("ENDPOINT_UPLOAD")
}

func CORS() string {
	return os.Getenv("CORS")
}

func Sentry() string {
	return os.Getenv("DNS_SENTRY")
}
