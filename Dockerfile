FROM postgres:15.20
WORKDIR /Users/issa/Desktop/workspace/Station_GraphQL_API/
COPY init /docker-entrypoint-initdb.d/