package postgres

import (
	"context"
	"log"

	"github.com/temur-shamshidinov/Content_service/genproto/content_service"
)

type ContentRepo struct{}

func NewContentRepo () ContentRepoI {
	return &ContentRepo{}
}

func (c *ContentRepo) CreateContent(ctx context.Context, content *content_service.CreateContentReq) (*content_service.Content, error) {
	
	log.Println("sent data:", content)

	return nil, nil
}