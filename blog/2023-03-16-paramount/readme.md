CBS: all videos failing

## Checklist

- [x] I'm reporting a broken site support
- [x] I've verified that I'm running youtube-dl version **2021.12.17**
- [x] I've checked that all provided URLs are alive and playable in a browser
- [x] I've checked that all URLs and arguments with special characters are properly quoted or escaped
- [x] I've searched the bugtracker for similar issues including closed ones

## Verbose log

using an `assetType` of `Downloadable` (not DRM):

https://cbs.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_

~~~
[debug] Command-line args: ['--verbose', 'cbs:YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_']
[debug] Encodings: locale cp1252, fs utf-8, out utf-8, pref cp1252
[debug] youtube-dl version 2021.12.17
[debug] Python version 3.10.4 (CPython) - Windows-10-10.0.18363-SP0
[debug] exe versions: ffmpeg 2022-06-09-git-5d5a014199-essentials_build-www.gyan.dev
[debug] Proxy map: {}
[CBS] YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_: Downloading XML
ERROR: YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_: Failed to parse XML  (caused by
ParseError('no element found: line 1, column 0'))
~~~

same with `assetType` of `StreamPack` (not DRM):

<https://paramountplus.com/movies/video/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_>

~~~
[debug] Command-line args: ['--verbose', 'cbs:wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_']
[debug] Encodings: locale cp1252, fs utf-8, out utf-8, pref cp1252
[debug] youtube-dl version 2021.12.17
[debug] Python version 3.10.4 (CPython) - Windows-10-10.0.18363-SP0
[debug] exe versions: ffmpeg 2022-06-09-git-5d5a014199-essentials_build-www.gyan.dev
[debug] Proxy map: {}
[CBS] wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_: Downloading XML
ERROR: wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_: Failed to parse XML  (caused by
ParseError('no element found: line 1, column 0'))
~~~

## Description

looks like the issue is that these type URLs are now returning empty responses:

<http://can.cbs.com/thunder/player/videoPlayerService.php?partner=cbs&contentId=wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_>

I believe this started in the last day or two, not sure if temporary or permanent issue.
