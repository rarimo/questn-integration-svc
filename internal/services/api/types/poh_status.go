package types

type ProofOfHumanityRequest struct {
	Address string `url:"address"`
}

type ProofOfHumanityResponseData struct {
	Result bool `json:"result"`
}

type ProofOfHumanityResponse struct {
	BaseResponse
	Data *ProofOfHumanityResponseData `json:"data"`
}
