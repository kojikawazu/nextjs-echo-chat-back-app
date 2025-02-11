package utils_uuid

import "github.com/google/uuid"

// IsUUID は指定された文字列が UUID 形式かどうかを判定する
func IsUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
