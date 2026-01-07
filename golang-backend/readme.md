create a project folder and run the following commands

go mod init github.com/[git_user]/[proj_name]

go get github.com/githubnemo/CompileDaemon@latest

go install github.com/githubnemo/CompileDaemon@latest - Get Compile Daemon, used to cimple program after every change

go get github.com/joho/godotenv   

go get -u github.com/gin-gonic/gin  - Get Gin framework

go get -u gorm.io/gorm - Get GORM

go get -u gorm.io/driver/postgres - Get Postgres

CompileDaemon -command="./[proj_name].exe" - to run the .go file

Easiest postgreSQL would be Elephant SQL