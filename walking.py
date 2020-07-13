import os
import subprocess

path = 'c:\\suka\\'
_ffmpeg = "C:\\ffmpeg-20200628-4cfcfb3-win64-static\\bin\\ffmpeg.exe"
files = []
# r=root, d=directories, f = files
for r, d, f in os.walk(path):
    for _file in f:
        if '.mp4' in _file:
            files.append(os.path.join(r, _file))

for f in files:
    # print(f)
    # print('{} -i "{}" -c:v libx264 -c:a libmp3lame -b:a 384K "{}.avi"'.format(_ffmpeg, f, f))
    subprocess.Popen('{} -i "{}" -c:v libx264 -c:a libmp3lame -b:a 384K "{}.avi"'.format(_ffmpeg, f, f))
