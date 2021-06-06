package controller

import (
	"strconv"

	helper "order-food-app-golang/helper"
	"order-food-app-golang/middleware"
	"order-food-app-golang/model"
	"order-food-app-golang/server/request"
	responses "order-food-app-golang/server/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct{}

// ListUser ...
func (U *User) ListUser(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = helper.DefaultPage
	}
	limit, offset := helper.PaginationPageOffset(page, limit)
	userModel := model.UserModel{}
	data, count, err := userModel.GetAll(offset, limit)
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
func (U *User) FindByID(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	param := c.Params.ByName("id")
	id, _ := strconv.Atoi(param)

	userModel := model.UserModel{}
	data, err := userModel.FindByID(float64(id))
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

// Login ...
func (U *User) Login(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200

	var body request.Login

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

	userModel := model.UserModel{}
	data, err := userModel.FindByEmail(body.Email)

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

	// Compare Password
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(body.Password))
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

	// Generate Token
	token, err := middleware.CreateToken(data)
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

	// data.Password = ""

	params := map[string]interface{}{
		"meta": map[string]interface{}{
			"token": token,
		},
		"payload": data,
	}
	c.JSON(statusCode, helper.OutputAPIResponseWithPayload(params))
}

// Register ...
func (U *User) Register(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200

	var body request.Register

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

	// Set Request
	req := request.Register{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	// Hash Password
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		statusCode = 406

		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": "Error hashing password",
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}
	req.Password = string(bytes)

	userModel := model.UserModel{}
	userData, _ := userModel.FindByEmail(body.Email)

	if userData.Email != "" {
		statusCode = 400

		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": "User already exist",
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	data, err := userModel.Create(req)
	if err != nil {
		statusCode = 400

		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": "Error creating user",
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	// Set Response
	res := responses.UserModel{
		ID:       data,
		Name:     body.Name,
		Email:    body.Email,
		Password: string(bytes),
	}

	params := map[string]interface{}{
		"meta":    "success",
		"payload": res,
	}
	c.JSON(statusCode, helper.OutputAPIResponseWithPayload(params))
}
