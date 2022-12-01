package utils

import (
	"github.com/bagasunix/traveltours/pkg/errors"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func DecodeAssignPermissionsToRoleEndpoint(g *gin.Context) (request interface{}, err error) {
	vars := g.Param("id")
	if vars == "" {
		return nil, errors.ErrInvalidAttributes("id")
	}
	var uuidId uuid.UUID
	if uuidId, err = uuid.FromString(vars); err != nil {
		return nil, errors.ErrInvalidAttributes("id")
	}

	var req requests.AssignPermissionsToRole
	if err = g.Bind(&req); err != nil {
		return nil, errors.ErrInvalidAttributes("permissions")
	}
	req.RoleId = uuidId
	return &req, nil
}

func DecodeRemovePermissionsFromRoleEndpoint(g *gin.Context) (request interface{}, err error) {
	vars := g.Param("id")
	if vars == "" {
		return nil, errors.ErrInvalidAttributes("id")
	}

	var uuidId uuid.UUID
	if uuidId, err = uuid.FromString(vars); err != nil {
		return nil, errors.ErrInvalidAttributes("id")
	}
	var req requests.RemovePermissionsFromRole
	if err = g.Bind(&req); err != nil {
		return nil, errors.ErrInvalidAttributes("permissions")
	}
	req.RoleId = uuidId
	return &req, nil
}
