#!/bin/sh
#get the url for zip file
chal=$(curl -s https://hackattic.com/challenges/reading_qr/problem?access_token=9e115aa83183d27a)
qrurl=$(echo "$chal" | jq -r '.image_url')
# download the challenge url
rm qr.png
wget -q "$qrurl" -O qr.png
# using zbarimg for processing the qr code because the challenge is only 10 points
qrData=$(zbarimg --quiet qr.png)
# echo "$qrData"
code=$(echo "$qrData" | sed 's/.*://')
echo "QR data : $code"
json_data='{"code":"'"$code"'"}'
echo "json_data $json_data"
url="https://hackattic.com/challenges/reading_qr/solve?access_token=9e115aa83183d27a"
res=$(curl -X POST -H -S "Content-Type: application/json" -d "$json_data" "$url")
echo "Response $res"

