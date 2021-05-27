# Tablo for Channels

Exposes a simple M3U playlist for the tuners on a Tablo device, for use in [Channels](https://getchannels.com).

[Channels](https://getchannels.com) supports [custom channels](https://getchannels.com/docs/channels-dvr-server/how-to/custom-channels/) by utilizing streaming sources via M3U playlists.

You will need the private/internal IP of your Tablo device, which you can find via https://api.tablotv.com/assocserver/getipinfo/

## Set Up

Running the container is easy. Fire up the container as usual. You can set which port it runs on. Make sure to set your Tablo device IP.

    docker run -d --restart unless-stopped --name tablo-for-channels -e TABLO_IP=x.x.x.x -p 8050:80 tmm1/tablo-for-channels

You can retrieve the playlist URL via the status page.

    http://127.0.0.1:8050

## Development

```
$ git clone https://github.com/tmm1/tablo-for-channels
$ cd tablo-for-channels
$ go build
```

```
$ ./tablo-for-channels -ip 10.0.1.94
```

```
$ curl localhost:8080/playlist.m3u
#EXTM3U

#EXTINF:-1 channel-id="KEYT-HD" tvg-chno="3.1",KEYT-HD
http://localhost:8080/watch/12

#EXTINF:-1 channel-id="CBS" tvg-chno="3.2",CBS
http://localhost:8080/watch/13

#EXTINF:-1 channel-id="NOW" tvg-chno="3.3",NOW
http://localhost:8080/watch/14
```
