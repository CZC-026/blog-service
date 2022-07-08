package service

import (
	"github.com/go-programming-tour-book/blog-service/internal/dao"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

type CountArticleRequest struct {
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"description" binding:"max=1000" `
	Content       string `form:"content" binding:"max=10000"`
	CoverImageUrl string `form:"coverimageurl"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"description" binding:"max=1000" `
	Content       string `form:"content" binding:"max=10000"`
	CoverImageUrl string `form:"coverimageurl"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,max=100"`
	Desc          string `form:"desc" binding:"max=1000" `
	Content       string `form:"content" binding:"max=10000"`
	CoverImageUrl string `form:"coverimageurl"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
}
type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"description" binding:"max=1000" `
	Content       string `form:"content" binding:"max=10000"`
	CoverImageUrl string `form:"coverimageurl"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

//func (svc *Service) CountArticle(param *CountArticleRequest) (int, error) {
//	return svc.dao.CountArticle(param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State)
//}

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		CreatedBy:     param.CreatedBy,
	})
}

//func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
//	article, err := svc.dao.CreateArticle(&dao.Article{
//		Title:         param.Title,
//		Desc:          param.Desc,
//		Content:       param.Content,
//		CoverImageUrl: param.CoverImageUrl,
//		State:         param.State,
//		CreatedBy:     param.CreatedBy,
//	})
//	if err != nil {
//		return err
//	}
//	err = svc.dao.CreateArticleTag(article.ID, param.TagID, param.CreatedBy)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteTArticle(param.ID)
}
func (svc *Service) DeleteArticleUnscoped(param *DeleteArticleRequest) error {
	return svc.dao.DeleteTArticle1(param.ID)
}
