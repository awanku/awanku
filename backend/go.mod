module github.com/awanku/awanku/backend

go 1.14

replace github.com/asasmoyo/pq-hansip => ./lib/pq-hansip

require (
	cloud.google.com/go v0.57.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200428143746-21a406dcc535 // indirect
	github.com/asasmoyo/pq-hansip v0.0.0-20190502052219-d515e288ee85
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-ozzo/ozzo-validation/v4 v4.2.1
	github.com/go-pg/pg/v10 v10.0.0-beta.1
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/google/go-github v17.0.0+incompatible
	github.com/gorilla/securecookie v1.1.1
	github.com/segmentio/encoding v0.1.13 // indirect
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
	google.golang.org/api v0.26.0
	google.golang.org/genproto v0.0.0-20200605102947-12044bf5ea91 // indirect
)
