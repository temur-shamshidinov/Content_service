package service

import (
	"context"

	"github.com/temur-shamshidinov/Content_service/genproto/content_service"
	"github.com/temur-shamshidinov/Content_service/storage"
)

type ContentService struct {
	Storage storage.StorageI
	content_service.UnimplementedContentServiceServer
}

func NewContentServices (storage storage.StorageI) *ContentService {
	return &ContentService{Storage: storage}
}

func (c * ContentService) CreateContent( ctx context.Context, req *content_service.CreateContentReq) (*content_service.Content, error){
	
	c.Storage.GetContentRepo().CreateContent(ctx,req)
	
	return nil, nil
}