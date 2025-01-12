# apps

create Android 6 device. install user certificate. start video. after the
commercial you might get an error, try again.

## paramount tv intl

~~~
sources\com\cbs\app\BuildConfig.java
put("swisscom", "6d5824edfa1e56d6");
put("timvision", "893b6cb2e9112879");
put("vodafone", "ace4afb584a31528");

sources\com\cbs\app\config\DefaultAppSecretProvider.java
return "e55edaeb8451f737";

sources\com\cbs\app\config\SetTopBoxAppSecretProvider.java
return "e55edaeb8451f737";
~~~

- https://apkmirror.com/apk/viacomcbs-streaming/paramount-android-tv
- https://play.google.com/store/apps/details?id=com.cbs.ca

## paramount phone intl

~~~java
sources\tt\a.java
a4 = aVar.a("ab7520c40734f8aa");
~~~

- https://apkmirror.com/apk/viacomcbs-streaming/paramount-4
- https://play.google.com/store/apps/details?id=com.cbs.ca

## paramount phone us

~~~
sources\com\cbs\app\dagger\DataLayerModule.java
dataSourceConfiguration.setCbsAppSecret("4fb47ec1f5c17caa");

sources\com\cbs\app\dagger\SharedComponentModule.java
return new ci.a("{\"amazon_tablet\":\"c4abf90e3aa8131f\",
\"amazon_mobile\":\"c1353af7ed0252d8\",\"google_mobile\":\"8c4edb1155a410e4\"}");
~~~

- https://apkmirror.com/apk/cbs-interactive-inc/paramount
- https://play.google.com/store/apps/details?id=com.cbs.app

## paramount tv us

- https://apkmirror.com/apk/cbs-interactive-inc/paramount-2
- https://play.google.com/store/apps/details?id=com.cbs.ott

~~~
> play -i com.cbs.ott -leanback
details[8] = 0 USD
details[13][1][4] = 15.0.52
details[13][1][16] = Dec 10, 2024
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = http://legalterms.cbsinteractive.com/privacy
downloads = 12.95 million
name = Paramount+
size = 93.86 megabyte
version code = 211505243
~~~

result:

~~~
sources\com\cbs\app\dagger\module\DataLayerModule.java
dataSourceConfiguration.setCbsAppSecret("2a1a5e98a02a1e19");
~~~

---

## cbs phone us

- https://apkmirror.com/apk/cbs-interactive-inc/cbs
- https://play.google.com/store/apps/details?id=com.cbs.tve

~~~
> play -i com.cbs.tve
details[8] = 0 USD
details[13][1][4] = 15.0.50
details[13][1][16] = Nov 26, 2024
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = http://legalterms.cbsinteractive.com/privacy
downloads = 6.11 million
name = CBS
size = 85.97 megabyte
version code = 211505038
~~~

## cbs tv us

- https://apkmirror.com/apk/cbs-interactive-inc/cbs-android-tv
- https://play.google.com/store/apps/details?id=com.cbs.tve

~~~
> play -i com.cbs.tve -leanback
details[8] = 0 USD
details[13][1][4] = 15.0.50
details[13][1][16] = Nov 26, 2024
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = http://legalterms.cbsinteractive.com/privacy
downloads = 6.11 million
name = CBS
size = 90.59 megabyte
version code = 211505039
~~~
