package rabbitmq

const (
	// LocationDiscoveredQueue is the queue through which LocationDiscovered events travel
	LocationDiscoveredQueue = LocationDiscoveredExchange
	// LocationCreatedQueue is the queue through which LocationCreated events travel
	LocationCreatedQueue = LocationCreatedExchange
	// TicketDiscoveredQueue is the queue through which TicketDiscovered events travel
	TicketDiscoveredQueue = TicketDiscoveredExchange
)
