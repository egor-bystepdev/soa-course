
docker build -t dofe71/mafia_client -f client/Dockerfile .
docker build -t dofe71/mafia_server -f server/Dockerfile .
docker build -t dofe71/soa_graphql -f sessions_stat_service/Dockerfile .
