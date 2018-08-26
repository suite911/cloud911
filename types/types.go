package types

import (
	"crypto/rand"
	"encoding/json"
	"encoding/hex"
)

const (
	Self uint64 = 1 << iota
	Adult
	Staff
	Admin
)

type Auth struct {
	RowID   int64  `json:"rowid"`
	ID      int64  `json:"id"`
	Digest  string `json:"dig"`
	Entropy string `json:"ent"`
	Request uint64 `json:"req"`
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
