package middleware

import (
	"context"
	"fmt"
	"github.com/jaumeCloquellCapo/authGrpc/app/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"os"
	"strconv"
)

// Interceptor data structure
type Interceptor struct {
	authKey string
}

// NewInterceptor function, for init Interceptor object
func NewInterceptor(authKey string) *Interceptor {
	return &Interceptor{authKey}
}

//Auth function,
//or Unary interceptor
//additional security for our GRPC server
func (i *Interceptor) Auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "missing context metadata")
	}

	if len(meta["authorization"]) != 1 {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid authorization")
	}

	authorization := meta["authorization"][0]

	if authorization != i.authKey {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid authorization")
	}

	return handler(ctx, req)
}

// AuthStream

func (i *Interceptor) AuthStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	meta, ok := metadata.FromIncomingContext(stream.Context())

	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "missing context metadata")
	}

	if len(meta["authorization"]) != 1 {
		return grpc.Errorf(codes.Unauthenticated, "invalid authorization")
	}

	authorization := meta["authorization"][0]

	if authorization != i.authKey {
		return grpc.Errorf(codes.Unauthenticated, "invalid authorization")
	}

	return handler(srv, stream)
}

//VerifyToken ...
func VerifyTokenFromContext(ctx context.Context) (*jwt.Token, error) {
	tokenString := extractTokenFromContext(ctx)

	if tokenString == nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid authorization")
	}

	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, grpc.Errorf(codes.Unauthenticated, "invalid authorization")
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	return token, err
}

//ExtractToken extract token from Authorization header
func extractTokenFromContext(ctx context.Context) *string {

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	if len(meta["authorization"]) != 1 {
		return nil
	}

	return &meta["authorization"][0]
}

/**
ExtractTokenContext
*/
func GetAccessTokenFromContext(ctx context.Context) (AccessDetails *model.AccessDetails, err error) {
	token, err := VerifyTokenFromContext(ctx)
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
