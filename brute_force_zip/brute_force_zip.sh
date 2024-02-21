#!/bin/sh
#get the url for zip file
chal=$(curl -s https://hackattic.com/challenges/brute_force_zip/problem?access_token=9e115aa83183d27a)
zipurl=$(echo "$chal" | jq -r '.zip_url')
# delete previous chal file
rm locked.zip
# download the challenge url
wget -q "$zipurl" -O locked.zip
# crack the password
zippwd=$(fcrackzip -v -D -p wordlists/crunch_deez_nuts.txt locked.zip --use-unzip)
# extrack the password from the output
password=$(echo "$zippwd" | awk -F '== ' '/PASSWORD FOUND!!!!/ {print $2}')
echo "password is $password"
rm *.txt
unzip -P $password locked.zip
msg=$(cat secret.txt)
json_data='{"secret":"'"$msg"'"}'
echo "json_data $json_data"
url="https://hackattic.com/challenges/brute_force_zip/solve?access_token=9e115aa83183d27a"
res=$(curl -X POST -H "Content-Type: application/json" -d "$json_data" "$url")
echo "Response $res"

