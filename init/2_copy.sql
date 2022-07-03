
copy station(station_cd,station_g_cd,station_name,station_name_k,station_name_r,line_cd,pref_cd,post,address,lon,lat,open_ymd,close_ymd,e_status,e_sort)
  from '/docker-entrypoint-initdb.d/csv/station20220425free.csv' with csv header;

copy company(company_cd,rr_cd,company_name,company_name_k,company_name_h,company_name_r,company_url,company_type,e_status,e_sort)
  from '/docker-entrypoint-initdb.d/csv/company20220401.csv' with csv header;

copy junction(line_cd,station_cd1,station_cd2)
  from '/docker-entrypoint-initdb.d/csv/join20220413.csv' with csv header;

copy line(line_cd,company_cd,line_name,line_name_k,line_name_h,line_color_c,line_color_t,line_type,lon,lat,zoom,e_status,e_sort)
  from '/docker-entrypoint-initdb.d/csv/line20220323free.csv' with csv header;