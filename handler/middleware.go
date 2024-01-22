package handler

import (
	"github.com/samsul96maarif/auth-service/lib"
	"net/http"
)

func (handler *Handler) AuthMiddleware(next http.Handler, route_slag string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var skipPermission bool
		email := r.URL.Query().Get("email")
		ctx := r.Context()
		user, err := handler.BE.Usecase.FindUser(ctx, email)
		if err != nil || user == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userRole, err := handler.BE.Usecase.FindUserRole(ctx, int(user.Id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if userRole.RoleId == lib.ROLE_SUPER_ADMIN_ID {
			skipPermission = true
		}
		if !skipPermission {
			// ini bisa di cache

			// create-user
			// users/
			// post
			permission, err := handler.BE.Usecase.FindPermission(ctx, route_slag)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if permission == nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// ini bisa di cache
			_, err = handler.BE.Usecase.FindRolePermission(ctx, userRole.RoleId, permission.Slug)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
