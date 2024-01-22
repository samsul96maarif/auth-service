package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	core "github.com/samsul96maarif/auth-service"
	"github.com/samsul96maarif/auth-service/config"
	"github.com/samsul96maarif/auth-service/handler"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello home")
	return
}

func Init() {
	godotenv.Load()
}

func main() {
	Init()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	var (
		srvErr = make(chan error, 1)
		err    error
	)
	db, _ := config.NewDB()
	be := core.NewBe(db)
	be.Init()
	handler := handler.NewHandler(&be)
	router := mux.NewRouter()

	router.HandleFunc("/seed", func(w http.ResponseWriter, r *http.Request) {
		err = handler.BE.Usecase.CreateSuperAdmin(r.Context())
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		return
	})

	router.HandleFunc("/revoke-perm", func(w http.ResponseWriter, r *http.Request) {
		role := r.URL.Query().Get("role_id")
		slug := r.URL.Query().Get("slug")
		roleId, erro := strconv.Atoi(role)
		if erro != nil {
			fmt.Fprintln(w, erro.Error())
			return
		}
		if erro = handler.BE.Usecase.RevokePermission(r.Context(), uint(roleId), slug); erro != nil {
			fmt.Fprintln(w, erro.Error())
			return
		}
		fmt.Fprintln(w, "succeed")
		return
	}).Methods(http.MethodDelete)

	router.HandleFunc("/add-perm", func(w http.ResponseWriter, r *http.Request) {
		role := r.URL.Query().Get("role_id")
		slug := r.URL.Query().Get("slug")
		roleId, erro := strconv.Atoi(role)
		if erro != nil {
			fmt.Fprintln(w, erro.Error())
			return
		}
		if erro = handler.BE.Usecase.AdddPermission(r.Context(), uint(roleId), slug); erro != nil {
			fmt.Fprintln(w, erro.Error())
			return
		}
		fmt.Fprintln(w, "succeed")
		return
	}).Methods(http.MethodPost)

	router.HandleFunc("/hello", handler.AuthMiddleware(http.HandlerFunc(Home), "hello-v2").ServeHTTP)
	router.HandleFunc("/hello", handler.AuthMiddleware(http.HandlerFunc(Home), "add-hello").ServeHTTP).Methods(http.MethodPost)
	router.HandleFunc("/hello", handler.AuthMiddleware(http.HandlerFunc(Home), "edit-hello").ServeHTTP).Methods(http.MethodPut)
	router.HandleFunc("/hello", handler.AuthMiddleware(http.HandlerFunc(Home), "delete-hello").ServeHTTP).Methods(http.MethodDelete)

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "pong")
		return
	})

	headerOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Authorization"})
	methodOk := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT"})
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(headerOk, methodOk)(router),
	}
	go func() {
		srvErr <- srv.ListenAndServe()
	}()
	select {
	case err = <-srvErr:
		fmt.Println("error appear ", err.Error())
		return
	case <-ctx.Done():
		fmt.Println("shutdown gracefully ", ctx.Err(), " at ", time.Now())
		stop()
	}
	err = srv.Shutdown(ctx)
	return
}
