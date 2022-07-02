# Station_GraphQL_API
## Data source
https://ekidata.jp/dl/?p=1
## create DB refering to the sql_file
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
install xo  
`go get -u github.com/xo/xo`  
`mkdir models`  
`export PATH="$HOME/go/bin:$PATH"`  
generate Struct  

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
         left outer join station_join sjb on st.line_cd = sjb.line_cd and st.station_cd = sjb.station_cd2
         left outer join station_join sja on st.line_cd = sja.line_cd and st.station_cd = sja.station_cd1
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