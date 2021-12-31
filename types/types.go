package types

type CloseDrawRequest struct {
	Id      int64 `form:"id"`
    IsPrize int64 `form:"isPrize"`
}

type DrawingRequest struct {
	Id       int64  `form:"id"`
    PlayId   int64  `form:"playId"`
	Type     string `form:"type"`
    SourceId int64  `form:"sourceId"`
    Uin      int64  `form:"uin"`
	Content  string `form:"content"`
}

type SetBindRequest struct {
	Id        int64  `form:"id"`
	IncludeId int64  `form:"includeId"`
	Type      string `form:"type"`
	Status    int    `form:"status"`
}

type GetBindRequest struct {
	IncludeId int64  `form:"includeId"`
	Type      string `form:"type"`
}

type GetConfigListsRequest struct {
	Type      string `form:"type"`
	ChannelId int64  `form:"channelId"`
	Page      int  `form:"page,default=1"`
	Num       int  `form:"num,default=10"`
}

type ListRequest struct {
	Title       string `form:"title,optional"`
	Status      int    `form:"status,default=-1"`
	CreateSTime int64  `form:"createSTime,default=0"`
	CreateETime int64  `form:"createETime,default=0"`
	Page        int  `form:"page,default=1"`
	Num         int  `form:"num,default=10"`
}

type BindDrawRequest struct {
	Type      string `form:"type"`
	ChannelId int64  `form:"channelId"`
	id        int64  `form:"id"`
}

type InfoRequest struct {
	Id int64 `form:"id"`
}

type RecordRequest struct {
	Id        int64  `form:"id"`
	UserNick  string `form:"userNick,optional"`
	Phone     string `form:"phone,optional"`
	StartTime int64  `form:"startTime,default=0"`
	EndTime   int64  `form:"endTime,default=0"`
	Page      int64  `form:"page,default=1"`
	Num       int64  `form:"num,default=10"`
}

type Response struct {
	Code         int         `json:"code"`
	Data         interface{} `json:"data"`
	ErrorCode    int         `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
}

type ConfigRequest struct {
	Id             int64  `form:"id"`
	Title          string `form:"title"`
	Logo           string `form:"logo"`
	BannerLink     string `form:"bannerLink"`
	Intro          string `form:"intro"`
	Question       string `form:"question"`
	DeleteImg      string `form:"deleteImg"`
	URL            string `form:"url"`
	StartTime      int64  `form:"startTime"`
	EndTime        int64  `form:"endTime"`
	IsShowDeadline int64  `form:"isShowDeadline"`
}

