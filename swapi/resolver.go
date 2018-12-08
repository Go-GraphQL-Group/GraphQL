package swapi

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Lookup(ctx context.Context) (*LookupQuery, error) {
	panic("not implemented")
}
func (r *queryResolver) Browse(ctx context.Context) (*BrowseQuery, error) {
	panic("not implemented")
}
func (r *queryResolver) Search(ctx context.Context) (*SearchQuery, error) {
	panic("not implemented")
}
