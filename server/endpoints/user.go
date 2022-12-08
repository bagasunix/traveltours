package endpoints

import (
	"net/http"

	"github.com/bagasunix/traveltours/server/domains"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/utils"
	"github.com/gin-gonic/gin"
)

const (
	CREATE_USER = "CreateUser"
)

type UserEndpoint interface {
	CreateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	ListUser() gin.HandlerFunc
}

// type UserEndpoints struct {
// 	CreateUser Endpoint
// }

type userHandler struct {
	service domains.Service
}

// ListUser implements UserEndpoint
func (l *userHandler) ListUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.BaseList
		if err := ctx.Bind(&req); err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		dataAccount, err := l.service.ListUser(ctx, &req)
		if err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		ctx.JSON(http.StatusOK, dataAccount)
	}
}

// DeleteUser implements UserEndpoint
func (u *userHandler) DeleteUser() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := decodeByEntityIdEndpoint(g)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataAccount, err := u.service.DeleteUser(g, req.(*requests.EntityId))
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		g.JSON(http.StatusCreated, dataAccount)
	}
}

// CreateUser implements UserEndpoints
func (u *userHandler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.CreateUser
		if err := ctx.Bind(&req); err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		dataAccount, err := u.service.CreateUser(ctx, &req)
		if err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		ctx.JSON(http.StatusCreated, dataAccount)
	}
}

func NewUserEndpoint(svc domains.Service) UserEndpoint {
	return &userHandler{service: svc}
}

// func NewUserEndpoint(s domains.Service, mdw map[string][]Middleware) UserEndpoints {
// 	eps := UserEndpoints{}
// 	eps.CreateUser = makeCreateUserEndpoint(s)

// 	SetEndpoint(CREATE_USER, &eps.CreateUser, mdw)

// 	return eps
// }

// func makeCreateUserEndpoint(s domains.Service) Endpoint {
// 	return func(ctx *gin.Context) (res interface{}, err error) {
// 		var req *requests.CreateUser
// 		ctx.Set("request", req)
// 		return s.CreateUser(ctx, req)
// 	}
// return func(ctx *gin.Context) {
// 	var req requests.CreateUser
// 	if err := ctx.Bind(&req); err != nil {
// 		utils.EncodeError(ctx, err, ctx.Writer)
// 		return
// 	}
// 	dataAccount, err := s.CreateUser(ctx, &req)
// 	if err != nil {
// 		utils.EncodeError(ctx, err, ctx.Writer)
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, dataAccount)
// }
// }
