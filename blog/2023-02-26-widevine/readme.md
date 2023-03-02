# Roku

this episode works:

~~~
> roku -b 12cdc6139bb858f4a597a1aa51a1e37d
GET https://therokuchannel.roku.com/api/v2/homescreen/content/https:%2F%2Fcontent.sr.roku.com%2Fcontent%2Fv1%2Froku-trc%2F12cdc6139bb858f4a597a1aa51a1e37d%3Fexpand=series&include=episodeNumber%252CreleaseDate%252CrunTimeSeconds%252CseasonNumber%252Cseries.seasons.episodes.viewOptions%25E2%2580%2588%252Cseries.title%252Ctitle%252CviewOptions
GET https://therokuchannel.roku.com
POST https://therokuchannel.roku.com/api/v3/playback
GET https://ovp-selector.sr.roku.com/v1/ed8488ba-0a58-43c0-af8b-76a9f1e12593?format=dash
GET https://vod.delivery.roku.com/0a8718f2962d41fda1fd019c5088edbf/3ad88f65a8d5468a97e1967971bb5661/29fa2dea19214f098321a12d80a02477/1ed39e9befba477eb2f83a61e3e92310/index_audio_3_0_init.mp4
POST https://wv.service.expressplay.com/hms/wv/rights/?ExpressPlayToken=BgA1P5QRKb8AJDIzN2U4NTE4LTQwN2QtNDI3Zi05NTkyLWFmMTJiMzRkMmU0NwAAAJAbOlIOTjNmoTRKmdySCF5o-NJMdSvjQE4q2DvfxPselPydqq4L0g7fA9i2uv__h2F0Pj2LdEN2CrxeWZ57och5zIVSUoqeL2aQ-UwlJFyKSdn2AtgCtaEGo-fy3Netp4AFCU6Wyw2S9F-Xl8YQTXSgXVYl5VSy-3wF2NWhv2631vra2_nVF5L-qfw9ltZKbqfSpHdHb5FOwfK6cnO72vODVxtpkA
2.23%   1.05 megabyte   1.03 megabyte/s
~~~

this episode fails:

~~~
> roku -b 105c41ea75775968b670fbb26978ed76
GET https://therokuchannel.roku.com/api/v2/homescreen/content/https:%2F%2Fcontent.sr.roku.com%2Fcontent%2Fv1%2Froku-trc%2F105c41ea75775968b670fbb26978ed76%3Fexpand=series&include=episodeNumber%252CreleaseDate%252CrunTimeSeconds%252CseasonNumber%252Cseries.seasons.episodes.viewOptions%25E2%2580%2588%252Cseries.title%252Ctitle%252CviewOptions
GET https://therokuchannel.roku.com
POST https://therokuchannel.roku.com/api/v3/playback
GET https://ovp-selector.sr.roku.com/v1/3cdd1d9b-47a1-406a-9dda-075d3bbeb22a?format=dash
GET https://vod.delivery.roku.com/165b8483c0804377b1c26f709d70bc2d/14633f105435419fbcfe2821ab237ee5/89d6e79145cb40aa9440af3847affebd/fb61ec652afe4d1495d0870f742f61da/index_audio_3_0_init.mp4
POST https://wv-license.sr.roku.com/license/v1/license/wv?token=Lc0cBq4tz9LIcKdfrviuQx6BBNxUkbxy26VFgSSgHJ4GPTyluaAl_USOEfpfYk19ubSROJHRY8di5i8orG2an6mVJAfAF4lFZCMSVSmmDEKM9HrY2G-mfm3sbX6xIORKllMLb2DHFpJJIhTs4_iTSP5pyktnTOqU0quvOERvpJiioTumJBF78IeKLjdH_CjjAbRf8SZg8MYVv0-DbZlUwTg49xHDx8R0n5laCLEsEnyaXeglLy-ts-gPOzpCtOuPm7RW4izf6AW5JSvmUW_QND3wtkw5JqgC6hL80JOWJ-ji5n8Z3tCZrDQNoUe5hIxvSS4MzimsYQah_cw7A3i6YuJf-36Qd5fYoAnvVmY=&traceId=77b2a82e5ff2520c43e308c098dca728&ExpressPlayToken=none
panic: 500 Internal Server Error
~~~

the ID is valid:

https://therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76

and it plays in a browser. Maybe its the headers? Lets capture a working request,
then switch the headers for the ones I am using. If I switch the query string and
headers for the ones I am using, it still works. So the issue is in the request
body. Does it work with Get Widevine Keys? What is the PSSH?

~~~
AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQvfpNbNs5cC5baB+QYX+afhoKaW50ZXJ0cnVzdCIBKg==
~~~

it works:

~~~
bdfa4d6cdb39702e5b681f90617f9a7e:e258b67d75420066c8424bd142f84565
~~~

could try another CDM. These fail:

- <https://github.com/Jnzzi/4464_L3-CDM>
- https://github.com/srteamer/cdm-l3

could try dumping a new one:

https://github.com/i1v/mech/tree/main/widevine#how-to-dump-l3-cdm

API 24 fails:

~~~
key_dumps\Android Emulator 5554\private_keys\4464\2656745247
~~~

API 25 fails:

~~~
key_dumps\Android Emulator 5554/private_keys/4464/2055674068
~~~

we could try more CDM, but I think at this point its probably Widevine
certificate:

https://github.com/chris124567/hulu/blob/9bcd331a/widevine/pssh.go#L107-L115
