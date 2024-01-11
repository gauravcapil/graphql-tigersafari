value=$(cat createTiger.json)

curl 'http://localhost:9000/query' \
  -F "operations=$(cat createTiger.json)" \
    -F map='{ "0": ["variables.photo"] }' \
      -F 0=@test_pic.jpg
