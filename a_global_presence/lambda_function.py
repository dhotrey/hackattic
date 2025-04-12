import http.client
import os


def lambda_handler(event, context):
    logs = []
    lambda_region = os.environ["AWS_REGION"]
    logs.append(f"received event {event}")
    logs.append(f"Lambda running in {lambda_region}")

    if event["eventType"] == "warmup":
        logs.append("warmup event received")
        return {"event": event, "logs": logs}

    elif event["eventType"] == "hack":
        conn = http.client.HTTPSConnection("hackattic.com")
        conn.request("GET", f"/_/presence/{event["presence_token"]}")
        res = conn.getresponse()
        data = res.read()
        resp = data.decode("utf-8")
        logs.append(res)
        print(resp)
        return {
            "statusCode": 200,
            "event": event,
            "logs": logs,
        }

    return {"statusCode": 200, "event": event}
