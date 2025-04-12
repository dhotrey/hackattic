import cv2.data
import requests
import time
import cv2

def save_image_grid_to_disk(imageGrid):
    for row_idx in range(len(imageGrid)):
        for img_idx in range(len(imageGrid[row_idx])):
            name = f"Out/{row_idx}_{img_idx}.png"
            print(f"Creating {name}")
            cv2.imwrite(name,imageGrid[row_idx][img_idx])

def process_image(imagename):
    start_time = time.time()
    image = cv2.imread(imagename)
    assert image is not None
    height, width, channels = image.shape
    print(height,width)
    img2 = image
    H_SIZE = 8
    W_SIZE = 8
    imageGrid = []
    for ih in range(H_SIZE):
        subImageList = []
        for iw in range(W_SIZE):
            x = width/W_SIZE*iw
            y = height/H_SIZE*ih
            h = (height/H_SIZE)
            w = (width/W_SIZE)
            img = image[int(y):int(y+h),int(x):int(x+w)]
            subImageList.append(img)
            image = img2
        imageGrid.append(subImageList)
    end_time = time.time()

    print(f"Processed images in {end_time-start_time}")
    return imageGrid

def run_face_id(imageGrid):
    detected_faces = []
    for row_idx in range(len(imageGrid)):
        for img_idx in range(len(imageGrid[row_idx])):
            image = imageGrid[row_idx][img_idx]
            gray_image = cv2.cvtColor(image,cv2.COLOR_BGR2GRAY)
            face_classifier = cv2.CascadeClassifier(cv2.data.haarcascades + "haarcascade_frontalface_default.xml")

            face = face_classifier.detectMultiScale(gray_image,scaleFactor=1.1,minNeighbors=5,minSize=(40,40))
            if len(face) != 0:
                # print(f"{row_idx}-{img_idx}.png -> : {face}")
                detected_faces.append([row_idx,img_idx])
    return detected_faces

def send_solution(sol):
    url = "https://hackattic.com/challenges/basic_face_detection/solve"
    access_tok = ""
    querystring = {"access_token":access_tok}
    payload = {"face_tiles": sol}
    headers = {"Content-Type": "application/json"}
    response = requests.request("POST", url, json=payload, headers=headers, params=querystring)
    print(response.json())

def main():
    print("Hello from basic-face-detection!")
    access_token = ""
    url = "https://hackattic.com/challenges/basic_face_detection/problem"
    querystring = {"access_token": access_token}
    t1 = time.time()
    response = requests.request("GET", url, params=querystring)
    t2 = time.time()
    img_url = response.json().get("image_url")
    print(f"Got image url {img_url} in {t2-t1}s")
    t3 = time.time()
    img_response = requests.request("GET", img_url)
    t4 = time.time()
    print(f"Downloaded image successfully in {t4-t3}s")

    with open("faces.png", "wb") as img:
        img.write(img_response.content)
    print("Image saved to file 'faces.png' ")


if __name__ == "__main__":
    main()
    cropped_img_list = process_image("faces.png")
    face_locations = run_face_id(cropped_img_list)
    send_solution(face_locations)
    # save_image_grid_to_disk(cropped_img_list)