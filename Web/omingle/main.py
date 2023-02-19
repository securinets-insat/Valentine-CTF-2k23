from internal import cupid
from flask import Flask, render_template, request, send_file, make_response
import subprocess
import re

app = Flask(__name__, static_url_path='/static')

@app.get("/")
def index():
    visitor = request.cookies.get('visitor')
    out = cupid(visitor, subprocess)
    resp = make_response(render_template('index.html'))

    if out:
        resp.set_cookie('notification', out)
    return resp


@app.post("/")
def post():
    print("request.remote_addr", request.remote_addr)
    name = request.form['pretty-name']
    if not re.match(r'^[</>()\'a-zA-Z0-9_]+$', name):
        return render_template('index.html', name='mmmmm no')
    return render_template('index.html', name="Hi "+name)

@app.get("/code")
def code():
    return send_file(__file__)