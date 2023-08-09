package coreapi

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Table struct {
	Id    string `json:"id" binding:"required"`
	Label string `json:"label" binding:"required"`
}
type TableList struct {
	Data *[]Table `json:"data,omitempty"`
}
type Bill struct {
	CheckoutUrl string      `json:"checkoutUrl" binding:"required"`
	Items       []OrderItem `json:"items" binding:"required"`
}
type PutItemRequest struct {
	Name       string      `json:"name" binding:"required"`
	Pricing    int64       `json:"pricing" binding:"required"`
	Attributes []Attribute `json:"attributes" binding:"required"`
	Images     []string    `json:"images" binding:"required"`
}
type Account struct {
	Id    string `json:"id" binding:"required"`
	Email string `json:"email" binding:"required"`
	Role  Role   `json:"role" binding:"required"`
}
type PutRestaurantRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description,omitempty"`
}
type PutTableRequest struct {
	Label string `json:"label" binding:"required"`
}
type Uploading struct {
	Url string `json:"url" binding:"required"`
}
type ItemList struct {
	Data       []Item     `json:"data" binding:"required"`
	Pagination Pagination `json:"pagination" binding:"required"`
}
type Item struct {
	Id         string      `json:"id" binding:"required"`
	Name       string      `json:"name" binding:"required"`
	Pricing    int64       `json:"pricing" binding:"required"`
	Attributes []Attribute `json:"attributes" binding:"required"`
	Images     *[]string   `json:"images,omitempty"`
}
type Option struct {
	Label string `json:"label" binding:"required"`
	Extra int64  `json:"extra" binding:"required"`
}
type CreateSessionRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UpdatePasswordRequest struct {
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}
type RestaurantList struct {
	Data       []Restaurant `json:"data" binding:"required"`
	Pagination Pagination   `json:"pagination" binding:"required"`
}
type Session struct {
	Account     Account `json:"account" binding:"required"`
	Token       string  `json:"token" binding:"required"`
	TokenType   string  `json:"tokenType" binding:"required"`
	TokenFormat string  `json:"tokenFormat" binding:"required"`
	ExpiredAt   int64   `json:"expiredAt" binding:"required"`
	CreatedAt   int64   `json:"createdAt" binding:"required"`
}
type SessionVerificationRequest struct {
	Token string `json:"token" binding:"required"`
}
type CreateAccountRequest struct {
	VerificationCode *string `json:"verificationCode,omitempty"`
	Email            string  `json:"email" binding:"required"`
	Password         string  `json:"password" binding:"required"`
	Role             *Role   `json:"role,omitempty"`
}
type AccountList struct {
	Data       []Account  `json:"data" binding:"required"`
	Pagination Pagination `json:"pagination" binding:"required"`
}
type Pagination struct {
	Index int64 `json:"index" binding:"required"`
	Limit int64 `json:"limit" binding:"required"`
	Total int64 `json:"total" binding:"required"`
}
type Attribute struct {
	Label   string   `json:"label" binding:"required"`
	Options []Option `json:"options" binding:"required"`
}
type OrderItem struct {
	Options *[]Option `json:"options,omitempty"`
	Item    Item      `json:"item" binding:"required"`
}
type SessionVerification struct {
	Status SessionStatus `json:"status" binding:"required"`
}
type Restaurant struct {
	Id          string      `json:"id" binding:"required"`
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Itmes       *[]ItemList `json:"itmes,omitempty"`
}
type Ordering string

const ASCENDING Ordering = "ASCENDING"
const DESCENDING Ordering = "DESCENDING"

type Role string

const ROOT Role = "ROOT"
const ADMIN Role = "ADMIN"
const USER Role = "USER"

type SessionStatus string

const ACTIVED SessionStatus = "ACTIVED"
const EXPIRED SessionStatus = "EXPIRED"
const DISACTIVED SessionStatus = "DISACTIVED"

type AccountApiInterface interface {
	CreateAccount(gin_context *gin.Context, gin_body CreateAccountRequest)
	ListAccount(gin_context *gin.Context, ordering Ordering, index int64, limit int64)
	CreateSession(gin_context *gin.Context, gin_body CreateSessionRequest)
	UpdatePassword(gin_context *gin.Context, gin_body UpdatePasswordRequest)
	VerifySession(gin_context *gin.Context, gin_body SessionVerificationRequest)
	GetAccount(gin_context *gin.Context)
}

func CreateAccountBuilder(api AccountApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		var createAccountRequest CreateAccountRequest
		if err := gin_context.ShouldBindJSON(&createAccountRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.CreateAccount(gin_context, createAccountRequest)
	}
}
func ListAccountBuilder(api AccountApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		ordering := gin_context.Query("ordering")
		index := gin_context.Query("index")
		limit := gin_context.Query("limit")
		api.ListAccount(gin_context, Ordering(ordering), stringToInt64(index), stringToInt64(limit))
	}
}
func CreateSessionBuilder(api AccountApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		var createSessionRequest CreateSessionRequest
		if err := gin_context.ShouldBindJSON(&createSessionRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.CreateSession(gin_context, createSessionRequest)
	}
}
func UpdatePasswordBuilder(api AccountApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		var updatePasswordRequest UpdatePasswordRequest
		if err := gin_context.ShouldBindJSON(&updatePasswordRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.UpdatePassword(gin_context, updatePasswordRequest)
	}
}
func VerifySessionBuilder(api AccountApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		var sessionVerificationRequest SessionVerificationRequest
		if err := gin_context.ShouldBindJSON(&sessionVerificationRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.VerifySession(gin_context, sessionVerificationRequest)
	}
}
func GetAccountBuilder(api AccountApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		api.GetAccount(gin_context)
	}
}
func AccountApiInterfaceMounter(gin_router *gin.Engine, gwg_api_label AccountApiInterface) {
	gin_router.POST("/accounts", CreateAccountBuilder(gwg_api_label))
	gin_router.GET("/accounts", ListAccountBuilder(gwg_api_label))
	gin_router.POST("/account/session", CreateSessionBuilder(gwg_api_label))
	gin_router.PUT("/account/password", UpdatePasswordBuilder(gwg_api_label))
	gin_router.GET("/accounts/session/verification", VerifySessionBuilder(gwg_api_label))
	gin_router.GET("/account", GetAccountBuilder(gwg_api_label))
}

type RestaurantApiInterface interface {
	ListRestaurantItems(gin_context *gin.Context, id string)
	CreateItem(gin_context *gin.Context, id string, gin_body PutItemRequest)
	UpdateTable(gin_context *gin.Context, id string, gin_body PutTableRequest)
	DeleteTable(gin_context *gin.Context, id string)
	GetItem(gin_context *gin.Context, id string)
	UpdateItem(gin_context *gin.Context, id string, gin_body PutItemRequest)
	DeleteItem(gin_context *gin.Context, id string)
	CreateRestaurant(gin_context *gin.Context, gin_body PutRestaurantRequest)
	ListRestaurants(gin_context *gin.Context)
	UpdateRestaurant(gin_context *gin.Context, id string, gin_body PutRestaurantRequest)
	GetRestaurant(gin_context *gin.Context, id string)
	DeleteRestaurant(gin_context *gin.Context, id string)
	ListRestaurantTable(gin_context *gin.Context, id string)
	CreateTable(gin_context *gin.Context, id string, gin_body PutTableRequest)
	UploadItemImage(gin_context *gin.Context, id string)
}

func ListRestaurantItemsBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.ListRestaurantItems(gin_context, id)
	}
}
func CreateItemBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		var putItemRequest PutItemRequest
		if err := gin_context.ShouldBindJSON(&putItemRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.CreateItem(gin_context, id, putItemRequest)
	}
}
func UpdateTableBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		var putTableRequest PutTableRequest
		if err := gin_context.ShouldBindJSON(&putTableRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.UpdateTable(gin_context, id, putTableRequest)
	}
}
func DeleteTableBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.DeleteTable(gin_context, id)
	}
}
func GetItemBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.GetItem(gin_context, id)
	}
}
func UpdateItemBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		var putItemRequest PutItemRequest
		if err := gin_context.ShouldBindJSON(&putItemRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.UpdateItem(gin_context, id, putItemRequest)
	}
}
func DeleteItemBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.DeleteItem(gin_context, id)
	}
}
func CreateRestaurantBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		var putRestaurantRequest PutRestaurantRequest
		if err := gin_context.ShouldBindJSON(&putRestaurantRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.CreateRestaurant(gin_context, putRestaurantRequest)
	}
}
func ListRestaurantsBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		api.ListRestaurants(gin_context)
	}
}
func UpdateRestaurantBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		var putRestaurantRequest PutRestaurantRequest
		if err := gin_context.ShouldBindJSON(&putRestaurantRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.UpdateRestaurant(gin_context, id, putRestaurantRequest)
	}
}
func GetRestaurantBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.GetRestaurant(gin_context, id)
	}
}
func DeleteRestaurantBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.DeleteRestaurant(gin_context, id)
	}
}
func ListRestaurantTableBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.ListRestaurantTable(gin_context, id)
	}
}
func CreateTableBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		var putTableRequest PutTableRequest
		if err := gin_context.ShouldBindJSON(&putTableRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.CreateTable(gin_context, id, putTableRequest)
	}
}
func UploadItemImageBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.UploadItemImage(gin_context, id)
	}
}
func RestaurantApiInterfaceMounter(gin_router *gin.Engine, gwg_api_label RestaurantApiInterface) {
	gin_router.GET("/restaurants/:id/items", ListRestaurantItemsBuilder(gwg_api_label))
	gin_router.POST("/restaurants/:id/items", CreateItemBuilder(gwg_api_label))
	gin_router.PUT("/tables/:id", UpdateTableBuilder(gwg_api_label))
	gin_router.DELETE("/tables/:id", DeleteTableBuilder(gwg_api_label))
	gin_router.GET("/items/:id", GetItemBuilder(gwg_api_label))
	gin_router.PUT("/items/:id", UpdateItemBuilder(gwg_api_label))
	gin_router.DELETE("/items/:id", DeleteItemBuilder(gwg_api_label))
	gin_router.POST("/restaurants", CreateRestaurantBuilder(gwg_api_label))
	gin_router.GET("/restaurants", ListRestaurantsBuilder(gwg_api_label))
	gin_router.PUT("/restaurants/:id", UpdateRestaurantBuilder(gwg_api_label))
	gin_router.GET("/restaurants/:id", GetRestaurantBuilder(gwg_api_label))
	gin_router.DELETE("/restaurants/:id", DeleteRestaurantBuilder(gwg_api_label))
	gin_router.GET("/restaurants/:id/tables", ListRestaurantTableBuilder(gwg_api_label))
	gin_router.POST("/restaurants/:id/tables", CreateTableBuilder(gwg_api_label))
	gin_router.POST("/items/:id/image", UploadItemImageBuilder(gwg_api_label))
}
func stringToInt32(s string) int32 {
	if value, err := strconv.ParseInt(s, 10, 32); err == nil {
		return int32(value)
	}
	return 0
}
func stringToInt64(s string) int64 {
	if value, err := strconv.ParseInt(s, 10, 64); err == nil {
		return value
	}
	return 0
}
func stringToFloat32(s string) float32 {
	if value, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(value)
	}
	return 0
}
func stringToFloat64(s string) float64 {
	if value, err := strconv.ParseFloat(s, 64); err == nil {
		return value
	}
	return 0
}
