package shipx

import (
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/vela-ssoc/ssoc-common/validation"
	"github.com/vela-ssoc/ssoc-proto/muxtool"
	"github.com/xgfone/ship/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ErrorHandler struct {
	log *slog.Logger
}

func NewErrorHandler(log *slog.Logger) *ErrorHandler {
	return &ErrorHandler{
		log: log,
	}
}

func (*ErrorHandler) NotFound(*ship.Context) error { return ship.ErrNotFound }

func (h *ErrorHandler) HandleError(c *ship.Context, err error) {
	r := c.Request()
	pd := &muxtool.ProblemDetails{
		Status:   0,
		Title:    "",
		Detail:   "",
		Instance: r.URL.Path,
		Method:   r.Method,
		Host:     r.Host,
	}
	switch ev := err.(type) {
	case ship.HTTPServerError:
		pd.Status = ev.Code
	case *ship.HTTPServerError:
		pd.Status = ev.Code
	case *validation.ValidError:
		pd.Title = "参数校验错误"
	case *time.ParseError:
		pd.Title = "参数格式错误"
		pd.Detail = "时间格式错误，正确格式：" + ev.Layout
	case *net.ParseError:
		pd.Title = "参数格式错误"
		pd.Detail = ev.Text + " 不是有效的 " + ev.Type
	case base64.CorruptInputError:
		pd.Title = "参数格式错误"
		pd.Detail = "base64 编码错误：" + err.Error()
	case *json.SyntaxError:
		pd.Title = "格式错误"
		pd.Detail = "不合法的 JSON 数据"
	case *json.UnmarshalTypeError:
		pd.Title = "类型错误"
		pd.Detail = ev.Field + " 收到无效的数据类型"
	case *strconv.NumError:
		pd.Title = "数据类型错误"
		var msg string
		if sn := strings.SplitN(ev.Func, "Parse", 2); len(sn) == 2 {
			msg = ev.Num + " 不是 " + strings.ToLower(sn[1]) + " 类型"
		} else {
			msg = "类型错误：" + ev.Num
		}
		pd.Detail = msg
	case mongo.WriteException:
		var over bool
		for _, we := range ev.WriteErrors {
			if over {
				break
			}
			switch we.Code {
			case 11000:
				pd.Status = http.StatusConflict
				pd.Title = "数据已存在"
				pd.Detail = "数据已存在"
				kv := we.Raw.Lookup("keyValue").String()
				if kv != "" {
					pd.Detail += "：" + kv
				}
				over = true
			}
		}
	}

	if pd.Status < 100 || pd.Status > 999 {
		pd.Status = http.StatusBadRequest
	}
	if pd.Title == "" {
		pd.Title = http.StatusText(pd.Status)
	}
	if pd.Detail == "" {
		pd.Detail = err.Error()
	}

	h.log.Warn("统一错误处理", "detail", pd, "error", err)

	_ = c.JSON(pd.Status, pd)
}
