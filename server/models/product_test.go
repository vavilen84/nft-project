package models

//func TestProduct_ValidateOnCreate(t *testing.T) {
//
//	beforeTestRun()
//	db := store.GetTestDB()
//	ctx := store.GetDefaultDBContext()
//	conn, connErr := db.Conn(ctx)
//	if connErr != nil {
//		helpers.LogFatal(connErr)
//	}
//	defer conn.Close()
//	prepareTestDB(ctx, conn)
//
//	m := Product{}
//	err := validation.ValidateByScenario(constants.ScenarioCreate, m)
//	assert.NotNil(t, err)
//	assert.NotEmpty(t, err.(validation.Errors)[constants.ProductPriceField])
//	assert.NotEmpty(t, err.(validation.Errors)[constants.ProductSKUField])
//	assert.NotEmpty(t, err.(validation.Errors)[constants.ProductTitleField])
//
//	m = Product{
//		Title: "product",
//		SKU:   "sku-123_123",
//		Price: 100,
//	}
//	err = validation.ValidateByScenario(constants.ScenarioCreate, m)
//	assert.Empty(t, err)
//}
//
//func TestProduct_Create(t *testing.T) {
//
//	beforeTestRun()
//	db := store.GetTestDB()
//	ctx := store.GetDefaultDBContext()
//	conn, connErr := db.Conn(ctx)
//	if connErr != nil {
//		helpers.LogFatal(connErr)
//	}
//	defer conn.Close()
//	prepareTestDB(ctx, conn)
//
//	model := Product{
//		Title: "title",
//		SKU:   "sku",
//		Price: 1,
//	}
//	err := model.Create(ctx, conn)
//	assert.Nil(t, err)
//
//	modelFromDb, err := FindProductById(ctx, conn, model.GetId())
//	assert.Nil(t, err)
//	assert.Equal(t, model.Id, modelFromDb.Id)
//	assert.Equal(t, model.Title, modelFromDb.Title)
//	assert.Equal(t, model.SKU, modelFromDb.SKU)
//	assert.Equal(t, model.Price, modelFromDb.Price)
//}
