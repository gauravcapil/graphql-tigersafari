value=$(cat createTiger.json)
set -x
curl 'http://localhost:9000/query' -H 'Authorization: '$(./test_login.sh | jq ".data.login.token" | sed "s;\";;g") -H 'content-type: application/json' \
      -d @createUser.json
