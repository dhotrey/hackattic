import boto3
import requests


def send_solution():
    url = "https://hackattic.com/challenges/a_global_presence/solve"

    querystring = {"access_token": "9e115aa83183d27a"}

    payload = {}
    headers = {"Content-Type": "application/json"}

    response = requests.request(
        "POST", url, json=payload, headers=headers, params=querystring
    )

    print("Submitting solution...")
    print("Got response: ", response.text)


def warmup(regions):
    for region in regions:
        print("Warming up region: ", region)
        lambda_client = boto3.client("lambda", region_name=region)
        response = lambda_client.invoke(
            FunctionName="rd_ad_test",
            InvocationType="RequestResponse",
            Payload='{"warmup": true}',
        )
        print(f"Got response {}")

    print("Warmup complete.")


def main():
    print("Hello from a-global-presence!")

    regions = [
        "us-east-1",
        "ap-south-1",
        "ap-northeast-3",
        "ap-northeast-2",
        "ap-northeast-1",
        "ap-southeast-1",
        "ap-southeast-2",
    ]

    url = "https://hackattic.com/challenges/a_global_presence/problem"
    querystring = {"access_token": "9e115aa83183d27a"}
    response = requests.request("GET", url, params=querystring)
    presense_token = response.json().get("presence_token")
    print("Got presence token: ", presense_token)


if __name__ == "__main__":
    main()
