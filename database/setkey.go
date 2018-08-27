package database

import (
	"errors"

	"github.com/suite911/cloud911/types"

	"github.com/suite911/query911/query"

	"golang.org/x/crypto/sha3"
)

func SetKey(sk *SetKey) (*types.Session, error) {
	if sk == nil {
		return nil, errors.New("nil")
	}
	rid, id, key := sk.RowID, sk.ID, sk.Key
	if len(key) != 32 {
		return nil, errors.New("bad key")
	}
	dig := sha3.Sum256(key)
	now := time.Now().UTC().Unix()
	q := query.Query{DB: DB()}
	q.SQL = `UPDATE "Users" ` +
		`SET "key" = ?, "ses" = ? ` +
		`WHERE `
	if rid > 0 {
		q.SQL += '"_ROWID_" = ? && '
	}
	q.SQL += '"id" = ? && "key" = NULL;'
	if rid > 0 {
		q.Exec(dig, now, rid, id)
	} else {
		q.Exec(dig, now, id)
	}
	if err := q.Error; err != nil {
		return nil, err
	}
	s := new(types.Session)
	s.RowID = rid
	s.ID = id
	s.LoggedInAt = now
	return s, nil
}
