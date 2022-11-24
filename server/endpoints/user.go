package endpoints

import (
	"github.com/bagasunix/traveltours/server/domains"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/gin-gonic/gin"
)

const (
	CREATE_USER = "CreateUser"
)

// type UserEndpoint interface {
// 	CreateUser() gin.HandlerFunc
// }

type UserEndpoints struct {
	CreateUser Endpoint
}

// type userHandler struct {
// 	service domains.Service
// }

// CreateUser implements UserEndpoints
// func (u *userHandler) CreateUser() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var req requests.CreateUser
// 		if err := ctx.Bind(&req); err != nil {
// 			utils.EncodeError(ctx, err, ctx.Writer)
// 			return
// 		}
// 		dataAccount, err := u.service.CreateUser(ctx, &req)
// 		if err != nil {
// 			utils.EncodeError(ctx, err, ctx.Writer)
// 			return
// 		}
// 		ctx.JSON(http.StatusCreated, dataAccount)
// 	}
// }

// func NewUserEndpoint(svc domains.Service) UserEndpoint {
// 	return &userHandler{service: svc}
// }

func NewUserEndpoint(s domains.Service, mdw map[string][]Middleware) UserEndpoints {
	eps := UserEndpoints{}
	eps.CreateUser = makeCreateUserEndpoint(s)

	SetEndpoint(CREATE_USER, &eps.CreateUser, mdw)

	return eps
}

func makeCreateUserEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request interface{}) (res interface{}, err error) {
		return s.CreateUser(ctx, request.(*requests.CreateUser))
	}
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
}
