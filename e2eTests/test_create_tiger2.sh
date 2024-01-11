value=$(cat createTiger.json)

curl 'http://localhost:9000/query' \
-H 'Authorization: '$(./test_login.sh | jq ".data.login.token" | sed "s;\";;g") \
  -F "operations=$(cat createTiger2.json)" \
    -F map='{ "0": ["variables.photo"] }' \
      -F 0=@test_pic.jpg
