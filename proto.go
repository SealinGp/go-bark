package gobark

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type BarkRequest struct {
	BarkRequestOptions
	Text *Text `json:"text,omitempty"` //普通消息
}

func (r *BarkRequest) set(httpReq *http.Request) error {
	r.setOpts(httpReq)

	if r.Text != nil {
		r.Text.set(httpReq)
		return nil
	}

	return errors.New("text or ring required")
}

func (r *BarkRequest) setOpts(httpReq *http.Request) {
	q := httpReq.URL.Query()
	if r.Icon != "" {
		q.Set("icon", r.Icon)
	}

	if r.Group != "" {
		q.Set("group", r.Group)
	}

	if r.Level > 0 {
		q.Set("level", r.Level.String())
	}

	if r.Url != "" {
		q.Set("url", r.Url)
	}

	if r.Badge > 0 {
		q.Set("badge", fmt.Sprintf("%v", r.Badge))
	}

	if r.Copy != "" {
		q.Set("copy", r.Copy)
	}

	if r.AutomaticallyCopy {
		q.Set("authCopy", "1")
	}

	if r.Sound != "" {
		q.Set("sound", r.Sound)
	}

	httpReq.URL.RawQuery = q.Encode()
}

func (r *Text) set(httpReq *http.Request) {
	u := httpReq.URL
	subjects := make([]string, 0, 2)
	if r.Title != "" {
		subjects = append(subjects, r.Title)
	}

	if r.Content != "" {
		subjects = append(subjects, r.Content)
	}

	if r.IsArchive {
		u.Query().Set("isArchive", "1")
	}

	u.Path += "/" + strings.Join(subjects, "/")
}

type BarkRequestOptions struct {
	Sound             string `json:"sound,omitempty"`              //可在Bark App查看所有的铃声名
	Icon              string `json:"icon,omitempty"`               //自定义推送图标,图标将替换默认Bark图标,会自动缓存在客户端,https://xxx/xx.jpg
	Group             string `json:"group,omitempty"`              //消息分组名,同一个分组下的消息将会叠加显示
	Level             Level  `json:"level,omitempty"`              //消息时效性配置
	Url               string `json:"url,omitempty"`                //点击通知可跳转的url
	Badge             int    `json:"badge,omitempty"`              //通知角标
	Copy              string `json:"copy,omitempty"`               //下啦推送,锁屏界面左滑查看推送时,可选择复制推送内容,携带该参数时,将只复制copy参数的值
	AutomaticallyCopy bool   `json:"automatically_copy,omitempty"` //自动复制推送内容ios >=14.5 长按或下拉推送可出发自动复制
}

type Text struct {
	Title     string `json:"title,omitempty"`      //标题
	Content   string `json:"content,omitempty"`    //内容
	IsArchive bool   `json:"is_archive,omitempty"` //自动保存通知消息
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

type CommonResp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	Timestamp int64  `json:"timestamp"`
}

func (r *CommonResp) Error() error {
	if r.Code != 200 {
		return fmt.Errorf("code %v should be 200", r.Code)
	}

	return nil
}

type Responser interface {
	Error() error
}
