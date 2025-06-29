package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sikozonpc/social/internal/store"
)

type userKey string

const userCtx userKey = "user"

type FollowerUser struct {
	UserID int64 `json:"user_id"`
}

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {

	user := getUserFromContext(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromContext(r)

	// TODO: Revert back to auth userID from ctx
	var payload FollowerUser

	if err := readJSON(w, r, &payload); err != nil {
		app.statusBadRequest(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Follow.Follow(ctx, followerUser.ID, payload.UserID); err != nil {

		switch err {
		case store.ErrConflict:
			app.conflictResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}
func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowedUser := getUserFromContext(r)

	// TODO: Revert back to auth userID from ctx
	var payload FollowerUser

	if err := readJSON(w, r, &payload); err != nil {
		app.statusBadRequest(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Follow.UnFollow(ctx, unfollowedUser.ID, payload.UserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) useContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)

		if err != nil {
			app.statusBadRequest(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := app.store.User.GetById(ctx, userID)
		if err != nil {
			switch err {
			case store.ErrNotFound:
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
				return
			}
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}
