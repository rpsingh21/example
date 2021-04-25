from flask import Flask
from socket import gethostname, getaddrinfo


app = Flask(__name__)

@app.route('/')
def home():
    host = gethostname()
    return f'<h3>Welcome to the Flask from</h3> <b> Host : </b> {host}'


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8000, debug=True)
