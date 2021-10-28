module github.com/gwuhaolin/livego

go 1.13

require (
	github.com/auth0/go-jwt-middleware v0.0.0-20190805220309-36081240882b
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v7 v7.2.0
	github.com/kr/pretty v0.2.1
	github.com/lbryio/transcoder v0.14.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/negroni v1.0.0 // indirect
)

replace github.com/lbryio/transcoder => /home/randy/dev/src/github.com/OdyseeTeam/transcoder

replace github.com/floostack/transcoder => github.com/andybeletsky/transcoder v1.2.1
