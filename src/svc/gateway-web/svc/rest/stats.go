package rest

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/briand787b/ncx/src/svc/server/internal/model/stats"
// )

// // MostOrderedItemResponse represents the response object for MostOrderedItem requests
// type MostOrderedItemResponse struct {
// 	MenuItemID string `json:"menu_item_id"`
// 	Quantity   int    `json:"quantity"`

// 	m *stats.MostOrderedItem
// }

// // NewMostOrderedItemResponse creates a new MostOrderedItemResponse
// func NewMostOrderedItemResponse(mm *stats.MostOrderedItem) *MostOrderedItemResponse {
// 	resp := &MostOrderedItemResponse{m: mm}

// 	return resp
// }

// // Render processes a MostOrderedItemResponse before rendering in HTTP response
// func (m *MostOrderedItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
// 	if m.m == nil {
// 		return errors.New("No MostOrderedItem exists in MostOrderedItem response")
// 	}

// 	m.MenuItemID = m.m.MenuItemID
// 	m.Quantity = m.m.Quantity
// 	return nil
// }
