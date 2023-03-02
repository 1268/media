# Widevine

> Theatricality and deception, powerful agents to the uninitiated.
>
> But we are initiated, aren’t we, Bruce?
>
> The Dark Knight Rises (2012)

First Widevine commit was May 21 2022:

https://github.com/gucio321/mech/commit/9d3dff51

## L1

- Amazon needs AndroidCDM L1 for 720P/1080P/4K
- Disney+ needs Android L1 for 1080P+

## What is a CDM?

The way it works, is you need a key to decrypt the media. To get that key, you
make a request to a license server, and they give you the key back. However the
key returned from the license server, is **itself encrypted**, so before you can
use the key, you have to decrypt it. Thats what the CDM is for. Without the
CDM, you cant decrypt the key, and you cant then use the decrypted key to
decrypt some media. theres **a lot** more detail to it, but thats the high
level view of whats going on.

## Where did proto file come from?

https://github.com/TDenisM/widevinedump/tree/main/pywidevine/cdm/formats

## Where to download L3 CDM?

I cant host those here for legal reasons, but you should be able to download
them from here:

<https://github.com/Jnzzi/4464_L3-CDM>

or search:

https://github.com/search?q=L3+CDM

## How to dump L3 CDM?

Install [Android Studio][1]. Then create Android virtual device:

API Level | ABI | Target
----------|-----|--------------------------
24        | x86 | Android 7.0 (Google APIs)

Then download [Widevine Dumper][2]. Then install:

~~~
pip install -r requirements.txt
~~~

Then download [Frida server][3], example file:

~~~
frida-server-15.1.17-android-x86.xz
~~~

Then start Frida server:

~~~
adb root
adb push frida-server-15.1.17-android-x86 /data/frida-server
adb shell chmod +x /data/frida-server
adb shell /data/frida-server
~~~

Then start Android Chrome and visit [Shaka Player][4]. Click the green play
button. If you receive this prompt:

> bitmovin.com wants to play protected content. Your device’s identity will be
> verified by Google.

Click ALLOW. Then start dumper:

~~~
$env:PROTOCOL_BUFFERS_PYTHON_IMPLEMENTATION = 'python'
python dump_keys.py
~~~

Once you see "Hooks completed", go back to Chrome and click the green play
button again. Result:

~~~
2022-05-21 02:10:52 PM - Helpers.Scanner - 49 - INFO - Key pairs saved at
key_dumps\Android Emulator 5554/private_keys/4464/2770936375
~~~

[1]://developer.android.com/studio
[2]://github.com/wvdumper/dumper
[3]://github.com/frida/frida/releases
[4]://integration.widevine.com/player
