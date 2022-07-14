# Station_GraphQL_API
## Data source
https://ekidata.jp/dl/?p=1
## create DB referring to the sql_file
`docker volume create pg-data-station`  
`docker-compose up --build`
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
### install [xo](https://github.com/xo/xo)
> xo can automatically genarate struct of Golang referring to input query  

`go get -u github.com/xo/xo`  
`mkdir models`  
`export PATH="$HOME/go/bin:$PATH"`  
### generate Struct  

`xo query postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -M -B -T StationConn -o models/ << ENDSQL   
select li.line_name,
       li.line_name_h,
       li.line_cd,
       st.station_cd,
       st.station_g_cd,
       st.address,
       st.station_name,
       COALESCE(s2l.line_name, '')     as before_line_name,
       COALESCE(st2.station_cd, 0)    as before_station_cd,
       COALESCE(st2.station_name, '') as before_station_name,
       COALESCE(st2.address, '')      as before_address,
       COALESCE(s3l.line_name, '')     as after_line_name,
       COALESCE(st3.station_cd, 0)    as after_station_cd,
       COALESCE(st3.station_name, '') as after_station_name,
       COALESCE(st2.address, '')      as after_address,
       COALESCE(gli.line_name, '')    as transfer_line_name,
       COALESCE(gs.station_cd, 0)     as transfer_station_cd,
       COALESCE(gs.station_name, '')  as transfer_station_name,
       COALESCE(gs.address, '')       as transfer_address
from station st
         inner join line li on st.line_cd = li.line_cd
         left outer join junction sjb on st.line_cd = sjb.line_cd and st.station_cd = sjb.station_cd2
         left outer join junction sja on st.line_cd = sja.line_cd and st.station_cd = sja.station_cd1
         left outer join station st2 on sjb.line_cd = st2.line_cd and sjb.station_cd1 = st2.station_cd
         left outer join line s2l on st2.line_cd = s2l.line_cd
         left outer join station st3 on sja.line_cd = st3.line_cd and sja.station_cd2 = st3.station_cd
         left outer join line s3l on st3.line_cd = s3l.line_cd
         left outer join station gs on st.station_g_cd = gs.station_g_cd and st.station_cd <> gs.station_cd
         left outer join line gli on gs.line_cd = gli.line_cd
where st.station_cd = %%stationCD int%%
  and st.e_status = 0
order by st.e_sort
ENDSQL  `
### prevent overfetch by spliting resolver  
####　cdで駅を検索するクエリ
`xo query postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -M -B -T StationByCD -o models/ << ENDSQL
select l.line_cd, l.line_name, s.station_cd, s.station_g_cd, s.station_name, s.address
from station s
         inner join line l on s.line_cd = l.line_cd
where s.station_cd = %%stationCD int%%
  and s.e_status = 0
ENDSQL
`
#### 駅名で駅を検索する
`
xo query postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -M -B -T StationByName -o models/ << ENDSQL 
select l.line_cd, l.line_name, s.station_cd, s.station_g_cd, s.station_name, s.address 
from station s 
inner join line l on s.line_cd = l.line_cd 
where s.station_name = %%stationName string%% 
and s.e_status = 0 
ENDSQL
`
#### 乗り換え駅

`xo query postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -M -B -T Transfer -o models/ << ENDSQL
select s.station_cd,
       s.station_name,
       s.station_g_cd,
       s.address,
       ls.line_cd,
       ls.line_name,
       COALESCE(lt.line_cd, 0)     as transfer_line_cd,
       COALESCE(lt.line_name, '')   as transfer_line_name,
       COALESCE(t.station_cd, 0)   as transfer_station_cd,
       COALESCE(t.station_name, '') as transfer_station_name,
       COALESCE(t.address, '')      as transfer_address
from station s
        left outer join station t on s.station_g_cd = t.station_g_cd and s.station_cd <> t.station_cd
        left outer join line ls on s.line_cd = ls.line_cd
        left outer join line lt on t.line_cd = lt.line_cd
where s.station_cd = %%stationCD int%%
ENDSQL
`