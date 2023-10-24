package requests

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/questn-integration-svc/internal/services/api/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

func NewProofOfHumanityStatusRequest(r *http.Request) (*types.ProofOfHumanityRequest, error) {
	request := types.ProofOfHumanityRequest{}
	err := urlval.DecodeSilently(r.URL.Query(), &request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode url")
	}

	errs := validation.Errors{}
	errs["address"] = validation.Validate(
		request.Address,
		validation.Required, hexValidator,
	)

	if errs.Filter() != nil {
		return nil, errs.Filter()
	}

	return &request, nil
}
