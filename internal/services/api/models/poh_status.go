package models

import (
	"github.com/rarimo/questn-integration-svc/internal/services/api/types"
	"net/http"
)

func NewProofOfHumanityResponseError(w http.ResponseWriter, err error, status int) {
	render(w, &types.ProofOfHumanityResponse{
		BaseResponse: newBaseResponse(err, status),
		Data: &types.ProofOfHumanityResponseData{
			Result: false,
		},
	})
}

func NewProofOfHumanityResponse(w http.ResponseWriter, result bool) {
	render(w, &types.ProofOfHumanityResponse{
		Data: &types.ProofOfHumanityResponseData{
			Result: result,
		},
	})
}
