import boto3
import zipfile

REGIONS = [
    "eu-central-1",
    "ap-northeast-3",
    "ap-northeast-2",
    "ap-southeast-1",
    "ap-southeast-2",
    "us-east-1",
    "ap-south-1",
]


def create_zipfile():
    with zipfile.ZipFile("lambda_deployment.zip", "w") as zipf:
        zipf.write("lambda_function.py")

    print("Created deployment zip ", "lambda_deployment.zip")


def deploy_functions():
    create_zipfile()

    with open("lambda_deployment.zip", "rb") as f:
        zipped_code = f.read()

    for region in REGIONS:
        lambda_client = boto3.client("lambda", region_name=region)
        response = lambda_client.update_function_code(
            FunctionName="hackattic", ZipFile=zipped_code
        )

        if response["ResponseMetadata"]["HTTPStatusCode"] == 200:
            print(f"Deployed function hackattic to region {region} successfully!")
        else:
            print(f"Something went wrong while deploying code to region {region}")
            print(response)

        print("-" * 80)


if __name__ == "__main__":
    deploy_functions()
