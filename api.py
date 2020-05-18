from flask import Flask
from flask import request

app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Route'

@app.route('/hello')
def hello():
    return 'Hello, World'    

@app.route('/records', methods=["POST"])
def add_record():

    response = {
        "intReturnCode": 0,
        "strMessage": "Successful",
        "arrRecords": [{
            "name":"navigator",
            "type":"A",
            "content":"41.45.34.67",
            "ttl": 3600,
            "prio":10
        },
        {
            "name":"zipkin",
            "type":"A",
            "content":"35.89.34.67",
            "ttl": 3600,
            "prio":10
        }]}

    print(request.form['key'])
    print(request.form['sld'])
    print(request.form['tld'])
    return response