package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strings"
)

type JWT struct {
	jwt.RegisteredClaims
	ID string
	XH string
}

func AUTH() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("Authorization")
				if tokenString == "" {
					return nil, errors.New("token is nil")
				}

				auths := strings.SplitN(tokenString, " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], "passport") {
					return nil, errors.New("jwt token missing")
				}

				token, err := jwt.ParseWithClaims(auths[1], &JWT{}, func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(os.Getenv("JWT_KEY")), nil
				})

				if err != nil {
					return nil, err
				}

				if claims, ok := token.Claims.(*JWT); ok && token.Valid {
					// put CurrentUser into ctx
					ctx = WithContext(ctx, claims.XH)
				} else {
					return nil, errors.New("token Invalid")
				}
			}
			return handler(ctx, req)
		}
	}
}

func FromContext(ctx context.Context) string {
	return ctx.Value("xh").(string)
}

func WithContext(ctx context.Context, xh string) context.Context {
	return context.WithValue(ctx, "xh", xh)
}
