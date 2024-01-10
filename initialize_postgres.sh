source config.cfg
docker pull postgres
docker run -itd -e POSTGRES_USER=$PGUSER -e POSTGRES_PASSWORD=$PGPASS -p 5432:5432 -v /data:/var/lib/postgresql/data --name postgresql postgres
apt install odbc-psql postgresql-client

