# login

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
login_page1 = requests.post(
   auth_url, timeout=5, headers=login_headers,
   cookies=cookies,
   data=data,
   allow_redirects=False, verify=False
)
~~~
