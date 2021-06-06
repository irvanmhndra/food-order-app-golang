package controller

import (
	"strconv"
	"time"

	helper "order-food-app-golang/helper"
	"order-food-app-golang/model"
	"order-food-app-golang/server/request"

	"github.com/gin-gonic/gin"
)

// Item ...
type Item struct{}

// ListItem ...
func (U *Item) ListItem(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = helper.DefaultPage
	}
	limit, offset := helper.PaginationPageOffset(page, limit)
	ItemModel := model.ItemModel{}
	data, count, err := ItemModel.GetAll(offset, limit)
	if err != nil {
		statusCode = 400

		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}
	pagination := helper.PaginationRes(page, count, limit)
	params := map[string]interface{}{
		"payload": data,
		"meta":    pagination,
	}
	c.JSON(statusCode, helper.OutputAPIResponseWithPayload(params))
}

// FindOneByUUID ...
func (U *Item) FindByID(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	param := c.Params.ByName("id")
	id, err := strconv.Atoi(param)

	ItemModel := model.ItemModel{}
	data, err := ItemModel.FindByID(float64(id))
	if err != nil {
		statusCode = 400

		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}
	params := map[string]interface{}{
		"meta":    "success",
		"payload": data,
	}
	c.JSON(statusCode, helper.OutputAPIResponseWithPayload(params))
}

// Creat ...
func (U *Item) Create(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200

	var body request.SetItem

	// Validation req body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		statusCode = 406

		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	ItemModel := model.ItemModel{}
	err = ItemModel.Create(body)
	if err != nil {
		statusCode = 400

		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": "Error creating Item",
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	params := map[string]interface{}{
		"meta":    "success",
		"payload": body,
	}
	c.JSON(statusCode, helper.OutputAPIResponseWithPayload(params))
}

// Update ...
func (G *Item) Update(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	var body request.SetItem

	// Validation req body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		statusCode = 406

		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	now := time.Now().Local()

	ItemModel := model.ItemModel{}
	err = ItemModel.Update(id, body, now)
	if err != nil {
		statusCode = 400
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}
	params := map[string]interface{}{
		"meta": "success",
	}
	c.JSON(statusCode, helper.OutputAPIResponseWithPayload(params))
}

func (G *Item) UpdateStatus(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	status := c.Query("status")

	now := time.Now().Local()

	ItemModel := model.ItemModel{}
	err := ItemModel.UpdateStatus(id, status, now)
	if err != nil {
		statusCode = 400
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}
	params := map[string]interface{}{
		"meta": "success",
	}
	c.JSON(statusCode, helper.OutputAPIResponseWithPayload(params))
}
