package data

import (
	"context"
	"database/sql"
	"time"
)

type PermissionModel struct {
	DB *sql.DB
}

type Permissions []string

func (m PermissionModel) GetAllForUser(userID int64) (Permissions, error) {
	sqlQuery := `SELECT permissions.code
				FROM permissions
				INNER JOIN users_permissions ON users_permissions.permission_id = permissions.id
				INNER JOIN users ON users_permissions.user_id = users.id
				WHERE users.id = $1
				`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, sqlQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions Permissions

	for rows.Next() {
		var permission string

		err := rows.Scan(&permission)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return permissions, nil
}
