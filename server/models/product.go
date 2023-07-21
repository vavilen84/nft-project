package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/interfaces"
	"github.com/vavilen84/nft-project/orm"
	"github.com/vavilen84/nft-project/validation"
	"log"
	"regexp"
	"time"
)

type Product struct {
	Id    uint32 `json:"id" column:"id"`
	Title string `json:"title" column:"title"`
	SKU   string `json:"sku" column:"sku"`
	Price uint64 `json:"price" column:"price"`

	CreatedAt int64 `json:"created_at" column:"created_at"`
	UpdatedAt int64 `json:"updated_at" column:"updated_at"`
	DeletedAt int64 `json:"deleted_at" column:"deleted_at"`
}

func (m Product) GetId() uint32 {
	return m.Id
}

func (m Product) SetId(id uint32) interfaces.Model {
	m.Id = id
	return m
}

func (m Product) SetCreatedAt() interfaces.Model {
	m.CreatedAt = time.Now().Unix()
	return m
}

func (m Product) SetUpdatedAt() interfaces.Model {
	m.UpdatedAt = time.Now().Unix()
	return m
}

func (m Product) SetDeletedAt() interfaces.Model {
	m.DeletedAt = time.Now().Unix()
	return m
}

func (Product) GetTableName() string {
	return constants.ProductDBTable
}

func (Product) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioCreate: validation.FieldRules{
			constants.ProductTitleField: "required,min=1,max=255",
			constants.ProductSKUField:   "required,min=1,max=255,sku",
			constants.ProductPriceField: "required,min=0,max=999999999999",
		},
	}
}

func (Product) GetValidator() interface{} {
	v := validator.New()
	err := v.RegisterValidation("sku", ValidateSKU)
	if err != nil {
		helpers.LogError(err)
	}
	return v
}

func ValidateSKU(fl validator.FieldLevel) (r bool) {
	pattern := `^[a-z0-9_-]*$`
	r, err := regexp.Match(pattern, []byte(fl.Field().String()))
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func (m *Product) Create(ctx context.Context, conn *sql.Conn) (err error) {
	res, err := orm.Create(ctx, conn, *m)
	*m = res.(Product)
	if err != nil {
		log.Println(err)
	}
	return
}

func FindProductById(ctx context.Context, conn *sql.Conn, id uint32) (m Product, err error) {
	m.Id = id
	row := conn.QueryRowContext(ctx, `SELECT * FROM `+m.GetTableName()+` WHERE id = ?`, id)
	err = row.Scan(&m.Id, &m.Title, &m.SKU, &m.Price, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		log.Println(err)
	}
	return
}
