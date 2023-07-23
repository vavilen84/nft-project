package constants

import (
	"errors"
	"time"
)

var (
	ServerError       = errors.New("Server Error")
	BadRequestError   = errors.New("Bad Request")
	UnauthorizedError = errors.New("Unauthorized")
)

const (
	RoleUser  = 1
	RoleAdmin = 2

	FreeBillingPlan     = 1
	StandardBillingPlan = 2
	ProBillingPlan      = 3

	// email
	// TODO: replace with real domain
	ResetPasswordSubject        = "Benny.com: reset forgotten password"
	ResetPasswordHtmlBodyFormat = "In order to reset your password, please forward this link %s"
	EmailCharSet                = "UTF-8"

	// TODO: replace with real domen
	NoReplySenderEmail = "benny.com"

	//common
	SqlDsnFormat = `%s:%s@tcp(%s:%s)/%s`

	DevelopmentAppEnv = "development"
	ProductionAppEnv  = "production"
	TestingAppEnv     = "testing"

	// validation tags
	RequiredTag = "required"
	MinTag      = "min"
	MaxTag      = "max"
	EmailTag    = "email"

	// validation error messages
	RequiredErrorMsg = "%s resource: '%s' is required"
	MinValueErrorMsg = "%s resource: '%s' min value is %s"
	MaxValueErrorMsg = "%s resource: '%s' max value is %s"
	EmailErrorMsg    = "%s resource: email is not valid"

	// Server
	DefaultWriteTimout  = 60 * time.Second
	DefaultReadTimeout  = 60 * time.Second
	DefaultStoreTimeout = 60 * time.Second

	// scenarios
	ScenarioCreate = "create"
	ScenarioUpdate = "update"
	ScenarioDelete = "delete"
	ScenarioSignUp = "sign-up"
	ScenarioSignIn = "sign-in"
)
