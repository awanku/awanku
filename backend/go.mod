module github.com/awanku/awanku

go 1.14

replace github.com/asasmoyo/pq-hansip => ./lib/pq-hansip

require (
	cloud.google.com/go v0.61.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200428143746-21a406dcc535 // indirect
	github.com/asasmoyo/pq-hansip v0.0.0-20190502052219-d515e288ee85
	github.com/bxcodec/faker/v3 v3.5.0
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/cors v1.1.1
	github.com/go-openapi/jsonreference v0.19.4 // indirect
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-ozzo/ozzo-validation/v4 v4.2.1
	github.com/go-pg/pg/v9 v9.1.6
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/google/go-github v17.0.0+incompatible
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/swaggo/swag v1.6.7
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/tools v0.0.0-20200717024301-6ddee64345a6 // indirect
	google.golang.org/api v0.29.0
	google.golang.org/genproto v0.0.0-20200715011427-11fb19a81f2c // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)
