package utils

import "time"

var (
	Size   = 50
	Offset = 0

	UUIDKey      = "uuid"
	OpenIDKey    = "open_id"
	UnionIDKey   = "union_id"
	AppIDKey     = "app_id"
	TokenKey     = "token"
	SignatureKey = "signature"

	RedisLoginKeyPrefix = "login:%s"

	AuthTimeout       = 7 * 24 * time.Hour
	RsaPrivateKeyPath = "config/rsa_private.key"
	RsaPublicKeyPath  = "config/rsa_public.key"
	AuthorizationKey  = "Authorization"
	BearerKey         = "Bearer"
)
