from flask import Flask, render_template, request, send_file, make_response
from config import flag
from Crypto.Util.number import *
from internalMatch import getParams

a, b, p = getParams()

def func(x,a=a,b=b,p=p):
    return (a*x+b)%p

def matching(x,y):
    r = getRandomNBitInteger(20)
    for _ in range(r):
        x = func(x)
    return x==y

app = Flask(_name_, static_url_path='/static')


@app.route("/",methods=['GET',"POST"])
def index():
    if request.method == "GET":
        return render_template('index.html', a=a,b=b,p=p)
    else :
        try:
            x = int(request.form['your_name'])
            y = int(request.form['your_crush'])
            tmp = matching(x,y)
            print(tmp)
            if(tmp):
                return render_template("index.html",status=flag, a=a,b=b,p=p)
            else:
                return render_template('index.html',status="You Didn't match your crush vibe you weak", a=a,b=b,p=p)
        except Exception as e:
                print(e)
                return render_template('index.html',error="Names Are Numbers in Crypto", a=a,b=b,p=p)
        
@app.get("/code")
def code():
    return send_file(_file_)