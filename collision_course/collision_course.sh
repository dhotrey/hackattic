#!/bin/sh
#get the url for zip file
chal=$(curl -s https://hackattic.com/challenges/collision_course/problem?access_token=9e115aa83183d27a)
include=$(echo "$chal" | jq -r '.include')
echo "$include" > prefix_file
rm *.bin
./md5_fastcoll -q -p prefix_file

f1=$(cat msg1.bin | base64 | jq -sR)
f2=$(cat msg2.bin | base64 | jq -sR)

json_data='{"files":['$f1','$f2']}'
echo "json_data $json_data"
url="https://hackattic.com/challenges/collision_course/solve?access_token=9e115aa83183d27a"
res=$(curl -X POST -H "Content-Type: application/json" -d "$json_data" "$url")
echo "Response $res"


