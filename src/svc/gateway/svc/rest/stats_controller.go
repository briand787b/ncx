package rest

// import (
// 	"net/http"

// 	"github.com/briand787b/ncx/src/svc/pkg/cserr"
// 	"github.com/briand787b/ncx/src/svc/pkg/cslog"
// 	"github.com/briand787b/ncx/src/svc/server/internal/model/stats"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/render"
// )

// // StatsController controls the flow of Stats.Stats requests
// type StatsController struct {
// 	ss stats.Store
// 	l  cslog.Logger
// }

// // NewStatsController returns a new StatsController instance
// func NewStatsController(l cslog.Logger, ss stats.Store) *StatsController {
// 	return &StatsController{l: l, ss: ss}
// }

// func (c *StatsController) handleGetMostOrderedItemByLocationID(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	lID := chi.URLParam(r, "location_id")
// 	if lID == "" {
// 		render.Render(w, r, cserr.NewInternalServerHTTPError(ctx, "could not get location_id url param", c.l))
// 		return
// 	}

// 	moi, err := stats.GetMostOrderedItemByLocation(ctx, lID, c.ss)
// 	if err != nil {
// 		render.Render(w, r, cserr.NewHTTPErrorFromError(ctx, err, "could not get location by ID", c.l))
// 		return
// 	}

// 	render.Render(w, r, NewMostOrderedItemResponse(moi))
// }
