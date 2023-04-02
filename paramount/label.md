# label

this URL:

http://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_?format=preview

correctly returns separate series and episode title:

~~~
"cbs$SeriesTitle": "60 Minutes",
"title": "1/15/2023: Star Power, Hide and Seek, The Guru",
~~~

this URL:

http://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW?format=preview

incorrectly combines the series and episode titles:

~~~
"cbs$SeriesTitle": "SEAL Team",
"title": "SEAL Team - Tip of the Spear",
~~~

public URL:

https://paramountplus.com/shows/video/rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW

lets try Android client:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://www.paramountplus.com/shows/video/rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW/
~~~

this URL:

https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_.json?at=ABAtCJDahlKorBHq%2BQX3BPoqMWYKkTvTjVXAEi8vjHZRheaPLK2akw5ACDtkKHBfeBA%3D

correctly returns separate series and episode title:

~~~
"seriesTitle": "60 Minutes",
"label": "1/15/2023: Star Power, Hide and Seek, The Guru",
~~~

this URL:

https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW.json?at=ABAtCJDahlKorBHq%2BQX3BPoqMWYKkTvTjVXAEi8vjHZRheaPLK2akw5ACDtkKHBfeBA%3D

correctly returns separate series and episode title:

~~~
"seriesTitle": "SEAL Team",
"label": "Tip of the Spear"
~~~
