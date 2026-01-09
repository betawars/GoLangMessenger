create a project folder and run the following commands

go mod init github.com/[git_user]/[proj_name]

go get github.com/githubnemo/CompileDaemon@latest

go install github.com/githubnemo/CompileDaemon@latest - Get Compile Daemon, used to cimple program after every change

go get github.com/joho/godotenv   

go get -u github.com/gin-gonic/gin  - Get Gin framework

go get -u gorm.io/gorm - Get GORM

go get -u gorm.io/driver/postgres - Get Postgres

CompileDaemon -command="./[proj_name].exe" - to run the .go file

Easiest postgreSQL would be Elephant SQL(as Elephant SQL is no longer UP, I am using Supabase)

----Implementation of JWT authentication---

Additional packages required:

Go cryptography, can be found at: pkg.go.dev/golang.org/x/crypto
or
go get -u golang.org/x/crypto/bcrypt

go get -u github.com/golang-jwt/jwt/v4
