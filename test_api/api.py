from flask import Flask
from flask import request

app = Flask(__name__)

records = []


@app.route('/get-records', methods=["POST"])
def get_records():
    key = request.form['key']
    sld = request.form['sld']
    tld = request.form['tld']

    response = {
        "intReturnCode": 0,
        "strMessage": "Successful",
        "arrRecords": records}

    return response


@app.route('/update-records', methods=["POST"])
def update_records():

    key = request.form['key']
    sld = request.form['sld']
    tld = request.form['tld']

    print(request.form)

    index = 1

    records.clear()

    while True:
        if form_has_index(index, request.form):
            record = record_from_request(index, request.form)
            records.append(record_from_request(index, request.form))
            index += 1
        else:
            break

    response = {
        "intReturnCode": 0,
        "strMessage": "Successful"}

    return response


def form_has_index(index, form):
    return f'name{index}' in form.keys()


def record_from_request(index, form):
    return {
        "name": request.form[f'name{index}'],
        "type": request.form[f'type{index}'],
        "content": request.form[f'content{index}'],
        "ttl": request.form[f'ttl{index}'],
        "prio": request.form[f'prio{index}']
    }
