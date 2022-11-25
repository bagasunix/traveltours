package endpoints

import (
	"strconv"

	"github.com/bagasunix/traveltours/pkg/errors"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func decodeBaseListEndpoint(g *gin.Context) (request interface{}, err error) {
	r := g.Request
	query := r.URL.Query()
	qLimit := query["limit"]
	qKeywords := query["keyword"]

	requestBuilder := requests.NewBaseListBuilder()
	if !validation.IsEmpty(qKeywords) {
		requestBuilder.SetKeyword(qKeywords[0])
	}

	if validation.IsEmpty(qLimit) {
		requestBuilder.SetLimit(30)
		return requestBuilder.Build(), nil
	}

	limit, err := strconv.Atoi(qLimit[0])
	if err != nil {
		return nil, errors.ErrInvalidAttributes("limit")
	}

	if limit > 30 {
		requestBuilder.SetLimit(30)
		return requestBuilder.Build(), nil
	}

	return requestBuilder.SetLimit(int64(limit)).Build(), nil
}
