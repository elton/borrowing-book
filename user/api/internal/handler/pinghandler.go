package handler

import (
	"net/http"

	"github.com/elton/borrowing-book/user/api/internal/logic"
	"github.com/elton/borrowing-book/user/api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func pingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewPingLogic(r.Context(), ctx)
		err := l.Ping()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
