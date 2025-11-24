echo "Creating new VM instance"
gcloud compute instances create hackattic-vm \
    --machine-type=e2-micro \
    --zone=us-central1-a \
    --image-family=debian-12 \
    --image-project=debian-cloud \
    --boot-disk-size=30GB \
    --tags=http-server,https-server

echo "New VM created successfully"
echo "Enablind SSH"
gcloud compute firewall-rules update default-allow-ssh --no-disabled
echo "Enabled SSH successfully"
