package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}
type Comment struct {
	ID        int
	Post_id   int
	Parent_id int
	Content   string
	AddTime   int64
	Visible   bool
	Replies   []Comment
}

func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

func (s *Storage) addComment(comment Comment) error {
	_, err := s.db.Exec(context.Background(), `INSERT INTO comments_post(post_id,parent_id,content,addTime)
	VALUES($1,$2,$3,$4,$5)`,
		&comment.Post_id,
		&comment.Parent_id,
		&comment.Content,
		&comment.AddTime,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) comments(post_id int) ([]Comment, error) {
	rows, err := s.db.Query(context.Background(), `SELECT id,post_id,parent_id,content,addTime,visible FROM comments_post WHERE post_id=$1 ORDER BY addTime DESC`,
		post_id)
	if err != nil {
		return nil, err
	}
	var comments []Comment
	for rows.Next() {
		var comment Comment
		err = rows.Scan(
			&comment.ID,
			&comment.Post_id,
			&comment.Parent_id,
			&comment.Content,
			&comment.AddTime,
			&comment.Visible,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
