package repository

import (
	"news-app-be23/internal/features/articles"

	"gorm.io/gorm"
)

type ArticleModel struct {
	db *gorm.DB
}

func NewArticleModel(connection *gorm.DB) articles.Query {
	return &ArticleModel{
		db: connection,
	}
}

func (am *ArticleModel) InsertArticle(newArticle articles.Article) error {
	return am.db.Create(&newArticle).Error
}

func (am *ArticleModel) GetAllArticles() ([]articles.Article, error) {
	var articleList []articles.Article
	err := am.db.Find(&articleList).Error
	return articleList, err
}

func (am *ArticleModel) GetArticleByID(id uint) (*articles.Article, error) {
	var article articles.Article
	err := am.db.First(&article, id).Error
	return &article, err
}

func (am *ArticleModel) UpdateArticle(updatedArticle articles.Article) error {
	return am.db.Save(&updatedArticle).Error
}

func (am *ArticleModel) DeleteArticle(id uint) error {
	return am.db.Delete(&articles.Article{}, id).Error // Soft delete
}
