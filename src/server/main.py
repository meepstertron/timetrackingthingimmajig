import os

import fastapi
import uvicorn
import requests
import dotenv

app = fastapi.FastAPI()

dotenv.load_dotenv()

@app.get("/oauth/callback")
def hcauth_callback(request: fastapi.Request):
    print("callback")

    code = request.query_params.get("code")

    x = requests.post("https://auth.hackclub.com/oauth/token", json={
        "client_id": os.getenv("HCAUTH_CLIENTID"),
        "client_secret": os.getenv("HCAUTH_SECRET"),
        "redirect_uri": os.getenv("HCAUTH_REDIRECT_URI"),
        "code": code,
        "grant_type": "authorization_code",
    })

    print(x.json())

    return {"message": "Hello World"}



if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=7788)