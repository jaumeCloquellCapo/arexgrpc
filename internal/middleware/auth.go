package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jaumeCloquellCapo/authGrpc/app/model"
	"github.com/jaumeCloquellCapo/authGrpc/app/service"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type authMiddleware struct {
	//cache *storage.DbCache
	authService service.AuthServiceInterface
}

//AuthMiddlewareInterface ...
type AuthMiddlewareInterface interface {
	Handler() gin.HandlerFunc
}

//NewAuthMiddleware ...
func NewAuthMiddleware(authService service.AuthServiceInterface) AuthMiddlewareInterface {
	return &authMiddleware{
		authService,
	}
}

// Middleware functions take an gin.HandlerFunc and return a new one that can run code before and/or after calling the original handler — or it can decide not to call the original handler at all.
func (am authMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		am.ValidateAuth(c)
		c.Next()
	}
}

//ValidateAuth ...
func (am authMiddleware) ValidateAuth(c *gin.Context) {

	tokenAuth, err := ExtractTokenMetadata(c.Request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login first"})
		return
	}

	userID, err := am.authService.GetAuth(tokenAuth.AccessUUID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login first"})
		return
	}
	// Write user_id and user_model to the context
	c.Set("userID", userID)
}

//VerifyToken ...
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	return token, err
}

//ExtractToken extract token from Authorization header
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//ExtractTokenMetadata ...
func ExtractTokenMetadata(r *http.Request) (AccessDetails *model.AccessDetails, err error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {

		accessUUID, ok := claims["access_uuid"].(string)

		if !ok {
			return nil, err
		}

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)

		if err != nil {
			return nil, err
		}

		return &model.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}