#!/bin/sh

challenge=$(curl -s https://hackattic.com/challenges/hosting_git/problem?access_token="$token")
push_token=$(echo "$challenge" | jq -r '.push_token')
ssh_key=$(echo "$challenge" | jq -r '.ssh_key')
username=$(echo "$challenge" | jq -r .'username')
repo_path=$(echo "$challenge" | jq -r .'repo_path')

echo "Creating new user"
sudo useradd -m -s /bin/bash "$username"

cd /home/$username

echo "Creating git repo . . . "
sudo -u "$username" git init --bare $(echo "$repo_path") > /dev/null

echo "creating ssh dir"
sudo -u "$username" mkdir .ssh
sudo -u "$username" chmod 700 .ssh
echo "$ssh_key" | sudo -u "$username" tee -a .ssh/authorized_keys > /dev/null
sudo -u "$username" chmod 600 .ssh/authorized_keys
echo "setup complete!"

echo "Triggering the push"

curl -X POST https://hackattic.com/_/git/"$push_token" \
     -H "Content-Type: application/json" \
     -d '{"repo_host":"146.148.85.172"}'

echo "push triggered successfully!"

cd $repo_path
git config --global --add safe.directory /home/$username/$repo_path
secret=$(git show master:solution.txt)
echo "Got secret : $secret"

echo "Posting solution"

curl -X POST https://hackattic.com/challenges/hosting_git/solve?access_token="$token" \
     -H "Content-Type: application/json" \
     -d "{\"secret\":\"$secret\"}"

echo "Challenge completed successfully"
