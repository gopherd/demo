package module

// Service represents a gopherd service
type Service interface {
	// GetHammerComponent gets HammerComponent
	GetHammerComponent() HammerComponent
	// GetHornComponent gets HornComponent
	GetHornComponent() HornComponent
}

// HammerComponent represents hammer component
type HammerComponent interface {
	Weight() int // Weight of hammer
}

// HornComponent represents horn component
type HornComponent interface {
	Length() int // Length of horn
}
