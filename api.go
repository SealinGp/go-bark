package gobark

type Dog interface {
	Bark(req *BarkRequest) error
}

type BarkRequest struct {
	BarkRequestOptions
	TextBarkRequest *TextBarkRequest `json:"text_bark_request,omitempty"` //普通消息
	RingBarkRequest *RingBarkRequest `json:"ring_bark_request,omitempty"` //铃声消息
}

type BarkRequestOptions struct {
	Icon              string `json:"icon,omitempty"`               //自定义推送图标,图标将替换默认Bark图标,会自动缓存在客户端,https://xxx/xx.jpg
	Group             string `json:"group,omitempty"`              //消息分组名,同一个分组下的消息将会叠加显示
	Level             Level  `json:"level,omitempty"`              //消息时效性配置
	Url               string `json:"url,omitempty"`                //点击通知可跳转的url
	Badge             int    `json:"badge,omitempty"`              //通知角标
	Copy              string `json:"copy,omitempty"`               //下啦推送,锁屏界面左滑查看推送时,可选择复制推送内容,携带该参数时,将只复制copy参数的值
	AutomaticallyCopy bool   `json:"automatically_copy,omitempty"` //自动复制推送内容ios >=14.5 长按或下拉推送可出发自动复制
}

type TextBarkRequest struct {
	Title     string `json:"title,omitempty"`      //标题
	Content   string `json:"content,omitempty"`    //内容
	IsArchive bool   `json:"is_archive,omitempty"` //自动保存通知消息
}

type RingBarkRequest struct {
}


const (
	Active        Level = iota + 1 //立即亮屏通知
	TimeSensitive                  //时效性通知,可在专注状态下显示通知
	Passive                        //仅将通知添加到通知列表,不会亮屏提醒
)

type Level int8

func (l Level) String() string {
	switch l {
	case TimeSensitive:
		return "timeSensitive"
	case Passive:
		return "passive"
	default:
		return "active"
	}
}


