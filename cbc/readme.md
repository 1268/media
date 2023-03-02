# CBC

https://justwatch.com/ca/provider/cbc-gem

## Android client

https://play.google.com/store/apps/details?id=ca.cbc.android.cbctv

Install system certificate.

## How to create account?

Use Android client

## How to get X-Forwarded-For?

Based on this:

<https://github.com/firehol/blocklist-ipsets/blob/master/geolite2_country/country_ca.netset>

The largest Canada block is:

~~~
99.224.0.0/11
~~~

## How to get `apiKey`?

~~~
sources\vd\g.java
private static final String loginRadiusProdKey =
"3f4beddd-2061-49b0-ae80-6f1f2ed65b37";
~~~

https://github.com/skylot/jadx

## Why does this exist?

June 2 2022

https://gem.cbc.ca/media/downton-abbey/s01e05
