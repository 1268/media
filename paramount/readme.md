# Paramount+

## Android client

https://play.google.com/store/apps/details?id=com.cbs.app

Create device with Android API 23. Install user certificate. Start video. After
the commercial you might get an error, try again.

## CBS

212,360 titles:

https://justwatch.com/us/provider/cbs

which is the same as no filters:

https://justwatch.com/us

3,183 titles:

https://justwatch.com/us/provider/paramount-plus

## TRY PARAMOUNT+

1. https://paramountplus.com
2. click TRY IT FREE
3. click CONTINUE
4. make sure MONTHLY is selected, then under Essential click SELECT PLAN
5. if you see a bundle screen, click MAYBE LATER
6. click CONTINUE
7. uncheck Yes, I would like to receive marketing
8. click CONTINUE
9. click START PARAMOUNT+
10. https://paramountplus.com
11. select a profile
12. under profile click Account
13. click Cancel Subscription
14. click CONTINUE TO CANCEL
15. click I understand
16. click YES, CANCEL
17. click the first option
18. click COMPLETE CANCELLATION

## How to get app\_secret?

~~~
com\cbs\app\dagger\DataLayerModule.java
dataSourceConfiguration.setCbsAppSecret("6c70b33080758409");
~~~

## How to get secret\_key?

~~~
com\cbs\app\androiddata\retrofit\util\RetrofitUtil.java
SecretKeySpec secretKeySpec = new SecretKeySpec(b("302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"), "AES");
~~~
