package users

import (
	"github.com/dgrijalva/jwt-go"
	jwtRequest "github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/qianyan/go-gin-quickstart/infra"
	"net/http"
	"strings"
)

// Strips 'TOKEN ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &jwtRequest.PostExtractionFilter{
	jwtRequest.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var AccessTokenExtractor = &jwtRequest.MultiExtractor{
	AuthorizationHeaderExtractor,
	jwtRequest.ArgumentExtractor{"access_token"},
}

// A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, currentUserId uint) {
	var currentUserModel UserModel
	if currentUserId != 0 {
		db := infra.GetDB()
		db.First(&currentUserModel, currentUserId)
	}
	c.Set("currentUserId", currentUserId)
	c.Set("currentUserModel", currentUserModel)
}

// You can custom middlewares yourself as the doc: https://github.com/gin-gonic/gin#custom-middleware
//  r.Use(AuthMiddleware(true))
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUserModel(c, 0)
		token, err := jwtRequest.ParseFromRequest(c.Request, AccessTokenExtractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(infra.NBSecretPassword))
			return b, nil
		})

		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			currentUserId := uint(claims["id"].(float64))
			UpdateContextUserModel(c, currentUserId)
		}
	}
}
