//+build integration

package rdb

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tx, err := testRepo.NewWriteTx(context.Background())
	require.NoError(t, err)
	defer tx.Commit()

	type to struct {
		ID   int64
		Time Time
	}

	zero := Time{}
	t.Log(zero.time.Unix())
	res, err := tx.Exec(`INSERT INTO test.time (time) VALUES (?)`, zero)
	require.NoError(t, err)
	zid, _ := res.LastInsertId()

	now := FromTime(time.Now())
	t.Log(now.time.Unix())
	res, err = tx.Exec(`INSERT INTO test.time (time) VALUES (?)`, now)
	require.NoError(t, err)
	nid, _ := res.LastInsertId()

	zo := new(to)
	err = tx.Get(zo, `SELECT id, time FROM test.time WHERE id = ?`, zid)
	assert.NoError(t, err)

	assert.Equal(t, zero.ToTime().Unix(), zo.Time.ToTime().Unix())

	no1 := new(to)
	err = tx.Get(no1, `SELECT id, time FROM test.time WHERE id = ?`, nid)
	assert.NoError(t, err)
	assert.Equal(t, now.ToTime().Unix(), no1.Time.ToTime().Unix())

	no2 := new(to)
	q := fmt.Sprintf(`SELECT id, time FROM test.time WHERE id = %d`, nid)
	stm, err := tx.Preparex(q)
	assert.NoError(t, err)
	err = stm.Get(no2)
	assert.NoError(t, err)
	assert.Equal(t, now.ToTime().Unix(), no2.Time.ToTime().Unix())
}
