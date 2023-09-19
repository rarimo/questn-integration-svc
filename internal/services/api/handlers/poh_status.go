package handlers

import (
	"gitlab.com/rarimo/questn-integration-svc/internal/services/api/models"
	"gitlab.com/rarimo/questn-integration-svc/internal/services/api/requests"
	"net/http"
)

func ProofOfHumanityStatus(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewProofOfHumanityStatusRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		models.NewProofOfHumanityResponseError(w, err, http.StatusBadRequest)
		return
	}

	status, err := Subgraph(r).UserStatus(request.Address)
	if err != nil {
		Log(r).WithError(err).Error("failed to get status")
		models.NewProofOfHumanityResponseError(w, err, http.StatusInternalServerError)
		return
	}

	models.NewProofOfHumanityResponse(w, status)
}
