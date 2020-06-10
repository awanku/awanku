module github.com/awanku/awanku

go 1.14

replace github.com/asasmoyo/pq-hansip => ./lib/pq-hansip

require (
	cloud.google.com/go v0.58.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200428143746-21a406dcc535 // indirect
	github.com/asasmoyo/pq-hansip v0.0.0-20190502052219-d515e288ee85
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/cors v1.1.1
	github.com/go-ozzo/ozzo-validation/v4 v4.2.1
	github.com/go-pg/pg/v10 v10.0.0-beta.2
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/google/go-github v17.0.0+incompatible
	github.com/gorilla/schema v1.1.0
	github.com/gorilla/securecookie v1.1.1
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/api v0.26.0
)
