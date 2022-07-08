package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

type Article struct { //dao层将数据关联起来
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

//func (d *Dao) CountArticle(title string, desc string, content string, coverimageurl string, state uint8) (int, error) {
//	article := model.Article{
//		Title:         title,
//		Desc:          desc,
//		Content:       content,
//		CoverImageUrl: coverimageurl,
//		State:         state,
//	}
//	return article.Count(d.engine)
//}

func (d *Dao) GetArticleList(title string, desc string, content string, coverimageurl string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverimageurl,
		State:         state,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

//func (d *Dao) CreateArticle(title string, desc string, content string, coverimageurl string, state uint8, createdBy string) (*model.Article, error) {
//	article := model.Article{
//		Title:         title,
//		Desc:          desc,
//		Content:       content,
//		CoverImageUrl: coverimageurl,
//		State:         state,
//		Model:         &model.Model{CreatedBy: createdBy},
//	}
//	return article.Create(d.engine)
//}
func (d *Dao) CreateArticle(param *Article) error {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title string, desc string, content string, coverimageurl string, state uint8, modifiedBy string) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if title != "" {
		values["title"] = title
	}
	if desc != "" {
		values["desc"] = desc
	}
	if content != "" {
		values["content"] = content
	}
	if coverimageurl != "" {
		values["coverimageurl"] = coverimageurl
	}
	return article.Update(d.engine, values)
}

func (d *Dao) DeleteTArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}

func (d *Dao) DeleteTArticle1(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete1(d.engine)
}
