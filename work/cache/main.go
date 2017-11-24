package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/groupcache"
	"net/http"
	"time"
)

// 虚拟文件生成方法
func generateThumbnail(filename string) []byte {
	return []byte("------------this is fake file-------------")
}
func main() {
	me := "http://10.0.0.1"
	peers := groupcache.NewHTTPPool(me)
	//互备node
	peers.Set("http://10.0.0.1", "http://10.0.0.2", "http://10.0.0.3")
	//64M
	var thumbNails = groupcache.NewGroup("thumbnail", 64<<20, groupcache.GetterFunc(func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
		fileName := key
		dest.SetBytes(generateThumbnail(fileName))
		return nil
	}))

	groupcache.RegisterPeerPicker(func() groupcache.PeerPicker {
		return peers
	})

	router := gin.Default()
	router.GET("/files/:name", gin.HandlerFunc(
		func(c *gin.Context) {
			var data []byte
			name := c.Param("name")
			//获取缓存
			err := thumbNails.Get(ctx, name, groupcache.AllocatingByteSliceSink(&data))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "file not found",
				})
				return
			}
			http.ServeContent(c.Writer, c.Request, name, time.Now(), bytes.NewReader(data))
		}))
	router.Run("10.0.0.1:80")
}

/*
type Group struct {
	name       string
	getter     Getter
	perrsOnce  sync.Once
	peers      PeerPicker
	cacheBytes int64
	mainCache  cache
	hotCache   cache
	loadGroup  flightGroup
	Stats      Stats
}
*/

/*Getter
type Getter interface {
	Get(c Context, key string, dest Sink) error
}
*/

/*
type Sink interface {
	SetString(s string) error
	SetBytes(v []byte) error
	SetProto(m proto.Messag) error
}

type ProtoGetter interface {
	Get(c Context, in *pb.GetRequest, out *pb.GetResponse) error
}

type PeerPicker interface {
	PickPeer(key string) (peer ProtoGetter, ok bool)
}
*/
