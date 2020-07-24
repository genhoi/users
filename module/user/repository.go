package user

import (
	"database/sql"
	"strings"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAllByQuery(query string, limit uint16) ([]User, error) {
	query = strings.TrimSpace(query)
	tokens := strings.Split(query, " ")
	tokens = r.removeEmpty(tokens)
	tokens = r.addPerfix(tokens)
	tsQuery := r.toQuery(tokens)

	sqlQuery := `
SELECT id, login, full_name FROM users
WHERE 
	to_tsvector('russian', full_name) @@ to_tsquery('russian', '` + tsQuery + `') OR
	full_name ILIKE $1
LIMIT $2;
`

	rows, err := r.db.Query(sqlQuery, query+"%", limit)
	if err != nil {
		return nil, err
	}
	users := make([]User, 0, limit)

	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Login, &u.FullName)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *Repository) addPerfix(tokens []string) []string {
	prefixed := make([]string, 0, len(tokens))
	for _, t := range tokens {
		if t == "" {
			continue
		}
		prefixed = append(prefixed, t+":*")
	}

	return prefixed
}

func (r *Repository) toQuery(tokens []string) string {
	return strings.Join(tokens, " & ")
}

func (r *Repository) removeEmpty(tokens []string) []string {
	res := make([]string, 0, len(tokens))
	for _, t := range tokens {
		if t == "" {
			continue
		}
		res = append(res, t)
	}

	return res
}
