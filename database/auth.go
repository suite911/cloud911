import (
	"encoding/hex"
	"errors"

	"github.com/suite911/cloud911/types"

	"github.com/suite911/query911/query"

	"golang.org/x/crypto/sha3"
)

func Auth(auth *types.Auth) uint64 {
	if auth == nil {
		return 0
	}
	rid, id, dig, ent, req := auth.RowID, auth.ID, auth.Digest, auth.Entropy, auth.Request
	q := query.Query{DB: DB()}
	q.SQL = `SELECT "key", "flags" FROM "RegisteredUsers" WHERE `
	if rid > 0 {
		q.SQL += `"_ROWID_" = ? AND `
	}
	q.SQL += `"id" = ?;`
	if rid > 0 {
		q.Query(rid, id)
	} else {
		q.Query(id)
	}
	if !q.OK() {
		return 0
	}
	if !q.NextOrClose() {
		return 0
	}
	var key []byte
	var qFlags uint64
	q.ScanClose(&key, &qFlags)
	if !q.OK() {
		return 0
	}
	if hex.DecodedLen(dig) != 32 { // Anti-DoS
		return 0
	}
	digBytes, err := hex.DecodeString(dig)
	if err != nil {
		return 0
	}
	if hex.DecodedLen(ent) != 32 { // Anti-DoS
		return 0
	}
	entBytes, err := hex.DecodeString(ent)
	if err != nil {
		return 0
	}
	buf := append(entBytes, key)
	actual := sha3.Sum256(buf)
	if len(actual) != len(digBytes) { // Anti-DoS
		return 0
	}
	if actual != digBytes {
		return 0
	}
	return qFlags & req
}
