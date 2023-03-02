# Bandcamp

## Android client

2023-02-26

~~~
com.bandcamp.android
~~~

Bandcamp app as of the date above, does not monitor Bandcamp URLs:

~~~xml
<intent-filter>
   <action android:name="android.intent.action.VIEW"/>
   <category android:name="android.intent.category.DEFAULT"/>
   <category android:name="android.intent.category.BROWSABLE"/>
   <data android:scheme="x-bandcamp" android:host="open"/>
   <data android:scheme="x-bandcamp" android:host="messages"/>
   <data android:scheme="x-bandcamp" android:host="feed"/>
   <data android:scheme="x-bandcamp" android:host="linked_paypal_email"/>
   <data android:scheme="x-bandcamp" android:host="show_merch"/>
   <data android:scheme="x-bandcamp" android:host="show_tralbum"/>
   <data android:scheme="x-bandcamp" android:host="show_fan"/>
   <data android:scheme="x-bandcamp" android:host="login"/>
   <data android:scheme="x-bandcamp" android:host="signup"/>
</intent-filter>
~~~

so deep linking is not possible. The API used with Bandcamp Android requires
IDs for everything. This is the case going back to 2016 at least.

## Why does this exist?

2023-02-26

https://solarfields.bandcamp.com/track/sol-remix-remastered

Solar Fields hasn’t created any albums yet:

https://soundcloud.com/solarfields/albums

20 unavailable videos are hidden:

https://youtube.com/playlist?list=PLem8ejSphCJI6MybaQTuqIUIiWq75vCHg
