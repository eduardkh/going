import os
from os.path import isfile

path = 'c:\\suka\\'

files = []
# r=root, d=directories, f = files
for r, d, f in os.walk(path):
    for _file in f:
        if '.txt' in _file:
            files.append(os.path.join(r, _file))

counter = 1
for f in files:
    base = '\\'.join(f.split('\\')[0:-1])
    new_file = f.split('\\')[-1]
    print(f'{base}\\new_name_{counter}.txt')
    os.rename(f, f'{base}\\new_name_{counter}.txt')
    counter += 1
