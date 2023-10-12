package coreapi

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Restaurant struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Items       []Item   `json:"items"`
	Tables      []Table  `json:"tables"`
	Tags        []string `json:"tags"`
}
type PrinterList struct {
	Data       []Printer  `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Bill struct {
	Orders []Order    `json:"orders"`
	Id     string     `json:"id"`
	Status BillStatus `json:"status"`
}
type BillList struct {
	Data       []Bill     `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Session struct {
	Account     Account `json:"account"`
	Token       string  `json:"token"`
	TokenType   string  `json:"tokenType"`
	TokenFormat string  `json:"tokenFormat"`
	ExpiredAt   int64   `json:"expiredAt"`
	CreatedAt   int64   `json:"createdAt"`
}
type Pagination struct {
	Index int64 `json:"index"`
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
}
type Order struct {
	Item    Item   `json:"item"`
	Options []Pair `json:"options"`
}
type CreateBillRequest struct {
	Orders []Specification `json:"orders"`
}
type Uploading struct {
	Url string `json:"url"`
}
type PutItemRequest struct {
	Name       string      `json:"name"`
	Pricing    int64       `json:"pricing"`
	Attributes []Attribute `json:"attributes"`
	Images     []string    `json:"images"`
	Tags       []string    `json:"tags"`
	Printers   []string    `json:"printers"`
}
type Item struct {
	Tags       []string    `json:"tags"`
	Printers   []string    `json:"printers"`
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Pricing    int64       `json:"pricing"`
	Attributes []Attribute `json:"attributes"`
	Images     []string    `json:"images"`
}
type AccountList struct {
	Data       []Account  `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Specification struct {
	ItemId  string `json:"itemId"`
	Options []Pair `json:"options"`
}
type ItemList struct {
	Data       []Item     `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type SessionVerificationRequest struct {
	Token string `json:"token"`
}
type CreateAccountRequest struct {
	VerificationCode *string `json:"verificationCode,omitempty"`
	Email            string  `json:"email"`
	Password         string  `json:"password"`
	Role             *Role   `json:"role,omitempty"`
}
type PutPrinterRequest struct {
	Sn          string      `json:"sn"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Type        PrinterType `json:"type"`
}
type Printer struct {
	Id          string      `json:"id"`
	Sn          string      `json:"sn"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Type        PrinterType `json:"type"`
}
type UpdatePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
type RestaurantList struct {
	Data       []Restaurant `json:"data"`
	Pagination Pagination   `json:"pagination"`
}
type SessionVerification struct {
	Status SessionStatus `json:"status"`
}
type RestaurantsConfiguration struct {
	Height *int64    `json:"height,omitempty"`
	Tags   *[]string `json:"tags,omitempty"`
	Width  *int64    `json:"width,omitempty"`
}
type Table struct {
	X     int64  `json:"x"`
	Y     int64  `json:"y"`
	Id    string `json:"id"`
	Label string `json:"label"`
}
type TableList struct {
	Data *[]Table `json:"data,omitempty"`
}
type PutTableRequest struct {
	Label string `json:"label"`
	X     int64  `json:"x"`
	Y     int64  `json:"y"`
}
type Attribute struct {
	Label   string   `json:"label"`
	Options []Option `json:"options"`
}
type Option struct {
	Label string `json:"label"`
	Extra int64  `json:"extra"`
}
type Account struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}
type PutRestaurantRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
type Pair struct {
	Right string `json:"right"`
	Left  string `json:"left"`
}
type CreateSessionRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
type Role string

const ROOT Role = "ROOT"
const ADMIN Role = "ADMIN"
const USER Role = "USER"
const SUB_ACCOUNT Role = "SUB_ACCOUNT"

type SessionStatus string

const ACTIVED SessionStatus = "ACTIVED"
const EXPIRED SessionStatus = "EXPIRED"
const DISACTIVED SessionStatus = "DISACTIVED"

type Ordering string

const ASCENDING Ordering = "ASCENDING"
const DESCENDING Ordering = "DESCENDING"

type BillStatus string

const SUBMITTED BillStatus = "SUBMITTED"
const PAIED BillStatus = "PAIED"

type PrinterType string

const BILL PrinterType = "BILL"
const KITCHEN PrinterType = "KITCHEN"

type AccountApiInterface interface {
	GetAccount(gin_context *gin.Context)
	UpdatePassword(gin_context *gin.Context, gin_body UpdatePasswordRequest)
	CreateSession(gin_context *gin.Context, gin_body CreateSessionRequest)
	CreateAccount(gin_context *gin.Context, gin_body CreateAccountRequest)
	ListAccount(gin_context *gin.Context, ordering Ordering, index int64, limit int64)
}

func GetAccountBuilder(api AccountApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		api.GetAccount(gin_context)
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
func AccountApiInterfaceMounter(gin_router *gin.Engine, gwg_api_label AccountApiInterface) {
	gin_router.GET("/account", GetAccountBuilder(gwg_api_label))
	gin_router.PUT("/account/password", UpdatePasswordBuilder(gwg_api_label))
	gin_router.POST("/account/session", CreateSessionBuilder(gwg_api_label))
	gin_router.POST("/accounts", CreateAccountBuilder(gwg_api_label))
	gin_router.GET("/accounts", ListAccountBuilder(gwg_api_label))
}

type RestaurantApiInterface interface {
	UploadItemImage(gin_context *gin.Context, id string)
	ListBills(gin_context *gin.Context, restaurantId string, status string, startAt int64, endAt int64, tableId string)
	GetRestaurant(gin_context *gin.Context, id string)
	DeleteRestaurant(gin_context *gin.Context, id string)
	UpdateRestaurant(gin_context *gin.Context, id string, gin_body PutRestaurantRequest)
	CreateItem(gin_context *gin.Context, id string, gin_body PutItemRequest)
	ListRestaurantItems(gin_context *gin.Context, id string)
	CreateTable(gin_context *gin.Context, id string, gin_body PutTableRequest)
	ListRestaurantTable(gin_context *gin.Context, id string)
	CreateBill(gin_context *gin.Context, id string, gin_body CreateBillRequest)
	ListRestaurants(gin_context *gin.Context)
	CreateRestaurant(gin_context *gin.Context, gin_body PutRestaurantRequest)
	ListTablesBills(gin_context *gin.Context, id string)
	FinishBill(gin_context *gin.Context, id string)
	CreatePrinter(gin_context *gin.Context, id string, gin_body PutPrinterRequest)
	ListPrinters(gin_context *gin.Context, id string)
	GetTable(gin_context *gin.Context, id string)
	UpdateTable(gin_context *gin.Context, id string, gin_body PutTableRequest)
	DeleteTable(gin_context *gin.Context, id string)
	GetItem(gin_context *gin.Context, id string)
	UpdateItem(gin_context *gin.Context, id string, gin_body PutItemRequest)
	DeleteItem(gin_context *gin.Context, id string)
	UpdatePrinter(gin_context *gin.Context, id string, gin_body PutPrinterRequest)
	DeletePrinter(gin_context *gin.Context, id string)
	GetBill(gin_context *gin.Context, id string)
}

func UploadItemImageBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.UploadItemImage(gin_context, id)
	}
}
func ListBillsBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		restaurantId := gin_context.Query("restaurantId")
		status := gin_context.Query("status")
		startAt := gin_context.Query("startAt")
		endAt := gin_context.Query("endAt")
		tableId := gin_context.Query("tableId")
		api.ListBills(gin_context, restaurantId, status, stringToInt64(startAt), stringToInt64(endAt), tableId)
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
func ListRestaurantItemsBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.ListRestaurantItems(gin_context, id)
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
func ListRestaurantTableBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.ListRestaurantTable(gin_context, id)
	}
}
func CreateBillBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		var createBillRequest CreateBillRequest
		if err := gin_context.ShouldBindJSON(&createBillRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.CreateBill(gin_context, id, createBillRequest)
	}
}
func ListRestaurantsBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		api.ListRestaurants(gin_context)
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
func ListTablesBillsBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.ListTablesBills(gin_context, id)
	}
}
func FinishBillBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.FinishBill(gin_context, id)
	}
}
func CreatePrinterBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		var putPrinterRequest PutPrinterRequest
		if err := gin_context.ShouldBindJSON(&putPrinterRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.CreatePrinter(gin_context, id, putPrinterRequest)
	}
}
func ListPrintersBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.ListPrinters(gin_context, id)
	}
}
func GetTableBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.GetTable(gin_context, id)
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
func UpdatePrinterBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		var putPrinterRequest PutPrinterRequest
		if err := gin_context.ShouldBindJSON(&putPrinterRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.UpdatePrinter(gin_context, id, putPrinterRequest)
	}
}
func DeletePrinterBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.DeletePrinter(gin_context, id)
	}
}
func GetBillBuilder(api RestaurantApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.GetBill(gin_context, id)
	}
}
func RestaurantApiInterfaceMounter(gin_router *gin.Engine, gwg_api_label RestaurantApiInterface) {
	gin_router.POST("/items/:id/image", UploadItemImageBuilder(gwg_api_label))
	gin_router.GET("/bills", ListBillsBuilder(gwg_api_label))
	gin_router.GET("/restaurants/:id", GetRestaurantBuilder(gwg_api_label))
	gin_router.DELETE("/restaurants/:id", DeleteRestaurantBuilder(gwg_api_label))
	gin_router.PUT("/restaurants/:id", UpdateRestaurantBuilder(gwg_api_label))
	gin_router.POST("/restaurants/:id/items", CreateItemBuilder(gwg_api_label))
	gin_router.GET("/restaurants/:id/items", ListRestaurantItemsBuilder(gwg_api_label))
	gin_router.POST("/restaurants/:id/tables", CreateTableBuilder(gwg_api_label))
	gin_router.GET("/restaurants/:id/tables", ListRestaurantTableBuilder(gwg_api_label))
	gin_router.POST("/tables/:id/orders", CreateBillBuilder(gwg_api_label))
	gin_router.GET("/restaurants", ListRestaurantsBuilder(gwg_api_label))
	gin_router.POST("/restaurants", CreateRestaurantBuilder(gwg_api_label))
	gin_router.GET("/tables/:id/bills", ListTablesBillsBuilder(gwg_api_label))
	gin_router.PUT("/bills/:id/finish", FinishBillBuilder(gwg_api_label))
	gin_router.POST("/restaurants/:id/printers", CreatePrinterBuilder(gwg_api_label))
	gin_router.GET("/restaurants/:id/printers", ListPrintersBuilder(gwg_api_label))
	gin_router.GET("/tables/:id", GetTableBuilder(gwg_api_label))
	gin_router.PUT("/tables/:id", UpdateTableBuilder(gwg_api_label))
	gin_router.DELETE("/tables/:id", DeleteTableBuilder(gwg_api_label))
	gin_router.GET("/items/:id", GetItemBuilder(gwg_api_label))
	gin_router.PUT("/items/:id", UpdateItemBuilder(gwg_api_label))
	gin_router.DELETE("/items/:id", DeleteItemBuilder(gwg_api_label))
	gin_router.PUT("/printers/:id", UpdatePrinterBuilder(gwg_api_label))
	gin_router.DELETE("/printers/:id", DeletePrinterBuilder(gwg_api_label))
	gin_router.GET("/bills/:id", GetBillBuilder(gwg_api_label))
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
