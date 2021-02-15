package pedalboards

// Pedalboard represents a pedalboard stored in the db
type Pedalboard struct {
	ID        int     `json:"id,omitempty"`
	Brand     string  `json:"brand"`
	Name      string  `json:"name"`
	Width     float32 `json:"width"`
	Height    float32 `json:"height"`
	Image     string  `json:"image"`
	UpdatedAt string  `json:"-"`
	CreatedAt string  `json:"-"`
}

// Brand is the manufacturer associated with a pedalboard
type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
	Slug string `json:"slug"`
}

// Category represents a pedalboard category
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
