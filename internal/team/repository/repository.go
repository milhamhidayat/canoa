package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/huandu/go-sqlbuilder"
	"github.com/juju/errors"

	"soccer-api/internal/entity"
	"soccer-api/internal/team"
	"soccer-api/pkg/cursor"
)

type teamRepository struct {
	db *sql.DB
}

// New returns team repository
func New(db *sql.DB) team.Team {
	return &teamRepository{
		db: db,
	}
}

func (tr *teamRepository) Fetch(ctx context.Context, filter entity.QueryFilter) (teams []entity.Team, nextCursor string, err error) {
	sb := sq.NewSelectBuilder()
	sb.Select("id", "name", "founded", "stadium", "created_at").From("teams").OrderBy("created_at DESC")

	if filter.Num > 0 {
		sb.Limit(filter.Num)
	}

	if filter.Cursor != "" {
		decodedCursor, er := cursor.Decode(filter.Cursor)
		if err != nil {
			err = errors.Wrap(er, errors.New("decode team cursor"))
			return
		}
		sb.Where(sb.LessThan("created_at", decodedCursor))
	}

	q, args := sb.Build()
	rows, err := tr.db.QueryContext(ctx, q, args...)
	if err != nil {
		err = errors.Wrap(err, errors.New("select query"))
		return
	}

	for rows.Next() {
		t := entity.Team{}
		createdAt := time.Time{}
		err = rows.Scan(&t.ID, &t.Name, &t.Founded, &t.Stadium, &createdAt)
		if err != nil {
			err = errors.Wrap(err, errors.New("scan team row"))
			return
		}
		loc, _ := time.LoadLocation("Asia/Jakarta")
		t.CreatedAt = createdAt.In(loc)
		teams = append(teams, t)
	}

	nextCursor = filter.Cursor
	if len(teams) > 0 {
		createdAt := teams[len(teams)-1].CreatedAt
		nextCursor = cursor.Encode(createdAt)
	}

	return
}

func (tr *teamRepository) Get(ctx context.Context, ID int64) (team entity.Team, err error) {
	sb := sq.NewSelectBuilder()
	sb.Select("id", "name", "founded", "stadium")
	sb.From("teams").Where(sb.Equal("id", ID))

	query, args := sb.Build()
	row := tr.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&team.ID,
		&team.Name,
		&team.Founded,
		&team.Stadium,
	)
	if err == sql.ErrNoRows {
		err = entity.ErrNotFound
	}
	return

}

func (tr *teamRepository) Insert(ctx context.Context, t entity.Team) (team entity.Team, err error) {
	ib := sq.NewInsertBuilder()
	ib.InsertInto("teams").Cols("name", "founded", "stadium", "created_at")
	ib.Values(t.Name, t.Founded, t.Stadium, time.Now())

	q, args := ib.Build()
	res, err := tr.db.ExecContext(ctx, q, args...)
	if err != nil {
		err = errors.Wrap(err, errors.New("insert query"))
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		err = errors.Wrap(err, errors.New("inserted id"))
		return
	}
	t.ID = lastID
	team = t
	return
}
