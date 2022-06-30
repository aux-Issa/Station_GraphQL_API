FROM postgres:10.7
WORKDIR /Users/issa/Desktop/workspace/Station_GraphQL_API/
COPY init /docker-entrypoint-initdb.d/