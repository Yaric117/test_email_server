package routes

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"testproject/config"
	v1 "testproject/controller/http/v1"
	"testproject/docs"
	"testproject/middleware"
)

var Router = mux.NewRouter()

func init() {
	// middlewares
	Router.Use(middleware.AddCors)
	Router.Use(middleware.AppTokenAuthentication)

	docs.SwaggerInfo.Title = config.App.Name
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.App.Url
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	//api
	{
		apiGroup := Router.PathPrefix("/api/v1").Subrouter()
		apiGroup.HandleFunc("/test", v1.HandleTest).Methods(http.MethodGet, http.MethodOptions)
		apiGroup.HandleFunc("/create", v1.HandleCreateData).Methods(http.MethodPost, http.MethodOptions)
		apiGroup.HandleFunc("/windows", v1.HandleGetAllWindows).Methods(http.MethodGet, http.MethodOptions)
	}

	//api-docs
	Router.Handle("/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}
