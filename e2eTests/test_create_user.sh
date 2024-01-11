value=$(cat createTiger.json)

curl 'http://localhost:9000/query' -H 'content-type: application/json' \
      -d @createUser.json
