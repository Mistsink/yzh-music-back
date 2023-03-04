package router

import (
	"github.com/Mistsink/kuwo-api/doc"
	"github.com/Mistsink/kuwo-api/global"
	"github.com/Mistsink/kuwo-api/internal/middlerware"
	"github.com/Mistsink/kuwo-api/internal/router/api"
	"github.com/Mistsink/kuwo-api/pkg/app"
	"github.com/Mistsink/kuwo-api/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		//	TODO: create custom logger middleware
		r.Use(gin.Logger())
		r.Use(middlerware.Recovery())
	}

	r.GET("/ping", func(c *gin.Context) {
		response := app.NewResponse(c)
		response.ToResponse(&app.RespJSON{
			Code:   errcode.Success.Code,
			Result: "pong",
		})
	})

	/**
	歌单			playlist
	搜索			search
	歌手			artist
	歌曲相关		 music
	榜单相关 		 ssrank
	*/
	playlist := r.Group("/playlist")
	{
		playlistHandler := api.NewPlaylist()
		playlist.GET("/", playlistHandler.Default) //	default playlist
		playlist.GET("/rec", playlistHandler.Recommend)
		playlist.GET("/:id", playlistHandler.GetDetail)
		playlist.GET("/tag", playlistHandler.WithTag)       //	get playlist in the tag
		playlist.GET("/tag/:id", playlistHandler.WithTag)   //	get playlist in the tag
		playlist.GET("/tags", playlistHandler.GetTags)      //	get tags
		playlist.GET("/album", playlistHandler.AlbumDetail) //	get album info
		playlist.GET("/album/:id", playlistHandler.AlbumDetail)
	}

	search := r.Group("search")
	{
		searchHandler := api.NewSearch()
		search.GET("/", searchHandler.Music) //	default to /music
		search.GET("/hint", searchHandler.Hint)
		search.GET("/music", searchHandler.Music)
		search.GET("/album", searchHandler.Album)
		search.GET("/mv", searchHandler.MV)
		search.GET("/playlist", searchHandler.Playlist)
		search.GET("/artist", searchHandler.Artist)
	}

	artist := r.Group("/artist")
	{
		artistHandler := api.NewArtist()
		artist.GET("/:id", artistHandler.Music)   //	get songs of the artist
		artist.GET("/music", artistHandler.Music) // 	same to /
		artist.GET("/album", artistHandler.Album)
	}

	music := r.Group("/music")
	{
		musicHandler := api.NewMusic()
		music.GET("/", musicHandler.Url)
		music.GET("/:id", musicHandler.Url)
		music.GET("/lyric", musicHandler.Lyric)
		music.GET("/lyric/:id", musicHandler.Lyric)
		music.GET("/url", musicHandler.Url)
		music.GET("/url/:id", musicHandler.Url)
		music.GET("/mv", musicHandler.MV)
		music.GET("/mv/:id", musicHandler.MV)
		music.GET("/info", musicHandler.Info)
		music.GET("/info/:id", musicHandler.Info)
	}

	rank := r.Group("/rank")
	{
		rankHandler := api.NewRank()
		rank.GET("/", rankHandler.Rank) //	get rank list
		rank.GET("/:id", rankHandler.Detail)
	}

	//	No support
	// r.GET("/carousel", api.Carousel)

	r.NoRoute(func(c *gin.Context) {
		response := app.NewResponse(c)
		response.ToErrResponse(errcode.NotFound)
	})

	// api
	{
		docFunc := func(ctx *gin.Context) {
			ctx.Data(200, "html", doc.DocHTML)
		}
		r.GET("/", docFunc)
		r.GET("/api", docFunc)
		r.GET("/doc", docFunc)
		r.GET("/document", docFunc)
	}

	// docStaticFile := "./doc/doc.html"
	// r.StaticFile("/", docStaticFile)
	// r.StaticFile("/doc", docStaticFile)
	// r.StaticFile("/document", docStaticFile)

	return r
}
