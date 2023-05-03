package model

type TickleDbEnvDetail struct {
	TableName string
	Env       string
	Namespace string
}

type NotificationField struct {
	Cname               string
	Series              string
	Entity              string
	Template            string
	From                string
	From_Name           string
	ReplyTo             string
	To                  []string
	DomainBase          string
	UserType            string
	Category            string
	NotificationChannel []string
}

type AuthenticatioResponse struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	OrgId      string `json:"orgId"`
	Timezone   int64  `json:"timezone"`
	CompanyId  string `json:"companyId"`
	SessionKey string `json:"sessionKey"`
	Delegation struct {
		Delegated bool `json:"sessionKey"`
	}
}

type EmailResponse struct {
	JobId string `json:"jobId"`
}

type EmailField struct {
	Cname    string
	Template string
	From     string
	FromName string
	ReplyTo  string
	To       []string
	Schedule string
}

type AuditField struct {
	UserId    string
	ModuleId  string
	SeriesId  string
	PagSize   int32
	MinTime   int64
	MaxTime   int64
	TimeStamp string
	AuditType int64
}

type EventField struct {
	ChannelId  int64
	Data       string
	TenantId   string
	Source     string
	Authorizer string
}

type EventChannelField struct {
	Name        string
	Project     string
	Parallelism int32
}
