package request

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	ID         uint   `json:"ID"`
	Username   string `json:"username"`
	RoleId     uint   `json:"roleId"` // 角色Id
	BufferTime int64  `json:"bufferTime"`
	jwt.RegisteredClaims
}
