
function createSightings {
TigerID=$1
i=$2
newd=$(date +%Y-%m-%d" "%H:%M:%S.000)
cp createSightingTiger.json createSightingTigerNew.json
sed -i -e "s;TIGERID;$TigerID;g" -e "s;SEENTIME;$newd;g" -e "s;SEENLONG;$i;g" -e "s;SEENLAT;0.$i;g"  createSightingTigerNew.json

curl -H 'Authorization: '$(./test_login.sh | jq ".data.login.token" | sed "s;\";;g")  'http://localhost:9000/query' \
  -F "operations=$(cat createSightingTigerNew.json)" \
      -F map='{ "0": ["variables.photo"] }' \
            -F 0=@test_pic.jpg
rm createSightingTigerNew.json
echo "adding last seen at in a file for checking list tiger api later: $newd, for $i"


}

createSightings n1 $1
