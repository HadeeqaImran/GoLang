### Go Project for a Simple Hello RestAPI

#### Initialize Go Project
go mod init <module-name>

#### Install GoFiber using the following command
go get github.com/gofiber/fiber/v2
What is gofiber?

#### Install Fiber Swagger using the following
go get github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/fiber-swagger


What has to go in GOBIN and PATH env variables? and what they mean
he GOPATH environment variable represents the root of your Go workspace.
It should point to the directory where your Go projects, source code, and binaries are stored.
You can have multiple workspaces, but each workspace has its own GOPATH


For my system this command: "ls -l /Users/mac/go/bin/swag" gave this: "-rwxr-xr-x  1 mac  staff  14568960 Jan 23 23:14 /Users/mac/go/bin/swag"
So I had to set my environment variables as follows:  
export GOPATH=$HOME 
export GOBIN=/Users/mac/go/bin  
export PATH=$PATH:$GOBIN  

For windows you need to run: setx GOPATH "%USERPROFILE%\go"
setx PATH "%PATH%;%GOPATH%\bin"

<<<<<<< HEAD
and this enabled me to run my swag init command

#### Create a main.go
The code has sufficient comments

#### Run your app
go run main.go
