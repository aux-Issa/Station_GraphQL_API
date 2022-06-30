# Station_GraphQL_API
## Data source
https://ekidata.jp/dl/?p=1
## DBのセットアップ完了確認
`PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d postgres -c 'select * from station limit 10';`
