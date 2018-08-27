package types

import (
	"crypto/rand"
	"encoding/json"
	"encoding/hex"
)

const (
	Unlocked uint64 = 1 << iota // Account is not locked from logging in
	Adult                       // Account owner is an adult
	VerifiedEmail               // Account owner has verified their e-mail address
	VerifiedIdentity            // Account owner has verified their identity
	PaidBefore                  // Account owner has paid money before
	Staff                       // Accoupt owner is staff
	Admin                       // Accoupt owner is admin
)

type Auth struct {
	RowID   int64  `json:"rowid"` // Row ID for faster retrieval
	ID      int64  `json:"id"`    // Account ID
	Session int64  `json:"ses"`   // Session timestamp
	Request uint64 `json:"req"`   // Requested permissions
	Digest  string `json:"dig"`   // Digest
}

func NewAuth(rowid, id int64, key [32]byte, request uint64) (*Auth, error) {
	a, err := new(Auth).Init(rowid, id, key, request)
	return a, err
}

func (a *Auth) Init(rowid, id int64, key []byte, request uint64) (*Auth, error) {
	const lenEnt = 32
	a.RowID = rowid
	a.ID = id
	a.Request = request
	buf := make([]byte, lenEnt, lenEnt+len(key))
	if _, err := rand.Read(buf); err != nil {
		return nil, error
	}
	a.Entropy = hex.EncodeToString(buf[:lenEnt])
	buf = append(buf, key...)
	if len(buf) != lenEnt + len(key) {
		panic("Security")
	}
	dig := sha3.Sum256(buf)
	a.Digest = hex.EncodeToString(dig[:])
	return a, nil
}

type User struct {
	RowID               int64  `json:"rowid"`
	ID                  int64  `json:"id"`
	Email               string `json:"email"`
	Username            string `json:"un"`
	HasPassword         bool   `json:"pw"`
	Registered          int64  `json:"regd"`
	HasVerifiedEmail    bool   `json:"vemd"`
	HasVerifiedIdentity bool   `json:"vidd"`
	Session             int64  `json:"ses"`
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
