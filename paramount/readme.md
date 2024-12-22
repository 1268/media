# Paramount+

## Android client

create Android 6 device. install user certificate. start video. after the
commercial you might get an error, try again.

US:

https://play.google.com/store/apps/details?id=com.cbs.app

INTL:

https://play.google.com/store/apps/details?id=com.cbs.ca

## try paramount+

1. paramountplus.com
2. try it free
3. continue
4. make sure monthly is selected, then under essential click select plan
5. if you see a bundle screen, click maybe later
6. continue
7. uncheck yes, i would like to receive marketing
8. continue
9. start paramount+

## How to get app\_secret?

~~~
sources\com\cbs\app\dagger\DataLayerModule.java
dataSourceConfiguration.setCbsAppSecret("a624d7b175f5626b");
~~~

## How to get secret\_key?

~~~
com\cbs\app\androiddata\retrofit\util\RetrofitUtil.java
SecretKeySpec secretKeySpec = new SecretKeySpec(b("302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"), "AES");
~~~

## link.theplatform.com

why do we need link.theplatform.com? because its the only anonymous option.
logged out the web client has this request:

https://paramountplus.com/shows/mayor-of-kingstown/video/xhr/episodes/page/0/size/18/xs/0/season/3

logged in the web client embeds MPD in HTML:

https://paramountplus.com/movies/video/Oo75PgAbcmt9xqqn1AMoBAfo190Cfhqi

Android client needs cookie for INTL requests:

---

paramountplus.com/movies/video/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ
content_id: "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
json.itemList[0].downloadCountrySet[0].code = "DE";
json.itemList[0].downloadCountrySet[1].code = "MQ";
json.itemList[0].downloadCountrySet[2].code = "IM";
json.itemList[0].downloadCountrySet[3].code = "IT";
json.itemList[0].downloadCountrySet[4].code = "BL";
json.itemList[0].downloadCountrySet[5].code = "PM";
json.itemList[0].downloadCountrySet[6].code = "VA";
json.itemList[0].downloadCountrySet[7].code = "JE";
json.itemList[0].downloadCountrySet[8].code = "FR";
json.itemList[0].downloadCountrySet[9].code = "GG";
json.itemList[0].downloadCountrySet[10].code = "IE";
json.itemList[0].downloadCountrySet[11].code = "MF";
json.itemList[0].downloadCountrySet[12].code = "SM";
json.itemList[0].downloadCountrySet[13].code = "GF";
json.itemList[0].downloadCountrySet[14].code = "GP";
json.itemList[0].downloadCountrySet[15].code = "CH";
json.itemList[0].downloadCountrySet[16].code = "YT";
json.itemList[0].downloadCountrySet[17].code = "GB";
json.itemList[0].downloadCountrySet[18].code = "AT";
json.itemList[0].downloadCountrySet[19].code = "RE";

content_id: "ssc3CuuS4mrQ7EyVXILH0FEQSi5yBAsA",
json.itemList[0].downloadCountrySet[0].code = "JE";
json.itemList[0].downloadCountrySet[1].code = "GB";
json.itemList[0].downloadCountrySet[2].code = "IM";
json.itemList[0].downloadCountrySet[3].code = "GG";
json.itemList[0].downloadCountrySet[4].code = "IE";

content_id: "WNujiS5PHkY5wN9doNY6MSo_7G8uBUcX",
json.itemList[0].downloadCountrySet[0].code = "AU";

need login
GET https://www.intl.paramountplus.com/apps-api/v3.1/androidtv/irdeto-control/session-token.json?contentId=Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ&model=sdk_google_atv_x86&firmwareVersion=9&version=15.0.28&platform=PPINTL_AndroidTV&locale=en-us&at=ABBoPFHuygkRnnCKELRhypuq5uEAJvSiVATsY9xOASH88ibse11WuoLrFnSDf0Bv7EY%3D HTTP/2.0
cookie: CBS_COM=N0Q5NkYyRTE2QTRGQUJEQTE1QkYzQTQwREVEMjQxNUY3RjYyQkM0MkVDMzM2OD...




works anon
GET https://www.paramountplus.com/apps-api/v3.1/androidphone/irdeto-control/anonymous-session-token.json?contentId=Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ&model=AOSP%20on%20IA%20Emulator&firmwareVersion=9&version=15.0.28&platform=PP_AndroidApp&locale=en-us&locale=en-us&at=ABBoPFHuygkRnnCKELRhypuq5uEAJvSiVATsY9xOASH88ibse11WuoLrFnSDf0Bv7EY%3D HTTP/2.0
cookie: CBS_COM=N0Q5NkYyRTE2QTRGQUJEQTE1QkYzQTQwREVEMjQxNUY3RjYyQkM0MkVDMzM2OD...

works anon
GET https://www.intl.paramountplus.com/apps-api/v3.0/androidtv/movies/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ.json?includeTrailerInfo=true&includeContentInfo=true&locale=en-us&at=ABDSbrWqqlbSWOrrXk8u9NaNdokPC88YiXcPvIFhPobM3a%2FJWNOSwiCMklwJDDJq4c0%3D HTTP/2.0
cookie: CBS_COM=N0Q5NkYyRTE2QTRGQUJEQTE1QkYzQTQwREVEMjQxNUY3RjYyQkM0MkVDMzM2OD...

works anon
GET https://www.intl.paramountplus.com/apps-api/v2.0/androidtv/video/cid/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ.json?locale=en-us&at=ABA3WXXZwgC0rQPN9WtWEUmpHsGCFJb6NP4tGjIFVLTuScgId9WA3LdC44hdHUJysQ0%3D HTTP/2.0
cookie: CBS_COM=N0Q5NkYyRTE2QTRGQUJEQTE1QkYzQTQwREVEMjQxNUY3RjYyQkM0MkVDMzM2OD...



