package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models"
	metakit "github.com/Nicolas-ggd/gorm-metakit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"
)

type RepositorySuite struct {
	db *sql.DB
	suite.Suite
}

// Initialize database mock
func (rs *RepositorySuite) InitDBMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func (rs *RepositorySuite) TestInsert() {
	rs.Run("Test Insert Success", func() {
		db, mock := rs.InitDBMock()
		defer db.Close()

		// initialize notification repository with db
		r := NotificationRepository{DB: db}

		// bind notification
		n := &models.Notification{
			Type:    "info",
			Time:    time.Now(),
			Message: "This is a test notification",
		}

		rows := sqlmock.NewRows([]string{"id", "type", "time", "message"}).
			AddRow(1, n.Type, n.Time, n.Message)

		mock.ExpectQuery(`INSERT INTO notifications \(type, time, message\)`).
			WithArgs(n.Type, n.Time, n.Message).
			WillReturnRows(rows)

		// call actual method
		result, err := r.Insert(n)

		assert.NoError(rs.T(), err)
		assert.NotNil(rs.T(), result)
		assert.Equal(rs.T(), uint(1), result.ID)
		assert.Equal(rs.T(), n.Type, result.Type)
		assert.Equal(rs.T(), n.Time, result.Time)
		assert.Equal(rs.T(), n.Message, result.Message)

		err = mock.ExpectationsWereMet()
		assert.NoError(rs.T(), err)
	})
}

func (rs *RepositorySuite) TestList() {
	rs.Run("Test List Success", func() {
		db, mock := rs.InitDBMock()
		defer db.Close()

		r := NotificationRepository{DB: db}

		// initialize metakit for list function
		meta := &metakit.Metadata{
			Sort:          "id",
			SortDirection: "asc",
		}

		mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM notifications").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		mock.ExpectQuery("SELECT \\* FROM notifications").
			WillReturnRows(sqlmock.NewRows([]string{"id", "type", "time", "message"}).
				AddRow(1, "info", "2023-07-15 12:00:00", "This is a test notification"))

		// call actual List method
		result, meta, err := r.List(meta)

		assert.NoError(rs.T(), err)
		assert.NotNil(rs.T(), result)
		assert.Equal(rs.T(), int64(1), meta.TotalRows)
		assert.Equal(rs.T(), 1, len(*result))
		assert.Equal(rs.T(), uint(0), (*result)[0].ID)
		assert.Equal(rs.T(), "", (*result)[0].Type)
		assert.Equal(rs.T(), time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), (*result)[0].Time)
		assert.Equal(rs.T(), "", (*result)[0].Message)

		err = mock.ExpectationsWereMet()
		assert.NoError(rs.T(), err)
	})
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
