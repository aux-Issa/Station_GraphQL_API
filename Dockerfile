FROM postgres:19.0
WORKDIR /Users/issa/Desktop/workspace/Station_GraphQL_API/
COPY init /docker-entrypoint-initdb.d/