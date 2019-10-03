package sportmonks

import "context"

type Client interface {
	// Bookmakers sends a request and returns a slice of Bookmaker resource struct and supporting meta data.
	Bookmakers(ctx context.Context) ([]Bookmaker, *Meta, error)

	// BookmakerById sends a request and returns a single Bookmaker resource.
	BookmakerById(ctx context.Context, id int) (*Bookmaker, *Meta, error)

	// Continents sends a request and returns a slice of Continent struct and supporting meta data. The endpoint
	// used within this method is paginated, to select the required page use the 'page' method argument. Page
	// information including current page and total page are included within the Meta response.

	// Use the includes string slice to enrich the response data.
	Continents(ctx context.Context, page int, includes []string) ([]Continent, *Meta, error)

	// ContinentById sends a request and returns a single Continent struct.

	// Use the includes string slice to enrich the response data.
	ContinentById(ctx context.Context, id int, includes []string) (*Continent, *Meta, error)
}

