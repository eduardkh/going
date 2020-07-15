import requests


_session = requests.Session()

"""get the cookie"""
URL = 'http://192.168.56.132/dvwa/login.php'
res = _session.get(url=URL)
cookies = res.cookies.get_dict()
print(cookies)

"""login to the site"""
URL2 = 'http://192.168.56.132/dvwa/login.php'
DATA2 = {
    'username': 'admin',
    'password': 'password',
    'Login': 'Login',
}
res2 = _session.post(url=URL2, data=DATA2)

"""get the correct URL page"""
URL3 = 'http://192.168.56.132/dvwa/vulnerabilities/exec/'
res3 = _session.get(url=URL3)
if 'Ping for FREE' in res3.text:
    print('ok')

"""post the correct URL page"""
URL4 = 'http://192.168.56.132/dvwa/vulnerabilities/exec/'
DATA4 = {
    'ip': '127.0.0.1',
    'submit': 'submit',
}
res4 = _session.post(url=URL4, data=DATA4)
if '64 bytes from 127.0.0.1' in res4.text:
    print('ok')
