package utils

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

	RsaPrivateKeyPath = "config/rsa_private.key"
	RsaPublicKeyPath  = "config/rsa_public.key"
	AuthorizationKey  = "Authorization"
	BearerKey         = "Bearer"
)
