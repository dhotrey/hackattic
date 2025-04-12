import boto3
import requests
from deploy import REGIONS
import json
import time


def trigger(event, region):
    t1 = time.time()
    client = boto3.client("lambda", region_name=region)
    response = client.invoke(
        FunctionName="hackattic",
        InvocationType="RequestResponse",
        Payload=json.dumps(event),
    )
    t2 = time.time()
    ttaken = t2 - t1
    return response, ttaken


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


def main():
    print("Hello from a-global-presence!")
    url = "https://hackattic.com/challenges/a_global_presence/problem"
    access_token = ""
    querystring = {"access_token": access_token}

    print("Warming up lambda functions")
    warm_up_event = {"eventType": "warmup"}
    for region in REGIONS:
        resp, timetaken = trigger(warm_up_event, region)
        print(f"Warmed up hackattic lambda in {region} in {timetaken}s")
        print(json.loads(resp["Payload"].read()))

    print("warmed up all lambda functions successfully!")
    response = requests.request("GET", url, params=querystring)
    presence_token = response.json().get("presence_token")
    start_time = time.time()
    print("Got presence token: ", presence_token)

    actual_event = {"eventType": "hack", "presence_token": presence_token}

    for region in REGIONS:
        resp, timetaken = trigger(actual_event, region)
        print(f"Triggered {region} lambda ")

    print("sending solution")
    send_solution()
    end_time = time.time()
    print(f"Time elapsed aftert requesting presense token {end_time-start_time}")


if __name__ == "__main__":
    main()
