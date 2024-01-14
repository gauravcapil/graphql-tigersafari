./test_create_user.sh
./test_create_tiger.sh
./test_create_tiger2.sh
./test_create_new_sighting.sh
./test_list_sighting.sh
./test_list_tigers.sh |   jq '.data.listTigers[]| "\(.SeenAt)|\(.Name)"' | while read line
do
	tigername=$( echo $line | awk -F\| '{print $NF}' | sed "s;\";;g")
	sightingreported=$( echo $line | awk -F\| '{print $1}' | sed "s;\";;g")
	echo tiger to check sightng value: $tigername, reported last sighting was: $sightingreported
	sightingposted=$(cat lastsightingposted_$tigername)
	if [[ $sightingposted != $sightingreported ]]
	then
		echo "Error: Sighting reported are incorrect for $tigername"
		exit 1
	fi
done

