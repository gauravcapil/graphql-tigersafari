package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"gaurav.kapil/tigerhall/dbutils"
	"gaurav.kapil/tigerhall/graph/model"
)

func Authenticate(ctx context.Context) error {

	token := ctx.Value(TauthCtxKey)
	if token == nil {
		return fmt.Errorf("no token was passed")
	}
	result := []*model.LoginData{}
	dbutils.DbConn.Where(&model.LoginData{Token: token.(string)}).First(&result)
	if len(result) == 0 || token == "" {
		return fmt.Errorf("operation not allowed without a valid token")
	}
	return nil
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

type TigerToken struct {
	tauth string
}

var TauthCtxKey = &TigerToken{tauth: "tauthctx"}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			ctx := context.WithValue(r.Context(), TauthCtxKey, tokenStr)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
