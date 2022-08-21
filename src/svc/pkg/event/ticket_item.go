package event

// TicketItemCreated is a value object that gets emitted
// when a TicketItem is inserted into the datbase
type TicketItemCreated struct {
	LocationID string `json:"location_id"`
	TicketID   string `json:"ticket_id"`
	MenuItemID string `json:"menu_item_id"`
	Quantity   int    `json:"quantity"`
}
