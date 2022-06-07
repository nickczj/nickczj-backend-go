Personal project for random stuff. See [routes.go](routes.go) for APIs.

HTTP web framework with gin-gonic, gorm, viper, redis.
Integrations with GCP app engine, auth0.
See [go.mod](go.mod) for everything.

See [here](https://github.com/nickczj/nickczj-backend-go/blob/b6460dd87aeefb0539277228e2a2d00a871d3e58/cache/redis.go#L40) for a generic cacheable. Only available with [go 1.18](https://go.dev/blog/intro-generics).
GCP app engine not supported yet (up to 1.16 as of 7 Jun 2022)

Environment variables:

```
APP_ENV
GCP_PROJECT_ID
GCP_BUCKET_NAME
```