package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/auth"
	"RSOI_CW/internal/pkg/middleware"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mailru/easyjson"
	"net/http"
	"os"
	"time"
)

type AuthHandler struct {
	repo auth.IRepo
}

func NewAuthHandler(repo auth.IRepo) *AuthHandler {
	return &AuthHandler{repo: repo}
}

func (h *AuthHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := easyjson.UnmarshalFromReader(r.Body, &user)
	if err != nil {
		middleware.Response(w, models.StatusBadUUID, nil)
		return
	}
	user, status := h.repo.AddUser(user)
	middleware.Response(w, status, user)
}

func (h *AuthHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, status := h.repo.GetUsers()
	middleware.Response(w, status, users)
}

func (h *AuthHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == http.MethodOptions {
		return
	}

	login, password, err := r.BasicAuth()
	if !err {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}
	user, state := h.repo.GetUser(login, password)
	if state != models.StatusOkey {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	tokenModel := models.Token{
		UserUUID: user.UUID,
		UserRole: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}

	SecretKey, err := os.LookupEnv("SECRET")
	if !err {
		panic("where is secret key")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenModel)

	jwtCookie, _ := token.SignedString([]byte(SecretKey))

	cookie := http.Cookie{Name: "Token", Value: jwtCookie}
	http.SetCookie(w, &cookie)
}

func (h *AuthHandler) CheckToken(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("Token")
	cookieValue := cookie.Value

	token, err := jwt.ParseWithClaims(cookieValue,
		&models.Token{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			SecretKey, _ := os.LookupEnv("SECRET")
			return []byte(SecretKey), nil
		})

	if err != nil {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	if ok := token.Valid; !ok {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	tk := token.Claims.(*models.Token)
	if tk.ExpiresAt >= time.Now().Unix() {
		middleware.Response(w, models.StatusOkey, nil)
	} else {
		middleware.Response(w, models.StatusNoAuth, nil)
	}
}

func (h *AuthHandler) CheckAdminToken(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("Token")
	cookieValue := cookie.Value

	token, err := jwt.ParseWithClaims(cookieValue,
		&models.Token{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			SecretKey, _ := os.LookupEnv("SECRET")
			return []byte(SecretKey), nil
		})

	if err != nil {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	if ok := token.Valid; !ok {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	tk := token.Claims.(*models.Token)
	if tk.ExpiresAt >= time.Now().Unix() {
		middleware.Response(w, models.StatusOkey, nil)
	} else {
		middleware.Response(w, models.StatusNoAuth, nil)
	}

}
