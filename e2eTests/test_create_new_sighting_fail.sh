curl -H "Authorization: test" 'http://localhost:9000/query' \
  -F "operations=$(cat createSightingTiger.json)" \
      -F map='{ "0": ["variables.photo"] }' \
            -F 0=@test_pic.jpg
	    echo first attempt should fail


