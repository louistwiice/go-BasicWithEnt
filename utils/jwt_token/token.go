package jwttoken

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/louistwiice/go/basicwithent/configs"
)

// GenerateToken generates token when user connects himself
func GenerateToken(user_id string) (map[string]string, error) {
	accessclaims := make(jwt.MapClaims)  // access token claim
	refreshclaims := make(jwt.MapClaims)  //refresh token claim
	now_time := time.Now().UTC()

	// Generate access token
	accessclaims["authorized"] = true
	accessclaims["sub"] = user_id
	accessclaims["exp"] = now_time.Add(time.Hour * time.Duration(configs.GetInt("ACCESS_TOKEN_HOUR_LIFESPAN"))).Unix()
	accessclaims["iat"] = now_time.Unix()
	accessclaims["nbf"] = now_time.Unix()

	accesstoken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessclaims)
	at, err := accesstoken.SignedString([]byte(configs.GetString("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshclaims["authorized"] = true
	refreshclaims["sub"] = user_id
	refreshclaims["exp"] = now_time.Add(time.Hour * time.Duration(configs.GetInt("REFRESH_TOKEN_HOUR_LIFESPAN"))).Unix()
	refreshclaims["iat"] = now_time.Unix()
	refreshclaims["nbf"] = now_time.Unix()

	refreshtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessclaims)
	rt, err := refreshtoken.SignedString([]byte(configs.GetString("REFRESH_TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	return map[string]string{"access_token": at, "refresh_token": rt}, nil
}

// ExtractToken retrieves the token send by the user in every request
func ExtractToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	// If the token doesn not start with a good prefix: like Bearer, Token, ..., we return empty string
	if !strings.HasPrefix(bearerToken, configs.GetString("TOKEN_PREFIX")) {
		return ""
	}

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// IsTokenValid check if a token is valid and not expired
func IsTokenValid(ctx *gin.Context) error {
	tokenString := ExtractToken(ctx)

	token , err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.GetString("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return err
}

// ExtractTokenID extracts the user ID on a Token from Access token
func ExtractClaimsFromAccess(ctx *gin.Context) (jwt.MapClaims, error) {
	tokenString := ExtractToken(ctx)
	token , err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.GetString("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if err != nil {
			return nil, err
		}
		return claims, nil
	}

	return nil, nil
}

// ExtractTokenID extracts the user ID on a Token from refresh token
func ExtractClaimsFromRefresh(refresh_string string) (jwt.MapClaims, error) {
	token , err := jwt.Parse(refresh_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.GetString("REFRESH_TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if err != nil {
			return nil, err
		}
		return claims, nil
	}

	return nil, nil
}

// Check refresh token
func RefreshToken(ctx *gin.Context, refresh_string string) (map[string]string, error) {
	rtoken , err := jwt.Parse(refresh_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.GetString("REFRESH_TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if refreshClaim, ok := rtoken.Claims.(jwt.MapClaims); ok && rtoken.Valid {
		return GenerateToken(refreshClaim["sub"].(string))
	}
	return nil, err
}