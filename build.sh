cd cli && go build -o agenda
cd ../services && go build -o agenda-server && rm -f test.db
