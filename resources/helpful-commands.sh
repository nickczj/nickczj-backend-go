# Handy docker stuff
docker container prune
docker volume prune
docker stop container_name

# Show Nginx config
docker exec nick_nginx-proxy_1 cat /etc/nginx/conf.d/default.conf

# ARCHIVE
docker exec hosting_plausible_db_1 psql -U postgres -d plausible_db -c "show tables;"
docker run -it --net hosting_default --rm --link plausible_events_db_1:clickhouse-server yandex/clickhouse-client --host plausible_plausible_events_db_1  --query "CREATE DATABASE IF NOT EXISTS plausible_events_db"