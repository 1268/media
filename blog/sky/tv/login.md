# login

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

and:

~~~py
cookie_url = 'https://raw.githubusercontent.com/sunsettrack4/zattoo_tvh/additional/cookies.json'
~~~

and:

~~~py
cookie_page = requests.get(cookie_url, timeout=5, verify=False)
~~~

and:

~~~py
cookie_check = cookie_page.json()
~~~

and:

~~~py
login_url = 'https://show.sky.ch/de/login?forceClassicalTvLogin=True'
~~~

and:

~~~py
login_page = requests.get(
   login_url, timeout=5, headers=login_headers, verify=False
)
~~~

and:

~~~py
asp_cookie = login_page.cookies.get(cookie_check["asp"])
~~~

and:

~~~py
cookie_token = login_page.cookies.get(cookie_check["rvt"])
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

and:

~~~py
login_page_parse = BeautifulSoup(login_page.content, 'html.parser')
~~~

and:

~~~py
app_token_reference = login_page_parse.find(
   'input', {'name': cookie_check["rvtp"]}
)
~~~

and:

~~~py
page_token = app_token_reference.get("value")        
~~~

and:

~~~py
data = {
   'username': __login, 'password': __password,
   'returnUrl': 'Home/HomeTv', 'subscriptionUrl': '/de/subscription',
   cookie_check["rvtp"]: page_token,
}
~~~

and:

~~~py
login_page = requests.post(
   auth_url, timeout=5, headers=login_headers,
   cookies=cookies,
   data=data,
   allow_redirects=False, verify=False
)
~~~
