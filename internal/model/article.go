package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

//func (a Article) Count(db *gorm.DB) (int, error) {
//	var count int
//	//是否具体查某一篇文章
//	if a.Title != "" {
//		db = db.Where("Title=?", a.Title)
//	}
//	db = db.Where("state=?", a.State)
//	if err := db.Model(&a).Where("is_del=?", 0).Count(&count).Error; err != nil {
//		return 0, err
//	}
//	return count, nil
//}

// List 查询一系列文章列表
func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db = db.Where("title=?", a.Title)
	}
	db = db.Where("state=?", a.State)
	if err = db.Where("is_del=?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? AND state = ? AND is_del = ?", a.ID, a.State, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}

	return article, nil
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

//func (a Article) Create(db *gorm.DB) (*Article, error) {
//	if err := db.Create(&a).Error; err != nil {
//		return nil, err
//	}
//
//	return &a, nil
//}
func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("id=? AND is_del=?", a.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id=? AND is_del=?", a.Model.ID, 0).Delete(&a).Error
}
func (a Article) Delete1(db *gorm.DB) error {
	return db.Where("id=? AND is_del=?", a.Model.ID, 0).Unscoped().Delete(&a).Error
}
