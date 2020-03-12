package po

type Profile struct {
	Model

	// ID 			int64 	  `gorm:"primary_key"`
	Nickname string `json:"nickname"`
	Age      int8   `json:"age"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
}
