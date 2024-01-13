source ../config.cfg
psql -h localhost -U $PGUSER -d postgres -c "select tiger_id,seen_at,seen_at_lat,seen_at_lon, photo_location, user_name, name, date_of_birth from (select distinct on (s.tiger_id) * from sightings s order by s.tiger_id, s.seen_at desc) join tiger_data t on t.id = tiger_id limit $1 offset $2;"
