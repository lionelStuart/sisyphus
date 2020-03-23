package po

type Auth struct {
	// Model

	ID         string `gorm:"primary_key" json:"id"`
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
	DeletedOn  int    `json:"deleted_on"`

	//ID       int    `gorm:"primary_key" json:"id"`
	// Uid      string `json:"uid"` //base32
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	State    int    `json: "state"`

	// ProfileID  int `json:"profile_id" gorm:"index"`
	Profile Profile `json:"profile"`
}
