package app

import (
	"testing"
)

func TestConfigProperties(t *testing.T) {
	if KeySecret() == "" {
		t.Error("propiedad jwt.secret no tiene valor")
	}

	if AppName() == "" {
		t.Error("propiedad app.name no tiene valor")
	}

	if AppVersion() == "" {
		t.Error("propiedad app.version no tiene valor")
	}

	if RoleAdmin() == "" {
		t.Error("propiedad app.roleAdmin no tiene valor")
	}

	if UrlAwsS3() == "" {
		t.Error("propiedad aws.urlS3 no tiene valor")
	}

	if CORS() == "" {
		t.Error("propiedad app.cors no tiene valor")
	}

	if Sentry() == "" {
		t.Error("propiedad sentry.url no tiene valor")
	}
}
