# Movie Server
Movie Server with [MongoDB]("https://www.mongodb.com/") and [Go]("https://golang.org") and [Gin]("https://gin-gonic.com/") and [Yaml]("https://yaml.org") for enviroment

## Steps
- Clone the repo `git clone https://github.com/en1tan/movie-server-go.git`
- Run Projet `go run ./main.go`

### Project Structure
```.
├── config
│   ├── config.go
│   ├── local.yaml
│   └── prod.yaml
├── controllers
│   └── movie-controller.go
├── database
│   └── database.go
├── go.mod
├── go.sum
├── main.go
├── models
│   └── movie-model.go
├── movie-server
├── README.md
├── routes
│   ├── router.go
│   └── routes.go
├── services
│   └── movie-service.go
└── utils
    └── errors.go
