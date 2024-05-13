package user

import (
	"net/http"
	"wire/domain"
)

type hanlder struct {
	svc domain.UserService
}

func (h *hanlder) FetchByUserName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := h.svc.FetchByUserName(r.Context(), "test")
		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(user.UserName))
			return
		}
		return
	}
}
