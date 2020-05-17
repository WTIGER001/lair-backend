package internal

// NewDataStore - Creates the correct type of datastore
func NewDataStore() Datastore {
	return NewRAMDatastore()
}