package linkhub

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/xtaci/smux"
)

func NewHTTP(next Handler) http.Handler {
	return &httpHandler{
		next: next,
		wsup: &websocket.Upgrader{
			HandshakeTimeout: 10 * time.Second,
			CheckOrigin:      func(r *http.Request) bool { return true },
		},
	}
}

type httpHandler struct {
	next Handler
	wsup *websocket.Upgrader
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := h.wsup.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	conn := ws.NetConn()
	sess, err := smux.Server(conn, nil)
	if err != nil {
		return
	}

	h.next.Handle(sess)
}
