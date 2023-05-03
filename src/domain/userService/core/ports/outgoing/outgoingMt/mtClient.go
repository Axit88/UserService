package outgoing

import (
	"context"

	"github.com/Axit88/UserService/src/domain/userService/core/model"
	authCommon "github.com/MindTickle/content-protos/pb/common"
	AuditLog "github.com/MindTickle/platform-protos/pb/auditlogsservice"
	AuditCommon "github.com/MindTickle/platform-protos/pb/common"
	"github.com/MindTickle/platform-protos/pb/event"
	sqlStore "github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
)

type TickleDbCreateTableClient interface {
	CreateTable(dbDetail model.TickleDbEnvDetail) error
}

type TickleDbInsertClient interface {
	InsertRow(id string, field model.User, url string, tableName string, reqContext sqlStore.RequestContext, authMeta sqlStore.AuthMeta) error
}

type AuthorizationClient interface {
	GetCompnanyRolePermission(url string, companyId string, reqMeta authCommon.RequestMeta, recMeta authCommon.RecordMeta) error
	GetUserRolePermission(url string, userId string, reqMeta authCommon.RequestMeta, recMeta authCommon.RecordMeta) error
}

type EventClient interface {
	CreateEvents(ctx context.Context, url string, channelId int64, eventData model.EventField) (*event.CreateEventsResponse, error)
}

type EventChannelClient interface {
	CreateEventChannel(ctx context.Context, url string, channelData model.EventChannelField) error
}

type AuditLogClient interface {
	AddLog(myClient AuditLog.AuditLogsServiceClient, url string, reqM AuditCommon.RequestMeta, schemaField []AuditLog.IngestField, field model.AuditField) error
}
