package shipx

import (
	"encoding/base64"
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/vela-ssoc/ssoc-common/data/problem"
	"github.com/vela-ssoc/ssoc-common/validation"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func NotFound(*ship.Context) error {
	return ship.ErrNotFound.Newf("资源不存在")
}

func HandleError(c *ship.Context, e error) {
	pd := &problem.Details{
		Type:     c.Host(),
		Title:    "请求错误",
		Status:   http.StatusBadRequest,
		Detail:   e.Error(),
		Instance: c.RequestURI(),
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
	case *mysql.MySQLError:
		switch err.Number {
		case 1062:
			pd.Detail = "数据已存在"
		default:
			c.Errorf("SQL 执行错误：%v", e)
			pd.Status = http.StatusInternalServerError
			pd.Detail = "内部错误"
		}
	default:
		switch {
		case err == gorm.ErrRecordNotFound:
			pd.Detail = "数据不存在"
		}
	}

	c.SetRespHeader(ship.HeaderContentType, "application/problem+json; charset=utf-8")
	c.SetRespHeader(ship.HeaderContentLanguage, "zh")

	_ = c.JSON(pd.Status, pd)
}
