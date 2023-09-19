package models

import "gitlab.com/rarimo/questn-integration-svc/internal/services/api/types"

func newBaseResponse(err error, status int) types.BaseResponse {
	return types.BaseResponse{
		Error: &types.ResponseError{
			Code:    status,
			Message: err.Error(),
		},
	}
}
