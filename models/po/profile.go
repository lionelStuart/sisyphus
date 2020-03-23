package po

type Profile struct {
	// Model

	ID         string `gorm:"primary_key" json:"id"`
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
	DeletedOn  int    `json:"deleted_on"`

	// ID 			int64 	  `gorm:"primary_key"`
	Nickname string `json:"nickname"`
	Age      int8   `json:"age"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
}
