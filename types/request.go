package types

type DrawConfigRequest struct {
	Id                  int    `form:"id"`
	Intro               string `form:"intro" binding:"required"`
	WinningInstructions int    `form:"winningInstructions"`
	Type                int    `form:"type" binding:"required"`
	ShowResult          int    `form:"showResult" binding:"required"`
	LimitNumOfWin       int    `form:"limitNumOfWin" binding:"required"`
	IsLimitNumOfWin     int    `form:"isLimitNumOfWin" binding:"required"`
	TotalPlay           int    `form:"totalPlay" binding:"required"`
	Times               int    `form:"times" binding:"required"`
	DeleteImg           int    `form:"DeleteImg"`
	PlayConfigs         string `form:"playConfigs" binding:"required"`
}

type PlayConfigs []struct {
	ID             int    `json:"id"`
	StartTime      int64  `json:"startTime"`
	EndTime        int64  `json:"endTime"`
	IsShowBanner   int    `json:"isShowBanner"`
	IsShowPrize    int    `json:"isShowPrize"`
	ShowType       int    `json:"showType"`
	Banner         string `json:"banner"`
	URL            string `json:"url"`
	WinningRate    int    `json:"winningRate"`
	SpecialSwitch  int    `json:"specialSwitch"`
	SpecialContent string `json:"specialContent"`
	Condition      int    `json:"condition"`
	JoinNum        int    `json:"joinNum"`
	CountDown      int    `json:"countDown"`
	PrizeType      int    `json:"prizeType"`
	ExchangeType   int    `json:"exchangeType"`
	ExpiredTime    string `json:"expiredTime"`
	PrizeConfigs   []struct {
		Name       string `json:"name" binding:"required"`
		PrizeAlias string `json:"prizeAlias"`
		Sum        int    `json:"sum" binding:"required"`
		Icon       string `json:"icon"`
		URL        string `json:"url"`
		ID         int    `json:"id"`
	} `json:"prizeConfigs"`
}
