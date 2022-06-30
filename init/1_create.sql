create table station
(
    station_cd integer not null,
    station_g_cd integer not null,
    station_name varchar not null,
    station_name_k varchar,
    station_name_r varchar,
    line_cd integer,
    pref_cd integer,
    post varchar,
    address varchar,
    lon float,
    lat float,
    open_ymd varchar ,
    close_ymd varchar,
    e_status integer,
    e_sort integer,
    PRIMARY KEY (station_cd)
);
comment on table station is 'station20220425free.csv';
