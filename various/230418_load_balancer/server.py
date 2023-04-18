import sys
from flask import Flask
from datetime import datetime

app = Flask(__name__)

# server_name = "Server-1"
server_name = sys.argv[1]
port = int(sys.argv[2])


@app.route("/")
def hello():
    print(f" {datetime.now()} Server hit!")
    return server_name


if __name__ == "__main__":
    app.run(port=port)
