import requests

_session = requests.Session()

URL = 'http://192.168.56.132/dvwa/login.php'
DATA = {
    'username': 'admin',
    'password': 'password',
    'Login': 'Login',
}
res = _session.post(url=URL, data=DATA)
print(res.cookies.get_dict())
# print(res.text)

URL = 'http://192.168.56.132/dvwa/vulnerabilities/exec/'
res2 = _session.get(url=URL)
print(res.cookies.get_dict())
if 'Ping for FREE' in res2.text:
    print('ok')


# URL = 'http://192.168.56.132/dvwa/vulnerabilities/exec/'
# res = requests.get(url=URL)
# DATA = {
#     'username': 'admin',
#     'submit': 'submit',
# }
# print(res.cookies.get_dict())
# print(res.text)

# DATA = '127.0.0.1'