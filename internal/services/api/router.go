package api

import (
	"gitlab.com/rarimo/questn-integration-svc/internal/services/api/handlers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *api) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxConfig(s.cfg),
			handlers.CtxSubgraph(s.cfg.Subgraph()),
		),
	)

	r.Route("/questn", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/poh_status", handlers.ProofOfHumanityStatus)
		})
	})

	return r
}
