# tv

if you request TV app, phone app is returned:

~~~
> play -i homedia.sky.sport -leanback
details[8] = 0 USD
details[13][1][4] = 1.18.1.142
details[13][1][16] = Jan 21, 2025
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 8.0 and up
details[15][18] = https://support.sky.ch/hc/en-us/articles/9520105066140
downloads = 468.36 thousand
name = Sky
size = 35.09 megabyte
version code = 584

> play -i homedia.sky.sport
details[8] = 0 USD
details[13][1][4] = 1.18.1.142
details[13][1][16] = Jan 21, 2025
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 8.0 and up
details[15][18] = https://support.sky.ch/hc/en-us/articles/9520105066140
downloads = 468.36 thousand
name = Sky
size = 35.09 megabyte
version code = 584
~~~

~~~py
auth_url = 'https://show.sky.ch/de/Authentication/Login'
~~~

and:

~~~
login_headers.update({"referer": "https://show.sky.ch/de/tv/", "tv": "Emulator"})
~~~

and:

~~~py
headers = {
   'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:127.0) Gecko/20100101 Firefox/127.0',
   "accept-language": "de"
}
~~~

and:

~~~py
login_headers = headers
~~~

---------------------------------------------------------------------------------

~~~py
login_page = requests.post(auth_url, timeout=5, headers=login_headers,
cookies=cookies, data=data, allow_redirects=False, verify=False)
~~~

and:

~~~py
cookies = {
   cookie_check["asp"]: asp_cookie,
   cookie_check["rvt"]: cookie_token,
   "SkyTvDevice": '''
{
  "isSky": true,
  "type": {
    "code": "Desktop"
  },
  "year": "",
  "keys": {
    "enter": 13,
    "back": 461,
    "up": 38,
    "down": 40,
    "left": 37,
    "right": 39,
    "play": 415,
    "pause": 19,
    "playPause": -1,
    "ff": 417,
    "rew": 412,
    "stop": 413,
    "search": -1,
    "rew10": -1,
    "ff10": -1,
    "key0": -1,
    "key1": -1,
    "key2": -1,
    "key3": -1,
    "key4": -1,
    "key5": -1,
    "key6": -1,
    "key7": -1,
    "key8": -1,
    "key9": -1
  }
}
   '''
}
~~~
