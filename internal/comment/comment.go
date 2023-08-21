package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("Failed to fetch comment by Id")
	ErrNotImplemented  = errors.New("Not implemented")
)

// Comment encapsulates article comments
type Comment struct {
	Id     string
	Slug   string
	Body   string
	Author string
}

type Storer interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service encapsulates our service
type Service struct {
	Store Storer
}

// NewService returns a pointer to a new Service
func NewService(store Storer) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Printf("Retrieving comment %s", id)

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
