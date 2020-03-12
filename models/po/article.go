package po

type Article struct {
	Model

	TagID int `json:"tag_id" mapstructure:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url" mapstructure:"cover_image_url"`
	CreatedBy     string `json:"created_by" mapstructure:"created_by"`
	OwnerUid      string `json:"owner_uid" mapstructure:"owner_uid"`
	ModifiedBy    string `json:"modified_by" mapstructure:"modified_by"`
	State         int    `json:"state"`
}
