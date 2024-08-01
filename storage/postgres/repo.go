package postgres

import (
	"context"

	"github.com/temur-shamshidinov/Content_service/genproto/content_service"
)

type ContentRepoI interface{
	CreateContent (ctx context.Context, content *content_service.CreateContentReq) (*content_service.Content, error)
	
}