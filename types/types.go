package types

const (
	Unlocked uint64 = 1 << iota // Account is not locked from logging in
	Adult                       // Account owner is an adult
	VerifiedEmail               // Account owner has verified their e-mail address
	VerifiedIdentity            // Account owner has verified their identity
	PaidBefore                  // Account owner has paid money before
	Staff                       // Accoupt owner is staff
	Admin                       // Accoupt owner is admin
)

type APICall struct {
	API     string `json:"api"` // API to call
	Payload []byte `json:"dat"` // Payload (specific to the named API)
}

type Identity struct {
	RowID int64  `json:"rowid"` // Row ID for faster retrieval (optional)
	ID    int64  `json:"id"`    // Account ID
}

/*
type Register struct {
	RowID int64  `json:"rowid"` // Row ID for faster retrieval
	ID    int64  `json:"id"`    // Account ID
	Key   string `json:"key"`   // Hex-encoded argon2-hashed password
}

type Login struct {
	RowID   int64  `json:"rowid"` // Row ID for faster retrieval
	ID      int64  `json:"id"`    // Account ID
	Session int64  `json:"ses"`   // Session timestamp
	Request uint64 `json:"req"`   // Requested permissions
	Digest  string `json:"dig"`   // Digest
}
*/

type User struct {
	RowID               int64  `json:"rowid"`
	ID                  int64  `json:"id"`
	Email               string `json:"email"`
	Username            string `json:"un"`
	HasPassword         bool   `json:"pw"`
	Registered          int64  `json:"regd"`
	HasVerifiedEmail    bool   `json:"vemd"`
	HasVerifiedIdentity bool   `json:"vidd"`
	Balance             int64  `json:"bal"`
	Captcha1            int    `json:"captcha1"`
	Captcha2            int    `json:"captcha2"`
	Captcha3            int    `json:"captcha3"`
	Captchas            int    `json:"captchas"`
	Flags               uint64 `json:"flags"`
	EmergencyWho        string `json:"emwho"`
	EmergencyHow        string `json:"emhow"`
	EmergencyRel        string `json:"emrel"`
}
