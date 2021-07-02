package service

type PostIsuConditionRequest struct {
	IsSitting bool   `json:"is_sitting"`
	Condition string `json:"condition"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type JIAServiceRequest struct {
	TargetIP   string `json:"target_ip"`
	TargetPort int    `json:"target_port"`
	IsuUUID    string `json:"isu_uuid"`
}

type PostIsuRequest struct {
	JIAIsuUUID string `json:"jia_isu_uuid"`
	IsuName    string `json:"isu_name"`
}

type PutIsuRequest struct {
	Name string `json:"name"`
}

type GetIsuSearchRequest struct {
	Name           *string
	Character      *string
	CatalogName    *string
	MinLimitWeight *uint
	MaxLimitWeight *uint
	CatalogTags    *string
	Page           *uint
}

type GetIsuConditionRequest struct {
	StartTime        *uint64
	CursorEndTime    uint64
	CursorJIAIsuUUID string
	ConditionLevel   string
	Limit            *uint
}
