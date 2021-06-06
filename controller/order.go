package controller

import (
	"strconv"
	"time"

	helper "order-food-app-golang/helper"
	"order-food-app-golang/model"
	"order-food-app-golang/server/request"

	"order-food-app-golang/middleware"

	"github.com/gin-gonic/gin"
)

// Order ...
type Order struct{}

// ListOrder ...
func (U *Order) ListOrder(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = helper.DefaultPage
	}
	limit, offset := helper.PaginationPageOffset(page, limit)
	OrderModel := model.OrderModel{}
	data, count, err := OrderModel.GetAll(offset, limit)
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
func (U *Order) FindByID(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	param := c.Params.ByName("id")
	id, err := strconv.Atoi(param)

	OrderModel := model.OrderModel{}
	data, err := OrderModel.FindByID(float64(id))
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
func (U *Order) Create(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200

	var body request.SetOrder
	User := middleware.GetUserCustom(c)
	UserID := User["user_id"].(int64)

	body.User = UserID

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

	OrderModel := model.OrderModel{}
	err = OrderModel.Create(body)
	if err != nil {
		statusCode = 400

		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": "Error creating Order",
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
func (G *Order) Update(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	var body request.SetOrder

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

	OrderModel := model.OrderModel{}
	err = OrderModel.Update(id, body, now)
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

func (G *Order) UpdateStatus(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	status := c.Query("status")

	now := time.Now().Local()

	OrderModel := model.OrderModel{}
	err := OrderModel.UpdateStatus(id, status, now)
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
