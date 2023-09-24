package data

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestLocations(t *testing.T) {
	assert := require.New(t)
	dns := fmt.Sprintf("postgres://%s:%s@db-test/%s?sslmode=disable", os.Getenv("DB_USER_TEST"),
		os.Getenv("DB_PASS_TEST"), os.Getenv("DB_NAME_TEST"))
	db, err := sql.Open("postgres", dns)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		t.Fatal(err)
	}

	m := NewModels(db)

	t.Run("insert", func(t *testing.T) {
		loc := &Location{
			Title:       "Test Location",
			Description: "Test Description",
			Address:     "Address",
		}
		err := m.Locations.Insert(loc)
		assert.NoError(err)
		assert.Equal(int64(1), loc.ID)
	})

	t.Run("select all", func(t *testing.T) {
		locs, err := m.Locations.SelectAll()
		assert.NoError(err)
		assert.Equal(1, len(locs))
	})

}
