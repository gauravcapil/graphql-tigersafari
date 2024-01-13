select * from (select distinct on (s.tiger_id) * from sightings s order by s.tiger_id, s.seen_at desc) join tiger_data t on t.id = tiger_id;
