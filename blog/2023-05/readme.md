# May 2023

the response body here is normal:

~~~
curl -i https://services.radio-canada.ca/ott/cbc-api/v2/assets/the-fall/s02e03
~~~

but the response status is different:

~~~
HTTP/1.1 426 Upgrade Required
Upgrade: 11.0.0
~~~


PS D:\git\mech\cbc> go test -run Ass
GET https://services.radio-canada.ca/ott/cbc-api/v2/assets/the-fall/s02e03
The Fall - S2 E3 - Beauty Hath Strange Power
{
 "playSession": {
  "URL": "https://services.radio-canada.ca/media/validation/v2?appCode=gem&idMedia=958273&manifestType=desktop&output=json&tech=hls"
 },
 "Series": "The Fall",
 "Title": "Beauty Hath Strange Power",
 "Season": 2,
 "Episode": 3,
 "Credits": {
  "releaseDate": "2014"
 }
}
PASS
ok      2a.pages.dev/mech/cbc   1.330s
PS D:\git\mech\cbc>






