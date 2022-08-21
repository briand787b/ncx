package event

// TicketDiscovered is a value object that gets emitted
// when a Ticket is discovered by the miner.
type TicketDiscovered struct {
	TicketID   string `json:"ticket_id"`
	LocationID string `json:"location_id"`
	Items      []struct {
		MenuItemID string `json:"menu_item_id"`
		Quantity   int    `json:"quantity"`
	} `json:"items"`
}
