package commits

import (
	"time"

	"github.com/chia-network/ecosystem-activity/internal/db"
)

// Commit represents all columns in one commit entry in the commits table
type Commit struct {
	ID     int
	RepoID int
	UserID int
	Date   time.Time
	SHA    string
	Notes  string
}

// SetNewRecord inserts one new record into the table
func SetNewRecord(c Commit) error {
	_, err := db.Exec(`INSERT INTO commits (repo_id,user_id,date,sha,notes) VALUES(?, ?, ?, ?, ?);`, c.RepoID, c.UserID, c.Date, c.SHA, c.Notes)
	return err
}
