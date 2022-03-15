package main

import (
	"errors"
	"example/web-service-gin/controllers"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("AUDIENCE", "api")
	os.Setenv("AUTHORITY", "https://demo.identityserver.io")
	os.Setenv("PORT", "8080")

	router := gin.Default()
	router.Use(EnsureValidToken())

	controllers.AddAlbumsController(router)

	router.Run()
}

// EnsureValidToken is a gin.HandlerFunc middleware that will check the validity of our JWT.
func EnsureValidToken() gin.HandlerFunc {
	issuerURL, err := url.Parse(os.Getenv("AUTHORITY"))
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 24*time.Hour)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUDIENCE")},
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	// errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
	// 	log.Printf("Encountered error while validating JWT: %v", err)
	// }

	// middleware := jwtmiddleware.New(
	// 	jwtValidator.ValidateToken,
	// 	jwtmiddleware.WithErrorHandler(errorHandler),
	// )

	return func(ctx *gin.Context) {
		// var encounteredError = false
		// var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// 	encounteredError = false
		// 	ctx.Request = r
		// 	ctx.Next()
		// }

		token, err := authHeaderTokenExtractor(ctx.Request)
		if err != nil {
			log.Println("Failed to extract bearer token:", err)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				map[string]string{"message": "Failed to validate JWT."},
			)
			return
		}

		validatedClaims, err := jwtValidator.ValidateToken(ctx.Request.Context(), token)
		if err != nil {

			log.Println("Failed to validate token:", err)

			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				map[string]string{"message": "Failed to validate JWT."},
			)
			return
		}

		log.Printf("claims: %+v", validatedClaims)

		// middleware.CheckJWT(handler).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func authHeaderTokenExtractor(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no JWT.
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}
