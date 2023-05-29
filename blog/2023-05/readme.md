# May 2023

this works:

~~~
cbc -a http://gem.cbc.ca/downton-abbey/s01e01
~~~

but this fails:

~~~
> cbc -a http://gem.cbc.ca/downton-abbey/s01e03
panic: Cette plate-forme n'est pas supportée
~~~

but if I visit the same page in the browser, it then works with the tool. what
is the browser doing? here is another that fails currently:

~~~
> cbc -a http://gem.cbc.ca/downton-abbey/s06e05 -log 2
GET /ott/catalog/v2/gem/show/downton-abbey/s06e05?device=web HTTP/1.1
Host: services.radio-canada.ca

GET /media/validation/v2?appCode=gem&idMedia=929058&manifestType=desktop&output=json HTTP/1.1
Host: services.radio-canada.ca
X-Claims-Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJIYXNIRCI6IkZhbHNlIiwiVGllciI6Ik1lbWJlciIsIkhhc0FkcyI6IlRydWUiLCJSY0lkIjoiZmE4N2M1NDAtMDRlNy00NDYxLTk3N2EtYzhiZDQ3MGQxNDBhIiwiTWF4aW11bU51bWJlck9mU3RyZWFtcyI6IjUiLCJSY1RlbGNvIjoiYXVjdW4iLCJQcGlkIjoiNGVkMmUyMWRhMjRmOTg1Yzg2ODJiMjJlYTQwMjI0NTAyODQ3ODE0MjAzNWZhZjVjNjk2MzRlNWFiZGU1ZDIzYSIsImV4cCI6MTY4NTQ1NjQ1MX0.LzdNbOvkpoY0qAceBxEjioHj6Y3eDyw-hQ1aEPmcJuQ
X-Forwarded-For: 99.224.0.0

panic: Cette plate-forme n'est pas supportée
~~~

here is another that currently fails:

~~~
> cbc -a http://gem.cbc.ca/downton-abbey/s06e06
GET https://services.radio-canada.ca/ott/catalog/v2/gem/show/downton-abbey/s06e06?device=web
GET https://services.radio-canada.ca/media/validation/v2?appCode=gem&idMedia=929064&manifestType=desktop&output=json
panic: Cette plate-forme n'est pas supportée
~~~
