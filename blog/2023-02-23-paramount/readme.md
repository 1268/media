# Paramount+

this:

https://paramountplus.com/movies/no-country-for-old-men/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC

redirects here:

https://paramountplus.com/movies/video/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC

from United States, I get 404. From United Kingdom or Ireland I get prompt to
Sign In. my tool fails from United States:

~~~
> paramount -b yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC
panic: 412 Precondition Failed

> paramount -b yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC -a 1
panic: 412 Precondition Failed

> paramount -b yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC -a 2
panic: 412 Precondition Failed
~~~

and fails from Ireland:

~~~
> paramount -b yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC -a 2
panic: 403 Forbidden
~~~

these all fail for United States, United Kingdom and Ireland:

~~~
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?assetTypes=StreamPack&formats=MPEG4,M3U
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?assetTypes=HLS_CLEAR&formats=MPEG4,M3U
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?assetTypes=DASH_CENC&formats=MPEG-DASH
~~~

these fail from United States:

~~~
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?format=SMIL&formats=M3U+none,MPEG4
<param name="exception" value="NoAssetTypeFormatMatches"/>

link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?format=SMIL&formats=MPEG-DASH,M3U,MPEG4
<param name="exception" value="GeoLocationBlocked"/>
~~~

these work from United Kingdom:

~~~
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?format=SMIL&formats=MPEG-DASH,M3U,MPEG4
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?formats=M3U
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?formats=MPEG-DASH,M3U
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?formats=MPEG-DASH,M3U,MPEG4
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?formats=MPEG-DASH,MPEG4
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?formats=M3U,MPEG4
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?formats=MPEG4,M3U
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC?formats=MPEG-DASH
~~~

this works:

<http://can.cbs.com/thunder/player/videoPlayerService.php?partner=cbs&contentId=wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_>

this fails from United States, United Kingdom and Ireland:

http://can.cbs.com/thunder/player/videoPlayerService.php?partner=cbs&contentId=yzKMWQVnTuK9dZrZCWJUUWEU8hF0yAgC

this fails from United Kingdom:

http://paramountplus.com/apps-api/v3.0/androidphone/irdeto-control/anonymous-session-token.json?at=ABAAAAAAAAAAAAAAAAAAAAAAgzpRfgenH4iAQ72Sn6EPpM0aUksiLRGIzMzSksra2UI

this works, but needs login:

http://paramountplus.com/apps-api/v3.0/androidphone/irdeto-control/session-token.json?at=ABAAAAAAAAAAAAAAAAAAAAAAgzpRfgenH4iAQ72Sn6EPpM0aUksiLRGIzMzSksra2UI
