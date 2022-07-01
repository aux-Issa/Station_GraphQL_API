# Station_GraphQL_API
## Data source
https://ekidata.jp/dl/?p=1
## verify completion of DB setup
`PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d postgres -c 'select * from station limit 10';`
## gqlgen: getting started 
`mkdir ./app`
`cd app`  
`go mod init github.com/aux-Issa/Station_GraphQL_API`  
`go get github.com/99designs/gqlgen`  
`go run github.com/99designs/gqlgen init`  
`go mod tidy`  

## generate Struct for Go from DB schema
install xo  
`go get -u github.com/xo/xo`  
`mkdir models`  
generate Struct  
``
