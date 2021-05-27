package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmm1/tablo-for-channels/tablo"
)

var (
	ipFlag = flag.String("ip", "", "ip address of tablo device")
)

func main() {
	flag.Parse()
	if *ipFlag == "" {
		flag.Usage()
		return
	}

	device := tablo.Device{IP: *ipFlag}

	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(200, "%s", INDEX)
	})
	r.GET("/playlist.m3u", func(c *gin.Context) {
		var list []string
		err := device.RequestAPI("GET", "/guide/channels", &list)
		if err != nil {
			log.Printf("[ERR] Failed to fetch channel list: %v", err)
			c.String(503, err.Error())
			return
		}

		c.Header("Content-Type", "application/x-mpegurl")
		fmt.Fprintf(c.Writer, "#EXTM3U\n\n")
		for _, ch := range list {
			var chinfo map[string]interface{}
			err = device.RequestAPI("GET", ch, &chinfo)
			if err != nil {
				log.Printf("[ERR] Failed to fetch info for %v: %v", ch, err)
				c.String(503, err.Error())
				return
			}
			info := chinfo["channel"].(map[string]interface{})
			callsign := info["call_sign"].(string)
			number := fmt.Sprintf("%v.%v", info["major"], info["minor"])
			fmt.Fprintf(c.Writer, "#EXTINF:-1 channel-id=\"%v\" tvg-chno=\"%v\",%s\nhttp://%s/watch/%v\n\n",
				callsign,
				number,
				callsign,
				c.Request.Host,
				chinfo["object_id"],
			)
		}
	})
	r.GET("/watch/:channel", func(c *gin.Context) {
		var resp map[string]interface{}
		device.RequestAPI("POST", "/guide/channels/"+c.Param("channel")+"/watch", &resp)
		c.Redirect(http.StatusTemporaryRedirect, resp["playlist_url"].(string))
	})
	r.Run()
}

const INDEX = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Tablo for Channels</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.1/css/bulma.min.css">
    <style>
      ul{
        margin-bottom: 10px;
      }
    </style>
  </head>
  <body>
  <section class="section">
    <div class="container">
      <h1 class="title">
        Tablo for Channels
      </h1>
      <ul>
        <li><a href="/playlist.m3u">Playlist</a></li>
      </ul>
    </div>
  </section>
  </body>
</html>
`
