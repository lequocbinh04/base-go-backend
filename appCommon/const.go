package appCommon

type DbType int

const (
	CurrentUser = "user"

	DBMain           = "postgres"
	PasetoProvider   = "paseto"
	PluginRedis      = "redis"
	PluginLocker     = "locker"
	PluginGrpcServer = "grpc-server"
)

const (
	DbTypeUser = iota
	DbTypeFingerprint
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

type TokenPayload struct {
	UId   int64  `json:"user_id"`
	URole string `json:"role"`
}

func (p TokenPayload) UserId() int64 {
	return p.UId
}

func (p TokenPayload) Role() string {
	return p.URole
}
