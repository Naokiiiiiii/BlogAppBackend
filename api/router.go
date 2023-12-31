package api

import (
	"database/sql"
	"net/http"

	"github.com/Naokiiiiiii/BlogApiPractice/api/middlewares"
	"github.com/Naokiiiiiii/BlogApiPractice/controllers"
	"github.com/Naokiiiiiii/BlogApiPractice/services"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

func NewRouter(db *sql.DB, config oauth2.Config) *mux.Router {
	ser := services.NewMyAppService(db, config)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)
	nCon := controllers.NewNiceController(ser)
	uCon := controllers.NewUserController(ser, config)
	r := mux.NewRouter()

	r.Use(middlewares.CorsMiddleware)
	r.Use(middlewares.LoggingMiddleware)

	// 認証が必要ないAPI
	// ログインAPI
	r.HandleFunc("/login", uCon.GoogleLoginHandler)
	r.HandleFunc("/token", uCon.GoogleTokenHandler)
	r.HandleFunc("/regenerate_token", uCon.RegenerateAccessTokenHandler).Methods(http.MethodPost, http.MethodOptions)

	// 認証が必要なAPI
	authRequired := r.PathPrefix("").Subrouter()

	authRequired.Use(middlewares.AuthMiddleware)

	// ユーザーAPI
	authRequired.HandleFunc("/user", uCon.SelectUserInfoHandler)
	authRequired.HandleFunc("/user/{id:[0-9]+}", uCon.UpdateUserHandler).Methods(http.MethodPut, http.MethodOptions)

	// 記事API
	authRequired.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost, http.MethodOptions)
	authRequired.HandleFunc("/article/{id:[0-9]+}", aCon.UpdateArticleHandler).Methods(http.MethodPut, http.MethodOptions)
	authRequired.HandleFunc("/article/{id:[0-9]+}", aCon.DeleteArticleHandler).Methods(http.MethodDelete, http.MethodOptions)
	authRequired.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet, http.MethodOptions)
	authRequired.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet, http.MethodOptions)

	// いいねAPI
	authRequired.HandleFunc("/article/nice", nCon.CreateOrDeleteNiceHandler).Methods(http.MethodPost, http.MethodOptions)

	// コメントAPI
	authRequired.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost, http.MethodOptions)
	authRequired.HandleFunc("/comment/{id:[0-9]+}", cCon.UpdateCommentHandler).Methods(http.MethodPut, http.MethodOptions)
	authRequired.HandleFunc("/comment/{id:[0-9]+}", cCon.DeleteCommentHandler).Methods(http.MethodDelete, http.MethodOptions)

	return r
}
