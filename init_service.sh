# using sh in windows with the help of moba xterm

#initializing mod file
go mod init gaurav.kapil/graphql-tigersafari

# getting dependencies for project bootstrapping
go get github.com/99designs/gqlgen

# gqlgen is to be added in tool.go
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go

# updating dependencies
go mod tidy

# initialize gqlgen config and generate the models.
go run github.com/99designs/gqlgen init
