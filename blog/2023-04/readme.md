# April 2023

OK from my comment here:

https://github.com/ytdl-org/youtube-dl/issues/31864#issuecomment-1474345789

I have this command:

~~~
curl -v link.theplatform.com/s/dJ5BDC/media/guid/2198311517/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_
~~~

which previously returned an HLS URL:

<https://cbsios-vh.akamaihd.net/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2020/05/07/1735196227871/0_0_3436402_ful01_2588_503000.mp4.csmil/master.m3u8>

but now does not:

~~~json
{
   "title": "No AssetType/ProtectionScheme/Format Matches",
   "description": "None of the available releases match the specified AssetType, ProtectionScheme, and/or Format preferences",
   "isException": true,
   "exception": "NoAssetTypeFormatMatches",
   "responseCode": "412"
}
~~~

and this is still a valid ID:

<https://paramountplus.com/movies/video/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_>

change user agent:

~~~
general.useragent.override
~~~

to:

~~~
Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Safari/605.1.15
~~~

https://whatismybrowser.com/guides/the-latest-user-agent/safari
