# Netflix

No free trial:

https://help.netflix.com/node/16282

## android

https://play.google.com/store/apps/details?id=com.netflix.mediaclient

~~~
> play -i com.netflix.mediaclient
details[8] = 0 USD
details[13][1][4] = 8.141.1 build 13 51230
details[13][1][16] = Dec 9, 2024
details[13][1][17] = APK APK APK APK
details[13][1][82][1][1] = 7.0 and up
details[15][18] = http://www.netflix.com/privacy
downloads = 2.84 billion
name = Netflix
size = 37.64 megabyte
version code = 51230
~~~

then:

~~~
adb install-multiple (Get-ChildItem *.apk)
~~~

results:

~~~
Sdk 35 crash after 12 seconds
Sdk 34 crash after 12 seconds
Sdk 33 (TIRAMISU) crash after 12 seconds
Sdk 32 (android 12L) crash after 12 seconds
Sdk 31
Sdk 30
Sdk 29 (android 10) crash after 12 seconds
Sdk 28
Sdk 27
Sdk 26
Sdk 25
sdk 24 crash after password
~~~
