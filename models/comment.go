package models

//
//type Comment struct {
//	ID        int    `gorm:"primary_key" json:"id"`
//	ArticleId int    `json:"article_id"`
//	Content   string `json:"content"`
//	CreatedBy string `json:"created_by"`
//}
//
//// AddComment Add a Tag
//func AddCommet(id int, content string, createdBy string) error {
//	tag := Comment{
//		ID:        id,
//		Content:   content,
//		CreatedBy: createdBy,
//	}
//
//	if err := db.Create(&tag).Error; err != nil {
//		return err
//	}
//	return nil
//}
