package types

const (
	Admin uint64 = 1 << iota
	Staff
	Minor
)

type RegisteredUser struct {
	RowID        int64  `json:"rowid"`
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"un"`
	HasPassword  bool   `json:"pw"`
	Registered   int64  `json:"regd"`
	Verified     int64  `json:"verd"`
	Balance      int64  `json:"bal"`
	Captcha1     int    `json:"captcha1"`
	Captcha2     int    `json:"captcha2"`
	Captcha3     int    `json:"captcha3"`
	Captchas     int    `json:"captchas"`
	Flags        uint64 `json:"flags"`
	EmergencyWho string `json:"emwho"`
	EmergencyHow string `json:"emhow"`
	EmergencyRel string `json:"emrel"`
}
