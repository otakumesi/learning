from flask import Flask
from redis import Redis, RedisError

import os
import socket

redis = Redis(host="redis", db=0, socket_connect_timeout=2, socket_timeout=2)

app = Flask(__name__)

@app.route("/")
def hello():
    try:
        visits = redis.incr("conunter")
    except RedisError:
        visits = "<i>cannot connet redis...</i>"

    html = "<h3>Hello {name}!</h3>\n" \
        "<b>Hostname:</b> {hostname}\n" \
        "<b>Visits:</b> {visits}\n"

    return html.format(name=os.getenv("NAME", "world"), hostname=socket.gethostname(), visits=visits)
if __name__ == "__main__":
    app.run("0.0.0.0", port=80)
