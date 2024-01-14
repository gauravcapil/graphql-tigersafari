./test_list_tigers.sh  | jq ".data.listTigers[].TigerID" | while read line
do
sed "s;TIGERID;$line;g" listSighting.json > listSightingNew.json
curl --insecure -H 'content-type: application/json' localhost:9000/query  -d @listSightingNew.json | jq
rm  listSightingNew.json
done
