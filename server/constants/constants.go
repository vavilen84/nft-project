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

	//common
	SqlDsnFormat = `%s:%s@tcp(%s:%s)/%s`

	// DEPRECATED! use default db & user for testing purposes
	// test db & test user
	TestingDB                   = "testing" // reference docker/db/docker-entrypoint-initdb.d/test.db-sql
	MysqlTestUserEnvVar         = "MYSQL_TEST_USER"
	MysqlTestUserPasswordEnvVar = "MYSQL_TEST_USER_PASSWORD"
	MysqlTestDBEnvVar           = "MYSQL_TEST_DATABASE"

	// env vars
	MysqlUserEnvVar     = "MYSQL_USER"
	MysqlDBEnvVar       = "MYSQL_DATABASE"
	MysqlPasswordEnvVar = "MYSQL_PASSWORD"

	AppRootEnvVar     = "APP_ROOT"
	ProjectRootEnvVar = "PROJECT_ROOT"

	SqlDriverEnvVar          = "SQL_DRIVER"
	MysqlPortEnvVar          = "MYSQL_PORT"
	MysqlHostEnvVar          = "MYSQL_HOST"
	DockerMysqlServiceEnvVar = "DOCKER_MYSQL_SERVICE"

	// app env
	AppEnvEnvVar = "APP_ENV"

	DevelopmentAppEnv = "development"
	ProductionAppEnv  = "production"
	TestingAppEnv     = "testing"

	// db tables
	MigrationDBTable            = "migration"
	ProductDBTable              = "product"
	CustomerDBTable             = "customer"
	OrderDBTable                = "order"
	OrderProductDBTable         = "order_product"
	OrderTaxDBTable             = "order_tax"
	OrderProductTaxDBTable      = "order_product_tax"
	OrderProductDiscountDBTable = "order_product_discount"
	OrderDiscountDBTable        = "order_discount"
	TaxDBTable                  = "tax"

	// migrations
	MigrationsFolder = "migrations"

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

	// model names
	MigrationModel = "Migration"

	// field names

	// migration
	MigrationVersionField   = "Version"
	MigrationFilenameField  = "Filename"
	MigrationCreatedAtField = "CreatedAt"
	MigrationUpdatedAtField = "UpdatedAt"

	//product
	ProductTitleField = "Title"
	ProductSKUField   = "SKU"
	ProductPriceField = "Price"

	//customer
	CustomerFirstNameField = "FirstName"
	CustomerLastNameField  = "LastName"
	CustomerEmailField     = "Email"

	//tax
	TaxTitleField      = "Title"
	TaxAmountField     = "Amount"
	TaxPercentageField = "Percentage"
	TaxTypeField       = "Type"

	//tax types
	TaxCartType     = 1
	TaxCategoryType = 2
	TaxProductType  = 3

	//discount
	DiscountTitleField      = "Title"
	DiscountAmountField     = "Amount"
	DiscountPercentageField = "Percentage"
	DiscountTypeField       = "Type"

	//discount types
	DiscountCartType     = 1
	DiscountCategoryType = 2
	DiscountProductType  = 3

	//order
	OrderCustomerIdField = "CustomerId"

	//order_product
	OrderOrderIdField   = "OrderId"
	OrderProductIdField = "ProductId"
	OrderQuantityField  = "Quantity"

	//order_tax
	OrderTaxOrderIdField = "OrderId"
	OrderTaxTaxIdField   = "TaxId"

	//order_discount
	OrderDiscountOrderIdField    = "OrderId"
	OrderDiscountDiscountIdField = "DiscountId"

	//order_product_tax
	OrderProductTaxOrderProductIdField = "OrderProductId"
	OrderProductTaxTaxIdField          = "TaxId"

	//order_product_discount
	OrderProductDiscountOrderProductIdField = "OrderProductId"
	OrderProductDiscountDiscountIdField     = "DiscountId"

	// common
	CommonCreatedAtField = "CreatedAt"
	CommonUpdatedAtField = "UpdatedAt"
	CommonDeletedAtField = "DeletedAt"

	// scenarios
	ScenarioCreate = "create"
	ScenarioUpdate = "update"
	ScenarioDelete = "delete"
	ScenarioSignUp = "sign-up"
	ScenarioSignIn = "sign-in"
)
