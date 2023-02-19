import textwrap
import random
import os
from flask import Flask, request
import time
app = Flask(__name__)

# not vuln
try:
    secret = os.environ['SECRETT']
except:
    secret = "nvm"
if secret == "nvm":
    port = 4000
else:
    port = 5000

#juicy part
@app.route("/")
def hello():
    x=random.randint(0,len(secret))
    try:
        return(textwrap.shorten(secret, width=x, placeholder="nope"))
    except:
        return("interesting humm..")

if __name__ == '__main__':
    app.run("0.0.0.0",port=port,debug=True)