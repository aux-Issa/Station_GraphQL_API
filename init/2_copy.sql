
copy station(station_cd,station_g_cd,station_name,station_name_k,station_name_r,line_cd,pref_cd,post,address,lon,lat,open_ymd,close_ymd,e_status,e_sort)
  from '/docker-entrypoint-initdb.d/csv/station20220425free.csv' with csv header;