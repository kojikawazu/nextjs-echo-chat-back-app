package repositories_auth_users

import (
	"nextjs-echo-chat-back-app/middlewares"
	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"
)

// FetchAuthUsers は `auth_users` テーブルからすべてのユーザー情報を取得する。
func (r *AuthUsersRepositoryImpl) FetchAuthUsers() ([]models.AuthUsers, error) {
	query := `
		SELECT * FROM auth_users
	`
	rows, err := middlewares.Pool.Query(middlewares.Ctx, query)
	if err != nil {
		logger.ErrorLog.Printf("Failed to fetch auth_users: %v", err)
		return nil, err
	}

	var authUsers []models.AuthUsers

	// 結果をスキャンしてブログデータをリストに追加
	for rows.Next() {
		var authUser models.AuthUsers

		err := rows.Scan(
			&authUser.ID,
			&authUser.Name,
			&authUser.Email,
			&authUser.Password,
			&authUser.CreatedAt,
			&authUser.UpdatedAt,
		)
		if err != nil {
			logger.ErrorLog.Printf("Failed to scan auth_users: %v", err)
			return nil, err
		}

		authUsers = append(authUsers, authUser)
	}

	if rows.Err() != nil {
		logger.ErrorLog.Printf("Failed to fetch auth_users: %v", rows.Err())
		return nil, rows.Err()
	}

	logger.InfoLog.Printf("Fetched %d auth_users", len(authUsers))
	logger.InfoLog.Println("Fetched auth_users successfully")
	return authUsers, nil
}
