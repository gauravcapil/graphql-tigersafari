
curl -H "Authorization: test" 'http://localhost:9000/query' \
  -F "operations=$(cat createSightingTiger.json)" \
    -F map='{ "0": ["variables.photo"] }' \
      -F 0=@test_pic.jpg
echo first attempt should fail


function createSightings {
TigerID=$1
for((i=10;i<20;i++))
do
cp createSightingTiger.json createSightingTigerNew.json
sed -i -e "s;TIGERID;$TigerID;g" -e "s;SEENTIME;$i;g" -e "s;SEENLONG;$i;g" -e "s;SEENLAT;$i;g"  createSightingTigerNew.json

curl -H 'Authorization: '$(./test_login.sh | jq ".data.login.token" | sed "s;\";;g")  'http://localhost:9000/query' \
  -F "operations=$(cat createSightingTigerNew.json)" \
      -F map='{ "0": ["variables.photo"] }' \
            -F 0=@test_pic.jpg
done


}

createSightings n1
createSightings n2
