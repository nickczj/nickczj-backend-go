package auth

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/url"
	"time"
)

func Jwt() gin.HandlerFunc {
	issuerURL, _ := url.Parse(viper.GetString("auth0.issuer-url"))
	audience := viper.GetString("auth0.audience")

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
	)

	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
		return nil
	}

	jwtMiddleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	return adapter.Wrap(jwtMiddleware.CheckJWT)
}
