package event

// LocationDiscovered defines the value object that gets
// emitted when a location is discovered by the miner. This
// does NOT necessarily mean that the location is new to
// the system.
type LocationDiscovered struct {
	ID string `json:"id"`
}

// LocationCreated defines the value object that gets
// emitted when a new Location is added to the database.
// This is a key differentiator between the LocationDiscovered
// event, which can be fired any time a Location is seen,
// even when that location is already known to the system
type LocationCreated struct {
	ID string `json:"id"`
}
