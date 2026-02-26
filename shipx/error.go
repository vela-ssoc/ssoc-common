package shipx

import (
	"encoding/base64"
	"encoding/json"
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

func NotFound(_ *ship.Context) error {
	return ship.ErrNotFound.Newf("资源不存在")
}

func HandleError(c *ship.Context, e error) {
	pd := &muxtool.ProblemDetails{
		Host:     c.Host(),
		Method:   c.Method(),
		Instance: c.RequestURI(),
		Status:   http.StatusBadRequest,
		Title:    "请求错误",
		Detail:   e.Error(),
	}

	switch err := e.(type) {
	case ship.HTTPServerError:
		pd.Status = err.Code
	case *ship.HTTPServerError:
		pd.Status = err.Code
	case *validation.ValidError:
		pd.Title = "参数校验错误"
	case *time.ParseError:
		pd.Title = "参数格式错误"
		pd.Detail = "时间格式错误，正确格式：" + err.Layout
	case *net.ParseError:
		pd.Title = "参数格式错误"
		pd.Detail = err.Text + " 不是有效的 " + err.Type
	case base64.CorruptInputError:
		pd.Title = "参数格式错误"
		pd.Detail = "base64 编码错误：" + err.Error()
	case *json.SyntaxError:
		pd.Title = "报文格式错误"
		pd.Detail = "请求报错必须是 JSON 格式"
	case *json.UnmarshalTypeError:
		pd.Title = "数据类型错误"
		pd.Detail = err.Field + " 收到无效的数据类型"
	case *strconv.NumError:
		pd.Title = "数据类型错误"
		var msg string
		if sn := strings.SplitN(err.Func, "Parse", 2); len(sn) == 2 {
			msg = err.Num + " 不是 " + strings.ToLower(sn[1]) + " 类型"
		} else {
			msg = "类型错误：" + err.Num
		}
		pd.Detail = msg
	case mongo.WriteException:
		for _, we := range err.WriteErrors {
			switch we.Code {
			case 11000:
				pd.Status = http.StatusConflict
				pd.Title = "数据已存在"
				pd.Detail = "数据已存在"
				kv := we.Raw.Lookup("keyValue").String()
				if kv != "" {
					pd.Detail += "：" + kv
				}
			}
		}
	}

	_ = c.JSON(pd.Status, pd)
}
