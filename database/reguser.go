package database

import (
	"errors"

	"github.com/suite911/cloud911/types"

	"github.com/suite911/query911/query"
)

func QueryRegisteredUser(auth *types.Auth, id int64) (*types.RegisteredUser, error) {
	aid := auth.ID
	priv := Auth(auth)
	q := query.Query{DB: DB()}
	q.SQL = `SELECT ` +
		`"_ROWID_", "email", "un", "pw", "regd", "verd", "bal", ` +
		`"conload", "conchange", "consubmit", "captcha", ` +
		`"flags", "emwho", "emhow", "emrel" ` +
		`FROM "RegisteredUsers" WHERE `
	if rid > 0 {
		q.SQL += `"_ROWID_" = ? AND `
	}
	q.SQL += `"id" = ?;`
	if rid > 0 {
		q.Query(rid, id)
	} else {
		q.Query(id)
	}
	if err := q.ErrorLogNow(); err != nil {
		return nil, err
	}
	if !q.NextOrClose() {
		return nil, q.ErrorLogNow() // probably nil, which is what we want: it means no result
	}
	ru := types.RegisteredUser
	var pw []byte
	q.ScanClose(
		&ru.RowID, &ru.Email, &ru.Username, &pw, &ru.Registered, &ru.Verified, &ru.Balance,
		&ru.Captcha1, &ru.Captcha2, &ru.Captcha3, &ru.Captchas,
		&ru.Flags, &ru.EmergencyWho, &ru.EmergencyHow, &ru.EmergencyRel,
	)
	if err := q.ErrorLogNow(); err != nil {
		return nil, err
	}
	result := new(types.RegisteredUser)
	result.RowID = ru.RowID
	result.ID = id
	result.HasPassword = len(pw) > 0
	result.Registered = ru.Registered
	result.Verified = ru.Verified
	authAsStaff := perm.Any(types.Admin|types.Staff)
	if authAsStaff || id == aid {
		result.Email = ru.Email
		result.Username = ru.Username
		result.Balance = ru.Balance
		result.EmergencyWho = ru.EmergencyWho
		result.EmergencyHow = ru.EmergencyHow
		result.EmergencyRel = ru.EmergencyRel
		result.Flags = ru.Flags
		if authAsStaff {
			result.Captcha1 = ru.Captcha1
			result.Captcha2 = ru.Captcha2
			result.Captcha3 = ru.Captcha3
			result.Captchas = ru.Captchas
		}
	}
	return result, nil
}
