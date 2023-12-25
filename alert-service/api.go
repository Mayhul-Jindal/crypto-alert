package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
)

type API struct {
	listenAddr string
	token      Maker
	auth       Auther
	validator  *validator.Validate
}

func NewAPI(listenAddr string, token Maker, auth Auther, validator *validator.Validate) *API {
	return &API{
		listenAddr: listenAddr,
		token:      token,
		auth:       auth,
		validator:  validator,
	}
}

func (a *API) Run(ctx context.Context) *http.Server {
	mux := chi.NewRouter()

	// public routes
	mux.Group(func(mux chi.Router) {
		mux.Get("/", a.handle(a.root))
		mux.Post("/signup", a.handle(a.signUp))
		mux.Get("/login", a.handle(a.login))
	})

	// private routes
	mux.Route("/alerts", func(mux chi.Router) {
		mux.Post("/", a.handle(a.authMiddleware(a.createAlert)))
	})

	server := &http.Server{
		Addr:    a.listenAddr,
		Handler: mux,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	return server
}

// Root handler
func (a *API) root(w http.ResponseWriter, r *http.Request) error {
	resp := map[string]string{"message": "ok"}
	return writeJSON(r.Context(), w, http.StatusOK, resp)
}

// Sign Up handler
func (a *API) signUp(w http.ResponseWriter, r *http.Request) error {
	var req SignUpUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return ErrBadRequest
	}

	err = a.validator.Struct(req)
	if err != nil {
		return ErrBadRequest
	}

	resp, err := a.auth.SignUp(r.Context(), req)
	if err != nil {
		return err
	}

	return writeJSON(r.Context(), w, http.StatusOK, resp)
}

// Login handler
func (a *API) login(w http.ResponseWriter, r *http.Request) error {
	var req LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return ErrBadRequest
	}

	err = a.validator.Struct(req)
	if err != nil {
		return ErrBadRequest
	}

	resp, err := a.auth.Login(r.Context(), req)
	if err != nil {
		return err
	}

	return writeJSON(r.Context(), w, http.StatusOK, resp)
}

// Create Alert handler
func (a *API) createAlert(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Create Alert"))
	return nil
}

// // Read Alert handler
// func handleRead(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Read Alert"))
// }

// // Update Alert handler
// func handleUpdate(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Update Alert"))
// }

// // Delete Alert handler
// func handleDelete(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Delete Alert"))
// }

// centralize error handling
type Handler func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func (a *API) handle(next Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			switch err {
			case ErrBadRequest, ErrNoAuthHeader, ErrInvalidAuthHeader, ErrUnsupportedAuthType:
				writeJSON(r.Context(), w, http.StatusBadRequest, ApiError{Error: err.Error()})

			case ErrNotAuthorized:
				writeJSON(r.Context(), w, http.StatusUnauthorized, ApiError{Error: err.Error()})

			default:
				log.Println("critical internal server error:", err)
				writeJSON(r.Context(), w, http.StatusInternalServerError, ApiError{Error: "internal server error"})
			}
		}
	}
}

// middlewares
func (a *API) authMiddleware(next Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		authorizationHeader := r.Header.Get("authorization")

		if len(authorizationHeader) == 0 {
			return ErrNoAuthHeader
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			log.Println("3")
			return ErrInvalidAuthHeader
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			log.Println("4")
			return ErrUnsupportedAuthType
		}

		accessToken := fields[1]
		payload, err := a.token.Verify(accessToken)
		if err != nil {
			return err
		}

		r = r.WithContext(context.WithValue(r.Context(), tokenPayload, payload))
		return next(w, r)
	}
}

// helper function
func writeJSON(ctx context.Context, w http.ResponseWriter, s int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)

	// centralized logging
	if apiErr, ok := v.(ApiError); ok {
		log.Println("api error", apiErr.Error)
	} else {
		log.Println("response", v)
	}

	return json.NewEncoder(w).Encode(v)
}
