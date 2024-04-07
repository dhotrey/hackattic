#!/bin/bash
#get the url for zip file
chal=$(curl -s https://hackattic.com/challenges/the_redis_one/problem?access_token=9e115aa83183d27a)
rdb=$(echo "$chal" | jq -r '.rdb')
getType=$(echo "$chal" | jq -r '.requirements' | jq -r '.check_type_of')
echo "$rdb" | base64 -d > malformed.rdb
printf "%b" "\x52\x45\x44\x49\x53" > header
tail -c +6 malformed.rdb > remaining_data
cat header remaining_data > chal.rdb
rm malformed.rdb header remaining_data
