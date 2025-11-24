echo "Destroying VM"
gcloud compute instances delete hackattic-vm --zone=us-central1-a --quiet
echo "Destroyed VM successfully"

echo "Disabling SSH"
gcloud compute firewall-rules update default-allow-ssh --disabled

echo "Cleanup complete."
