# May 2023

the response body here is normal, but the status is different:

~~~
> curl -i https://services.radio-canada.ca/ott/cbc-api/v2/assets/the-fall/s02e03
HTTP/1.1 426 Upgrade Required
Upgrade: 11.0.0

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
~~~

https://gem.cbc.ca/the-fall/s02e03

this is the HLS from the current web client:

~~~
https://cbcrcott-gem.akamaized.net/2b7d482a-9bc7-455c-8f4d-a25007a65ee4/CBC_THEFALL_S02E03.ism/desktop_master.m3u8
~~~

which comes from:

~~~
GET /media/validation/v2/?appCode=gem&connectionType=hd&deviceType=ipad&idMedia=958273&multibitrate=true&output=json&tech=hls&manifestType=desktop HTTP/1.1
Host: services.radio-canada.ca
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0
Accept: application/json, text/plain, */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: identity
Referer: https://gem.cbc.ca/
Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IjkzQURGMUNFNDhGNENCQUVCOTBDM0YyNEU5RDc0QkU5RjU0REJDMTIiLCJ4NXQiOiJrNjN4emtqMHk2NjVERDhrNmRkTDZmVk52QkkiLCJ0eXAiOiJKV1QifQ.eyJhenBDb250ZXh0IjoiY2JjZ2VtIiwiZW1haWwiOiJzcnBlbjZAZ21haWwuY29tIiwiZ2l2ZW5fbmFtZSI6InN0ZXZlbiIsImZhbWlseV9uYW1lIjoicGVubnkiLCJuYW1lIjoic3RldmVuIHBlbm55Iiwic3ViIjoiZmE4N2M1NDAtMDRlNy00NDYxLTk3N2EtYzhiZDQ3MGQxNDBhIiwicmNpZCI6ImZhODdjNTQwLTA0ZTctNDQ2MS05NzdhLWM4YmQ0NzBkMTQwYSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJpZHBVc2VySWQiOiIwZjY1MGY4ZDIyMGQ0MTJiYWFhNTc4ZDVjNTBiOTZlZCIsImxyYXQiOiI2NDVmYmE2ZS1iM2Q5LTQ4MzYtYWQwMC05MGZmMGI1YjAyMGMiLCJvaWQiOiI5ZDFkMGU1My1hZWNmLTQ4NTEtODFlNi0zZGYzMmE3YTQ1NGQiLCJpZHAiOiJyYWRpb2NhbmFkYSIsImp0aSI6ImJlZjFiNTM4LTE5NTAtNDI4My05YjI3LWIwOTZjYmMxODA3MF8yNjIwODhjNi00ZjQxLTQ4MTMtOGZiNS1mMDUwNzYzZmM5MjQiLCJub25jZSI6ImJiNDYzOTQ0LTdhODUtNDFhNS04YmI3LWEyODZhMzRjNmZjMiIsInNjcCI6ImVtYWlsIG1lZGlhLWRybXQgbWVkaWEtbWV0YSBtZWRpYS12YWxpZGF0aW9uIG1ldHJpayBvaWRjNHJvcGMgcHJvZmlsZSB0b3V0diB0b3V0di1wcmVzZW50YXRpb24gdG91dHYtcHJvZmlsaW5nIGlkLmFjY291bnQuY3JlYXRlIGlkLmFjY291bnQuZGVsZXRlIGlkLmFjY291bnQuaW5mbyBpZC5hY2NvdW50Lm1vZGlmeSBpZC5hY2NvdW50LnJlc2V0LXBhc3N3b3JkIGlkLmFjY291bnQuc2VuZC1jb25maXJtYXRpb24tZW1haWwgaWQud3JpdGUgbWVkaWEtdmFsaWRhdGlvbi5yZWFkIHN1YnNjcmlwdGlvbnMudmFsaWRhdGUgc3Vic2NyaXB0aW9ucy53cml0ZSBvdHQtcHJvZmlsaW5nIG90dC1zdWJzY3JpcHRpb24iLCJhenAiOiJmYzA1YjBlZS0zODY1LTQ0MDAtYTNjYy0zZGE4MmMzMzBjMjMiLCJ2ZXIiOiIxLjAiLCJpYXQiOjE2ODM5OTUyNDgsImF1ZCI6Ijg0NTkzYjY1LTBlZjYtNGE3Mi04OTFjLWQzNTFkZGQ1MGFhYiIsImV4cCI6MTY4NDAxNjg0OCwiaXNzIjoiaHR0cHM6Ly9sb2dpbi5jYmMucmFkaW8tY2FuYWRhLmNhL2JlZjFiNTM4LTE5NTAtNDI4My05YjI3LWIwOTZjYmMxODA3MC92Mi4wLyIsIm5iZiI6MTY4Mzk5NTI0OH0.FWelTYrxqBl3UNuY8ENh2qCSBb9BFzUoBqCcMJo_mbALtVW3emRy6mXTfqZpd_I0XVwHMU7LSOUlAH4IoWYeK1TZAstSoyfg6qTrtyFr-5JN0c0D30SThddlJVBQGZSmR7IiOQQQs2AQQeE_unESySiB6-RQRViapDGoBxzhkYk9Lr9o2D1NCMYwDytrXbsguhnY-VUXs_Nb-jbZ_b4SHrpfVrRBeh7D4S7cyRz5mf59S_lH_WiG13E2djNYxWoaq0yl7srX1Jc_JzUhrI5sQQTQxfbteWQgk0YlYBu5n1mueeLKU04l9EToQtDtSFK1H7MpqkZKjHWnPY_WGUj6Qw
x-claims-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJIYXNIRCI6IkZhbHNlIiwiVGllciI6Ik1lbWJlciIsIkhhc0FkcyI6IlRydWUiLCJSY0lkIjoiZmE4N2M1NDAtMDRlNy00NDYxLTk3N2EtYzhiZDQ3MGQxNDBhIiwiTWF4aW11bU51bWJlck9mU3RyZWFtcyI6IjUiLCJSY1RlbGNvIjoiYXVjdW4iLCJQcGlkIjoiNGVkMmUyMWRhMjRmOTg1Yzg2ODJiMjJlYTQwMjI0NTAyODQ3ODE0MjAzNWZhZjVjNjk2MzRlNWFiZGU1ZDIzYSIsImV4cCI6MTY4NDA4MTY0OX0.bpZ3o-WiDPDhw_1PUhDg94uKcxdgRkOp_4_lWPuakHI
Origin: https://gem.cbc.ca
Connection: keep-alive
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
content-length: 0
~~~
