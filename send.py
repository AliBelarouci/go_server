import time
import requests

url = "http://localhost:3000/public/getwilayainfos"

for i in range(200):
    try:
        time.sleep(0.1)
        response = requests.get(url)
        print(f"Response #{i+1}: {response.text}")
    except:
        time.sleep(0.3)
        print('error')

