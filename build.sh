mkdir -p "data/es"

docker-compose rm -sf
docker-compose up -d
echo -e "\n ======= build ==========  \n"
docker-compose ps

echo -e "\n ======== run ===========  \n"
go run main.go