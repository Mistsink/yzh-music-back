---
title: kuwo v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.15"

---

# kuwo

> v1.0.0

# search

## GET search-hint

GET /search/hint

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key|query|string| 是 |none|

> 返回示例

> search-hint

```json
{
  "code": 0,
  "result": [
    "RELWORD=稻香\r\nSNUM=274458\r\nRNUM=1000\r\nTYPE=0",
    "RELWORD=稻香  周杰伦\r\nSNUM=321490\r\nRNUM=1000\r\nTYPE=0",
    "RELWORD=稻香十小城夏天\r\nSNUM=124\r\nRNUM=1000\r\nTYPE=0",
    "RELWORD=稻香加小城夏天\r\nSNUM=44\r\nRNUM=1000\r\nTYPE=0",
    "RELWORD=稻香小城夏天\r\nSNUM=2954\r\nRNUM=1000\r\nTYPE=0",
    "RELWORD=稻香女生版\r\nSNUM=27\r\nRNUM=1000\r\nTYPE=0",
    "RELWORD=稻香伴奏\r\nSNUM=25\r\nRNUM=1000\r\nTYPE=0",
    "RELWORD=稻香DJ\r\nSNUM=16\r\nRNUM=1000\r\nTYPE=0",
    "RELWORD=稻香家小城夏天\r\nSNUM=9\r\nRNUM=1000\r\nTYPE=0",
    "RELWORD=稻香时代少年团\r\nSNUM=5\r\nRNUM=1000\r\nTYPE=0"
  ],
  "msg": "success",
  "details": null,
  "req_id": "257d9061105e327f299031b5b95a6da1"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|search-hint|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|[string]|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET artist

GET /search/artist

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key|query|string| 是 |none|
|p|query|string| 是 |页数 default：1|
|n|query|string| 是 |每页数量 default：30|

> 返回示例

> artist

```json
{
  "code": 0,
  "result": {
    "total_cnt": 14,
    "list": [
      {
        "country": "美国",
        "img_url": "https://star.kuwo.cn/star/starheads/120/78/99/1958514511.jpg",
        "music_num": 621,
        "img_120": "https://star.kuwo.cn/star/starheads/120/78/99/1958514511.jpg",
        "name": "Dev",
        "img_70": "https://star.kuwo.cn/star/starheads/70/78/99/1958514511.jpg",
        "id": 10562,
        "img_300": "https://star.kuwo.cn/star/starheads/300/78/99/1958514511.jpg"
      },
      {
        "country": "",
        "img_url": "https://star.kuwo.cn/star/starheads/120/41/61/1568427715.jpg",
        "music_num": 1,
        "img_120": "https://star.kuwo.cn/star/starheads/120/41/61/1568427715.jpg",
        "name": "周",
        "img_70": "https://star.kuwo.cn/star/starheads/70/41/61/1568427715.jpg",
        "id": 9588141,
        "img_300": "https://star.kuwo.cn/star/starheads/300/41/61/1568427715.jpg"
      }
    ]
  },
  "msg": "success",
  "details": null,
  "req_id": "aa204304913c742ae5c01a6f7f927a27"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|artist|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» total_cnt|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» country|string|true|none||none|
|»»» img_url|string|true|none||none|
|»»» music_num|integer|true|none||none|
|»»» img_120|string|true|none||none|
|»»» name|string|true|none||none|
|»»» img_70|string|true|none||none|
|»»» id|integer|true|none||none|
|»»» img_300|string|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET play-list

GET /search/playlist

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key|query|string| 是 |none|
|p|query|string| 是 |页数 default：1|
|n|query|string| 是 |每页数量 default：30|

> 返回示例

> play-list

```json
{
  "code": 0,
  "result": {
    "total_cnt": 150,
    "list": [
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/5/22/1548298834818_350689405_700.jpg",
        "total_cnt": "163",
        "user_name": "十月燕尾",
        "name": "【古风清雅】梅花雨落，棠梨煎雪，梦回稻香",
        "listen_cnt": "51176",
        "id": "2645676459"
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/50/97/1603547388228_528516250_150.jpg",
        "total_cnt": "40",
        "user_name": "沈煜钧",
        "name": "乡村民谣｜随着稻香继续奔跑",
        "listen_cnt": "37504",
        "id": "3095961179"
      }
    ]
  },
  "msg": "success",
  "details": null,
  "req_id": "4d7165033e3c9561622d2ca0654f9ba0"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|play-list|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» total_cnt|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» img_url|string|true|none||none|
|»»» total_cnt|string|true|none||none|
|»»» user_name|string|true|none||none|
|»»» name|string|true|none||none|
|»»» listen_cnt|string|true|none||none|
|»»» id|string|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET mv

GET /search/mv

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key|query|string| 是 |none|
|p|query|string| 是 |页数 default：1|
|n|query|string| 是 |每页数量 default：30|

> 返回示例

> mv

```json
{
  "code": 0,
  "result": {
    "total_cnt": 18,
    "list": [
      {
        "play_cnt": 0,
        "name": "稻香(Live)",
        "id": 196012380,
        "img_url": "https://img1.kuwo.cn/wmvpic/324/88/18/2397149238.jpg",
        "duration": {
          "duration": 229,
          "time_minutes": "03:49"
        },
        "artist": {
          "name": "2021中国好声音学员",
          "id": 8486338
        }
      },
      {
        "play_cnt": 0,
        "name": "稻香(Live)",
        "id": 931249,
        "img_url": "https://img1.kuwo.cn/wmvpic/324/48/12/261044058.jpg",
        "duration": {
          "duration": 222,
          "time_minutes": "03:42"
        },
        "artist": {
          "name": "周杰伦",
          "id": 336
        }
      },
      {
        "play_cnt": 0,
        "name": "稻香(Live)-2014女神的新衣现场",
        "id": 5860042,
        "img_url": "https://img1.kuwo.cn/wmvpic/324/82/32/1799651492.jpg",
        "duration": {
          "duration": 218,
          "time_minutes": "03:38"
        },
        "artist": {
          "name": "周杰伦",
          "id": 336
        }
      }
    ]
  },
  "msg": "success",
  "details": null,
  "req_id": "caac4d2fe77d14491864bd0b67b92a8e"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|mv|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» total_cnt|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» play_cnt|integer|true|none||none|
|»»» name|string|true|none||none|
|»»» id|integer|true|none||none|
|»»» img_url|string|true|none||none|
|»»» duration|object|true|none||none|
|»»»» duration|integer|true|none||none|
|»»»» time_minutes|string|true|none||none|
|»»» artist|object|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET album

GET /search/album

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key|query|string| 是 |none|
|p|query|string| 是 |页数 default：1|
|n|query|string| 是 |每页数量 default：30|

> 返回示例

> album

```json
{
  "code": 0,
  "result": {
    "total_cnt": 124,
    "list": [
      {
        "release_date": "2020-07-01",
        "album": {
          "info": "2000年，邰元龙首次发表词曲创作《有你为伴》。&lt;br&gt;2001年，为中国申奥成功作词作曲《激情澎湃》首次登上中央电视台“音画时尚”。&lt;br&gt;2002年，作词作曲的仙侠歌曲《雪月风花》登上当年最红综艺节目北京电视台“欢乐总动员”。&lt;br&gt;2003年，首次电视歌友会；参与制作抗击SARS非典疫情的公益歌曲《决不放弃》由李慧珍与邰上黄演唱，《光明在心》由央视青歌赛美声冠军得主王莉与邰上黄演唱。&lt;br&gt;2004年，作词作曲《美丽心情》收录在《认识你》群星合辑实体CD，频频登上央视“欢聚一堂”、“激情广场”、“同乐五洲”等节目；首次登上首都体育馆舞台。&lt;br&gt;2005年，作词作曲《寂寞的城市》登上当年最红综艺节目安徽卫视“超级大赢家”；在安徽卫视“超级偶像”新秀电视大赛担任评委（现在的超级网红&nbsp;薇娅Viya&nbsp;就是当年的参赛选手&nbsp;黄薇）。&lt;br&gt;2006年，&lt;br&gt;词曲创作《巨蟹和双鱼》的DEMO录音意外被流出在网路上走红；新浪博客流量过百万；&lt;br&gt;获得中央电视台“中国音乐电视”年度十大新人；&lt;br&gt;演唱了由晚会类作曲家龙寿创作的春晚歌曲《好媳妇》；词曲创作并演唱了晚会作品《梦里桃花源》。&lt;br&gt;2007年，&lt;br&gt;邰元龙的新浪博客获得新浪网络盛典“2006年度音乐博客”“2006年度人气博客”双冠大奖；&lt;br&gt;于香港获得&nbsp;“年度健康艺人”和“亚洲最具实力唱作人”奖项；&lt;br&gt;作词的歌曲《爱的形状》与韩国音乐家合作；《寂寞的城市》抒情摇滚版推出；&lt;br&gt;首次举办个人演唱会，系列名为【情歌盛典】；作词作曲《有史以来》。&lt;br&gt;2008年，&lt;br&gt;与朱茵、陈瑀涵、蔡淳佳、赵守镇等共同成为北京奥运会明星助威团的形象大使；&lt;br&gt;2001年为中国申奥成功作词作曲的《激情澎湃》改动歌词更名为《激情奥运》入选奥运歌曲专辑并拍摄公益MV；首次登上奥运倒计时100天等奥运会舞台；&lt;br&gt;为汶川大地震献唱晚会类作曲家龙寿全新填词独唱版《决不放弃》并拍摄公益MV；&lt;br&gt;在央视某场专为青海省循化地区定制的晚会上将当地撒拉族传统民歌《循化的撒拉》全新改编演唱。&lt;br&gt;2009年，&lt;br&gt;作词作曲的《蓝色爱情海》推出；首次登上央视春晚；注册建立音乐厂牌【ET华娱】。&lt;br&gt;2010年，&lt;br&gt;作词作曲的《只要你的爱》并推出实体CD；&lt;br&gt;将晚会类作曲家龙寿的代表作品《决不放弃》全新填词为禁毒公益歌曲《珍爱生命》并演唱。&lt;br&gt;2011年，&lt;br&gt;以海外华人爱华语的创作视角作词作曲了《华语情歌》备受赞誉；再次登上央视春晚。&lt;br&gt;当选了“第24届世界太太选美大赛北京总决赛”、“第51届国际小姐选美大赛江西总决赛”、“第40届洲际小姐选美大赛中国总决赛”、“第51届国际小姐选美大赛中国总决赛”评委。&lt;br&gt;9月11日，【情歌盛典】（2007年开启的个人演唱会系列）演出总第9场，该系列完美收官。&lt;br&gt;2012年，&lt;br&gt;上半年于苏州同里古镇首次尝试爵士与音乐剧风格表演形式的音乐会。&lt;br&gt;7月19日，邰元龙公开宣布不再以这个身份作为歌手出镜，名字邰元龙全面开始转战幕后，以作词人、作曲人、出品人身份出现；【ET华娱】签约邰上黄。&lt;br&gt;改编2004年的创作《美丽心情》成爵士摇滚版，由邰上黄在其全新系列演唱会“SuperWorld超级世界”首场（山西长治）上演唱。&lt;br&gt;2013年，&lt;br&gt;邰上黄在浙江卫视“转身遇到TA”节目上首唱的爵士新歌《情歌魅影》由邰元龙作词，&lt;br&gt;浙江卫视爆料邰上黄实为《999朵玫瑰》邰正宵的侄子，节目录制现场突发意外状况，女观众上台求爱，该节目成为当年收视率冠军。&lt;br&gt;2016年，&lt;br&gt;以出品人身份重推邰上黄，作词的爵士作品《欢迎光临我的演唱会》进入央视春晚联排后被刷；&lt;br&gt;4月，带团队赴泰国为邰上黄拍摄《蓝色爱情海》、《凡-高》（邰元龙作词作曲；该歌曲暂未推出）两首歌的音乐录影带；&lt;br&gt;8月，联合北京星光大道演播厅，举办【邰上黄SH.T“我是歌手我为歌狂”电视演唱会】全球同步网络直播，收视率1500万人次；&lt;br&gt;9月，为邰上黄制作量身改编成爵士的《月亮代表我的心》登上央视中秋晚会；&lt;br&gt;10月，作词作曲出《赵燕哝》；&lt;br&gt;客串出演古装网络连续剧《奇葩县衙》。&lt;br&gt;2017年，&lt;br&gt;带团队协办“2017世界超级模特大赛中国总决赛”，由邰上黄担任评委，并嘉宾演唱《华语情歌》，拍摄成音乐录影带；邰元龙领衔【ET华娱】团队于11月11日为邰上黄举办了全新系列个人演唱会【歌坛之神】首场首站北京站的演出，票房与声誉均大获成功。&lt;br&gt;词曲作品《有史以来》全网数字发行。&lt;br&gt;2018年，&lt;br&gt;作词作曲的BossaNova单曲《我是歌手太上皇》由邰上黄演唱，于1月湖南卫视“我是歌手”节目开播日上线全网发行；&lt;br&gt;2月，出品了由邰上黄与女星郭柯彤对唱的情歌《爱过你的是我》及其音乐录影带；&lt;br&gt;年末，正式携【ET华娱】团队、旗下艺人及商标“邰上黄”、IP：“歌坛之神”、“阿姆斯特朗”、“歌籁媒”、“蜜汝将”与《赵燕哝》、《华语情歌》等音乐版权加盟【托尼史塔克国际音乐文化公司】。&lt;br&gt;2019年…&lt;br&gt;加&nbsp;T007SH&nbsp;合作",
          "img_url": "https://img1.kuwo.cn/star/albumcover/300/81/69/2439824555.jpg",
          "name": "情歌王子误入晚会圈之《激情奥运》",
          "id": 14647054
        },
        "artist": {
          "name": "邰元龙",
          "id": 3079646
        }
      },
      {
        "release_date": "2020-07-07",
        "album": {
          "info": "邰元龙，中国华语流行音乐领域著名作词人、作曲人、出品人，国际巨星SH.T邰上黄身后的主要音乐创作人兼音乐版权持有人，音乐厂牌【ET华娱】的创始人、主要经营者兼出品人，【ET华娱】持有驰名商标“邰上黄”、“歌坛之神”、“阿姆斯特朗”、“歌籁媒”、“蜜汝将”等IP与音乐版权。邰元龙作词作曲的主要代表作品有《巨蟹和双鱼》、《有史以来》、《华语情歌》、《我是歌手太上皇》、《赵燕哝》等，作词的主要代表作品主要有《爱的形状》、《情歌魅影》、《欢迎光临我的演唱会》、《幸好,有你在（2020非常时期致敬钟南山李文亮及全体逆行者）》等…邰元龙自幼学习中国水墨画出身，亦是当代中国水墨画青年画家，同时酷爱美女类人物摄影，2007年曾以主编身份创立健身健美类实体杂志《围度》，创刊首期封面人物牟丛（中国首位阿诺德赛女子冠军）的封面照由邰元龙拍摄。&lt;br&gt;&lt;br&gt;【概述】&lt;br&gt;2000年，邰元龙首次发表词曲创作《有你为伴》。2001年，为中国申奥成功作词作曲《激情澎湃》。2006年，词曲创作《巨蟹和双鱼》的DEMO录音意外被流出在网路走红。2008年，《激情澎湃》改动歌词更名为《激情奥运》入选奥运歌曲专辑。2011年，以海外华人爱华语的创作视角作词作曲了《华语情歌》备受赞誉。2012年7月19日开始，名字邰元龙主要以作词人、作曲人、出品人身份出现在幕后。2016年10月，作词作曲出《赵燕哝》，该歌曲由邰上黄演唱，于2019年首次登上央视春晚，并在2019年开始逐渐走红。2020年，新型冠状病毒疫情期间作词《幸好,有你在》非常时期致敬钟南山李文亮及全体逆行者…&lt;br&gt;&lt;br&gt;【主要经历】&lt;br&gt;2000年，邰元龙首次发表词曲创作《有你为伴》。&lt;br&gt;2001年，为中国申奥成功作词作曲《激情澎湃》首次登上中央电视台“音画时尚”。&lt;br&gt;2002年，作词作曲的仙侠歌曲《雪月风花》登上当年最红综艺节目北京电视台“欢乐总动员”。&lt;br&gt;2003年，首次电视歌友会；参与制作抗击SARS非典疫情的公益歌曲《决不放弃》由李慧珍与邰上黄演唱，《光明在心》由央视青歌赛美声冠军得主王莉与邰上黄演唱。&lt;br&gt;2004年，&lt;br&gt;作词作曲《美丽心情》收录在《认识你》群星合辑实体CD，频频登上央视“欢聚一堂”、“激情广场”、“同乐五洲”等节目；首次登上首都体育馆舞台。&lt;br&gt;2005年，&lt;br&gt;作词作曲《寂寞的城市》登上当年最红综艺节目安徽卫视“超级大赢家”；&lt;br&gt;在安徽卫视“超级偶像”新秀电视大赛担任评委（现在的超级网红&nbsp;薇娅Viya&nbsp;就是当年的参赛选手&nbsp;黄薇）。&lt;br&gt;2006年，&lt;br&gt;词曲创作《巨蟹和双鱼》的DEMO录音意外被流出在网路上走红；新浪博客流量过百万；&lt;br&gt;获得中央电视台“中国音乐电视”年度十大新人；&lt;br&gt;演唱了由晚会类作曲家龙寿创作的春晚歌曲《好媳妇》；词曲创作并演唱了晚会作品《梦里桃花源》。&lt;br&gt;2007年，&lt;br&gt;邰元龙的新浪博客获得新浪网络盛典“2006年度音乐博客”“2006年度人气博客”双冠大奖；&lt;br&gt;于香港获得&nbsp;“年度健康艺人”和“亚洲最具实力唱作人”奖项；作词的歌曲《爱的形状》与韩国音乐家合作；&lt;br&gt;《寂寞的城市》抒情摇滚版推出；首次举办个人演唱会，系列名为【情歌盛典】；作词作曲《有史以来》。&lt;br&gt;2008年，&lt;br&gt;与朱茵、陈瑀涵、蔡淳佳、赵守镇等共同成为北京奥运会明星助威团的形象大使；&lt;br&gt;2001年为中国申奥成功作词作曲的《激情澎湃》改动歌词更名为《激情奥运》入选奥运歌曲专辑并拍摄公益MV；首次登上奥运倒计时100天等奥运会舞台；&lt;br&gt;为汶川大地震献唱晚会类作曲家龙寿全新填词独唱版《决不放弃》并拍摄公益MV；&lt;br&gt;在央视某场专为青海省循化地区定制的晚会上将当地撒拉族传统民歌《循化的撒拉》全新改编演唱。&lt;br&gt;2009年，&lt;br&gt;作词作曲的《蓝色爱情海》推出；首次登上央视春晚；&lt;br&gt;注册建立音乐厂牌【ET华娱】。&lt;br&gt;2010年，&lt;br&gt;作词作曲的《只要你的爱》并推出实体CD；&lt;br&gt;将晚会类作曲家龙寿的代表作品《决不放弃》全新填词为禁毒公益歌曲《珍爱生命》并演唱。&lt;br&gt;2011年，&lt;br&gt;以海外华人爱华语的创作视角作词作曲了《华语情歌》备受赞誉；再次登上央视春晚。&lt;br&gt;当选了“第24届世界太太选美大赛北京总决赛”、“第51届国际小姐选美大赛江西总决赛”、“第40届洲际小姐选美大赛中国总决赛”、“第51届国际小姐选美大赛中国总决赛”评委。&lt;br&gt;9月11日，【情歌盛典】（2007年开启的个人演唱会系列）演出总第9场，该系列完美收官。&lt;br&gt;2012年，&lt;br&gt;上半年于苏州同里古镇首次尝试爵士与音乐剧风格表演形式的音乐会。&lt;br&gt;7月19日，邰元龙公开宣布不再以这个身份作为歌手出镜，名字邰元龙全面开始转战幕后，以作词人、作曲人、出品人身份出现；【ET华娱】签约邰上黄。&lt;br&gt;改编2004年的创作《美丽心情》成爵士摇滚版，由邰上黄在其全新系列演唱会“SuperWorld超级世界”首场（山西长治）上演唱。&lt;br&gt;2013年，&lt;br&gt;邰上黄在浙江卫视“转身遇到TA”节目上…",
          "img_url": "https://img1.kuwo.cn/star/albumcover/300/70/76/3836428109.jpg",
          "name": "邰元龍版《循化的撒拉》全新电声编曲Kala伴奏Inst.",
          "id": 14733597
        },
        "artist": {
          "name": "邰元龙",
          "id": 3079646
        }
      }
    ]
  },
  "msg": "success",
  "details": null,
  "req_id": "fa2499264be0c0401854f136a7d1442b"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|album|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» total_cnt|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» release_date|string|true|none||none|
|»»» album|object|true|none||none|
|»»»» info|string|true|none||none|
|»»»» img_url|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|»»» artist|object|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET music

GET /search

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|p|query|string| 是 |页数 default：1|
|n|query|string| 是 |每页数量 default：30|
|key|query|string| 是 |none|

> 返回示例

> music

```json
{
  "code": 0,
  "result": {
    "total_cnt": 413,
    "list": [
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/81/79/2995200746.jpg",
        "name": "反方向的钟",
        "id": 215257,
        "score": 82,
        "has_loss_less": true,
        "release_date": "2000-11-07",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/81/79/2995200746.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/81/79/2995200746.jpg",
          "name": "Jay",
          "id": 0
        },
        "artist": {
          "name": "周杰伦",
          "id": 336
        },
        "duration": {
          "duration": 258,
          "time_minutes": "04:18"
        },
        "has_mv": 1,
        "mv_id": 7884573
      },
      {
        "img_url": "https://img2.kuwo.cn/star/starheads/500/12/40/3375344953.jpg",
        "name": "反方向的钟(片段)",
        "id": 161352056,
        "score": 60,
        "has_loss_less": false,
        "release_date": "1970-01-01",
        "img_120": "https://img2.kuwo.cn/star/starheads/120/12/40/3375344953.jpg",
        "album": {
          "img_url": "https://img2.kuwo.cn/star/starheads/500/12/40/3375344953.jpg",
          "name": "",
          "id": 0
        },
        "artist": {
          "name": "旺仔小乔",
          "id": 5237045
        },
        "duration": {
          "duration": 14,
          "time_minutes": "00:14"
        },
        "has_mv": 0,
        "mv_id": 0
      }
    ]
  },
  "msg": "success",
  "details": null,
  "req_id": "7d10a5a54fdd2d001b19e4a2d113f868"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|music|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» total_cnt|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» img_url|string|true|none||none|
|»»» name|string|true|none||none|
|»»» id|integer|true|none||none|
|»»» score|integer|true|none||none|
|»»» has_loss_less|boolean|true|none||none|
|»»» release_date|string|true|none||none|
|»»» img_120|string|true|none||none|
|»»» album|object|true|none||none|
|»»»» img_url|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|»»» artist|object|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|»»» duration|object|true|none||none|
|»»»» duration|integer|true|none||none|
|»»»» time_minutes|string|true|none||none|
|»»» has_mv|integer|true|none||none|
|»»» mv_id|integer|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

# play-list

## GET tmp

GET /

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET category-tag

GET /playlist/tags

> 返回示例

> category-tag

```json
{
  "code": 0,
  "result": [
    {
      "list": [
        {
          "img_url": "https://kwimg2.kuwo.cn/star/upload/73/86/1645691729194_.png",
          "name": "DJ",
          "id": "168"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/3/3/1536566652003_.jpg",
          "name": "抖音",
          "id": "2189"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/6/6/1507883159190_.png",
          "name": "经典",
          "id": "1265"
        },
        {
          "img_url": "https://kwimg3.kuwo.cn/star/upload/44/27/1640834798273_.jpeg",
          "name": "情歌",
          "id": "2200"
        },
        {
          "img_url": "https://kwimg4.kuwo.cn/star/upload/10/62/1640772246574_.jpg",
          "name": "BGM",
          "id": "2199"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/11/11/1517380562443_.png",
          "name": "古风",
          "id": "127"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/5/5/1517380561685_.png",
          "name": "喊麦",
          "id": "216"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/9/9/1517380562105_.png",
          "name": "游戏",
          "id": "1877"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/0/0/1507781397648_.png",
          "name": "轻音乐",
          "id": "173"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/7/7/1517380561751_.png",
          "name": "怀旧",
          "id": "155"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/8/8/1507781342248_.png",
          "name": "佛乐",
          "id": "220"
        },
        {
          "img_url": "https://kwimg1.kuwo.cn/star/upload/38/94/1640834931620_.jpeg",
          "name": "合唱",
          "id": "2201"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/13/13/1507883183437_.png",
          "name": "网络",
          "id": "621"
        },
        {
          "img_url": "https://kwimg3.kuwo.cn/star/upload/55/77/1645694524243_.jpg",
          "name": "儿童",
          "id": "171"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/10/10/1517468038218_.png",
          "name": "ACG",
          "id": "181"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/7/7/1517380562103_.png",
          "name": "影视",
          "id": "180"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/7/7/1517380562103_.png",
          "name": "网红",
          "id": "1879"
        },
        {
          "img_url": "https://kwimg4.kuwo.cn/star/upload/16/73/1551434397397_.jpg",
          "name": "3D",
          "id": "1366"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/14/14/1524130384286_.png",
          "name": "纯音乐",
          "id": "577"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/10/10/1507883200970_.png",
          "name": "KTV",
          "id": "361"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/15/15/1507883255871_.png",
          "name": "器乐",
          "id": "578"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/15/15/1536566688335_.jpg",
          "name": "翻唱",
          "id": "1848"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/15/15/1517380562303_.png",
          "name": "店铺专用",
          "id": "263"
        }
      ],
      "name": "主题",
      "fake_id": "5"
    },
    {
      "list": [
        {
          "img_url": "https://img3.kuwo.cn/star/upload/7/7/1507880388103_.png",
          "name": "伤感",
          "id": "146"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/3/3/1517380562355_.png",
          "name": "放松",
          "id": "62"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/13/13/1507881732397_.png",
          "name": "励志",
          "id": "58"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/1/1/1507881938529_.png",
          "name": "开心",
          "id": "143"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/13/13/1507881681997_.png",
          "name": "甜蜜",
          "id": "137"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/13/13/1517380561997_.png",
          "name": "兴奋",
          "id": "139"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/9/9/1507880598777_.png",
          "name": "安静",
          "id": "67"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/7/7/1517380562087_.png",
          "name": "治愈",
          "id": "66"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/6/6/1507880502006_.png",
          "name": "寂寞",
          "id": "147"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/9/9/1507880572457_.png",
          "name": "思念",
          "id": "160"
        }
      ],
      "name": "心情",
      "fake_id": "1"
    },
    {
      "list": [
        {
          "img_url": "https://img3.kuwo.cn/star/upload/4/4/1507882143860_.png",
          "name": "开车",
          "id": "376"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/13/13/1507882327133_.png",
          "name": "运动",
          "id": "366"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/10/10/1507882105930_.png",
          "name": "睡前",
          "id": "354"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/13/13/1507882306893_.png",
          "name": "跳舞",
          "id": "378"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/11/11/1507882229403_.png",
          "name": "学习",
          "id": "1876"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/1/1/1507882054209_.png",
          "name": "清晨",
          "id": "353"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/1/1/1507882289153_.png",
          "name": "夜店",
          "id": "359"
        },
        {
          "img_url": "https://img3.kuwo.cn/star/upload/6/6/1517380561894_.png",
          "name": "校园",
          "id": "382"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/8/8/1507882121720_.png",
          "name": "咖啡店",
          "id": "363"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/11/11/1507882267643_.png",
          "name": "旅行",
          "id": "375"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/12/12/1507882205772_.png",
          "name": "工作",
          "id": "386"
        },
        {
          "img_url": "https://kwimg3.kuwo.cn/star/upload/24/13/1611134337903_.jpg",
          "name": "广场舞",
          "id": "334"
        }
      ],
      "name": "场景",
      "fake_id": "3"
    },
    {
      "list": [
        {
          "img_url": "https://img2.kuwo.cn/star/upload/6/6/1507882866278_.png",
          "name": "70后",
          "id": "637"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/9/9/1507882844009_.png",
          "name": "80后",
          "id": "638"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/6/6/1507882820598_.png",
          "name": "90后",
          "id": "639"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/15/15/1507883023759_.png",
          "name": "00后",
          "id": "640"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/8/8/1517380562168_.png",
          "name": "10后",
          "id": "268"
        }
      ],
      "name": "年代",
      "fake_id": "143"
    },
    {
      "list": [
        {
          "img_url": "https://img4.kuwo.cn/star/upload/11/11/1507866722059_.png",
          "name": "流行",
          "id": "393"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/9/9/1507866744649_.png",
          "name": "电子",
          "id": "391"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/11/11/1507866778379_.png",
          "name": "摇滚",
          "id": "389"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/9/9/1517468516601_.png",
          "name": "民歌",
          "id": "1921"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/11/11/1507866837499_.png",
          "name": "民谣",
          "id": "392"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/12/12/1513764470636_.png",
          "name": "古典",
          "id": "390"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/8/8/1517468234136_.png",
          "name": "嘻哈",
          "id": "387"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/12/12/1507867025660_.png",
          "name": "乡村",
          "id": "399"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/6/6/1507866797318_.png",
          "name": "爵士",
          "id": "397"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/11/11/1507866815739_.png",
          "name": "R&B",
          "id": "394"
        }
      ],
      "name": "曲风流派",
      "fake_id": "8"
    },
    {
      "list": [
        {
          "img_url": "https://img3.kuwo.cn/star/upload/3/3/1507866482035_.png",
          "name": "华语",
          "id": "37"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/9/9/1507866499465_.png",
          "name": "欧美",
          "id": "35"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/5/5/1507886214677_.png",
          "name": "韩语",
          "id": "1093"
        },
        {
          "img_url": "https://img1.kuwo.cn/star/upload/14/14/1507866553486_.png",
          "name": "粤语",
          "id": "13"
        },
        {
          "img_url": "https://img2.kuwo.cn/star/upload/9/9/1507866530505_.png",
          "name": "日语",
          "id": "1091"
        },
        {
          "img_url": "https://img4.kuwo.cn/star/upload/5/5/1507883589845_.png",
          "name": "小语种",
          "id": "12"
        }
      ],
      "name": "语言",
      "fake_id": "6"
    },
    {
      "list": [],
      "name": "有声系列",
      "fake_id": "103"
    }
  ],
  "msg": "success",
  "details": null,
  "req_id": "9ad753387619d1427adae7f23ca9f0ef"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|category-tag|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|[object]|true|none||none|
|»» list|[object]|true|none||none|
|»»» img_url|string|true|none||none|
|»»» name|string|true|none||none|
|»»» id|string|true|none||none|
|»» name|string|true|none||none|
|»» fake_id|string|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET play-list-by-tag

GET /playlist/tag/2190

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|p|query|string| 是 |none|
|n|query|string| 是 |none|

> 返回示例

> play-list-by-tag

```json
{
  "code": 0,
  "result": {
    "total_cnt": 111,
    "list": [
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/5/19/1643169744591_517355605_500.jpg",
        "user_name": "莹火虫",
        "total_cnt": 26,
        "name": "恭喜发财，今年胜旧年",
        "listen_cnt": 57776,
        "id": 3250004788
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/81/96/1614388299650_515138181_500.jpg",
        "user_name": "迪莉_",
        "total_cnt": 16,
        "name": "〖元宵快乐〗新的开始，唯独喜欢依旧",
        "listen_cnt": 18024,
        "id": 3158153223
      }
    ],
    "n": 2,
    "p": 1
  },
  "msg": "success",
  "details": null,
  "req_id": "2ff578512224e48d38046e7047be77ac"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|play-list-by-tag|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» total_cnt|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» img_url|string|true|none||none|
|»»» user_name|string|true|none||none|
|»»» total_cnt|integer|true|none||none|
|»»» name|string|true|none||none|
|»»» listen_cnt|integer|true|none||none|
|»»» id|integer|true|none||none|
|»» n|integer|true|none||none|
|»» p|integer|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET recommand-play-list

GET /playlist/rec

> 返回示例

> recommand-play-list

```json
{
  "code": 0,
  "result": {
    "list": [
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/10/13/1649904612537_132026710_500.jpg",
        "user_name": "酷小编",
        "img_700": "https://img1.kuwo.cn/star/userpl2015/10/13/1649904612537_132026710_700.jpg",
        "img_300": "https://img1.kuwo.cn/star/userpl2015/10/13/1649904612537_132026710b.jpg",
        "img_500": "https://img1.kuwo.cn/star/userpl2015/10/13/1649904612537_132026710_500.jpg",
        "total_cnt": 1699,
        "name": "每日最新单曲推荐",
        "listen_cnt": 203475446,
        "id": 1082685104
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/62/77/1543559517343_368142962_500.jpg",
        "user_name": "张苏苏",
        "img_700": "https://img1.kuwo.cn/star/userpl2015/62/77/1543559517343_368142962_700.jpg",
        "img_300": "https://img1.kuwo.cn/star/userpl2015/62/77/1543559517343_368142962b.jpg",
        "img_500": "https://img1.kuwo.cn/star/userpl2015/62/77/1543559517343_368142962_500.jpg",
        "total_cnt": 32,
        "name": "【搞笑神曲】伙伴们注意了！这里有神曲出没",
        "listen_cnt": 1112660,
        "id": 2613400951
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/18/46/1585664790056_366896818_500.jpg",
        "user_name": "若水",
        "img_700": "https://img1.kuwo.cn/star/userpl2015/18/46/1585664790056_366896818_700.jpg",
        "img_300": "https://img1.kuwo.cn/star/userpl2015/18/46/1585664790056_366896818b.jpg",
        "img_500": "https://img1.kuwo.cn/star/userpl2015/18/46/1585664790056_366896818_500.jpg",
        "total_cnt": 20,
        "name": "古风戏腔|一如蝶衣，入戏太深，至死不懂出戏。",
        "listen_cnt": 973799,
        "id": 2986313346
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/58/10/1589252088713_415617058_500.jpg",
        "user_name": "王2014",
        "img_700": "https://img1.kuwo.cn/star/userpl2015/58/10/1589252088713_415617058_700.jpg",
        "img_300": "https://img1.kuwo.cn/star/userpl2015/58/10/1589252088713_415617058b.jpg",
        "img_500": "https://img1.kuwo.cn/star/userpl2015/58/10/1589252088713_415617058_500.jpg",
        "total_cnt": 73,
        "name": "【半吨兄弟·翻唱】逃不掉、忘不了的烟嗓味",
        "listen_cnt": 281002,
        "id": 3010694594
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/16/30/1564556379974_183963216_500.jpg",
        "user_name": "小精灵",
        "img_700": "https://img1.kuwo.cn/star/userpl2015/16/30/1564556379974_183963216_700.jpg",
        "img_300": "https://img1.kuwo.cn/star/userpl2015/16/30/1564556379974_183963216b.jpg",
        "img_500": "https://img1.kuwo.cn/star/userpl2015/16/30/1564556379974_183963216_500.jpg",
        "total_cnt": 30,
        "name": "快乐的童年，经典儿歌陪伴你",
        "listen_cnt": 1680656,
        "id": 2832261176
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/96/96/1593678940964_516582096_500.jpg",
        "user_name": "猫小酱",
        "img_700": "https://img1.kuwo.cn/star/userpl2015/96/96/1593678940964_516582096_700.jpg",
        "img_300": "https://img1.kuwo.cn/star/userpl2015/96/96/1593678940964_516582096b.jpg",
        "img_500": "https://img1.kuwo.cn/star/userpl2015/96/96/1593678940964_516582096_500.jpg",
        "total_cnt": 15,
        "name": "[伤感]我们分开了，从此一个人",
        "listen_cnt": 843546,
        "id": 3038180072
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/1/0/1609887644150_494940501_500.jpg",
        "user_name": "知返",
        "img_700": "https://img1.kuwo.cn/star/userpl2015/1/0/1609887644150_494940501_700.jpg",
        "img_300": "https://img1.kuwo.cn/star/userpl2015/1/0/1609887644150_494940501b.jpg",
        "img_500": "https://img1.kuwo.cn/star/userpl2015/1/0/1609887644150_494940501_500.jpg",
        "total_cnt": 176,
        "name": "热歌来袭丨潮流永不断",
        "listen_cnt": 1307635,
        "id": 3131127980
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/20/77/1523629510068_249594620_500.jpg",
        "user_name": "小舞",
        "img_700": "https://img1.kuwo.cn/star/userpl2015/20/77/1523629510068_249594620_700.jpg",
        "img_300": "https://img1.kuwo.cn/star/userpl2015/20/77/1523629510068_249594620b.jpg",
        "img_500": "https://img1.kuwo.cn/star/userpl2015/20/77/1523629510068_249594620_500.jpg",
        "total_cnt": 18,
        "name": "经典儿歌丨送给小朋友们的歌",
        "listen_cnt": 4621595,
        "id": 2449220152
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/5/92/1580873213226_449973605_500.jpg",
        "user_name": "天助",
        "img_700": "https://img1.kuwo.cn/star/userpl2015/5/92/1580873213226_449973605_700.jpg",
        "img_300": "https://img1.kuwo.cn/star/userpl2015/5/92/1580873213226_449973605b.jpg",
        "img_500": "https://img1.kuwo.cn/star/userpl2015/5/92/1580873213226_449973605_500.jpg",
        "total_cnt": 36,
        "name": "青春之歌丨可惜她己远去，消失在人海",
        "listen_cnt": 824443,
        "id": 2952682794
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/75/24/1551966928274_444980175_500.jpg",
        "user_name": "衾琳儿",
        "img_700": "https://img1.kuwo.cn/star/userpl2015/75/24/1551966928274_444980175_700.jpg",
        "img_300": "https://img1.kuwo.cn/star/userpl2015/75/24/1551966928274_444980175b.jpg",
        "img_500": "https://img1.kuwo.cn/star/userpl2015/75/24/1551966928274_444980175_500.jpg",
        "total_cnt": 20,
        "name": "『静听轻音』信步一曲春の私语",
        "listen_cnt": 120937,
        "id": 2703195565
      },
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/42/7/1602067131573_515737042_500.jpg",
        "user_name": "若若睏了",
        "img_700": "https://img1.kuwo.cn/star/userpl2015/42/7/1602067131573_515737042_700.jpg",
        "img_300": "https://img1.kuwo.cn/star/userpl2015/42/7/1602067131573_515737042b.jpg",
        "img_500": "https://img1.kuwo.cn/star/userpl2015/42/7/1602067131573_515737042_500.jpg",
        "total_cnt": 21,
        "name": "『伤感』我忘了回忆，你忘了忘记",
        "listen_cnt": 1181038,
        "id": 3088280577
      }
    ]
  },
  "msg": "success",
  "details": null,
  "req_id": "8aa8a48e5078c655b541a4170d4400a5"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|recommand-play-list|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» list|[object]|true|none||none|
|»»» img_url|string|true|none||none|
|»»» user_name|string|true|none||none|
|»»» img_700|string|true|none||none|
|»»» img_300|string|true|none||none|
|»»» img_500|string|true|none||none|
|»»» total_cnt|integer|true|none||none|
|»»» name|string|true|none||none|
|»»» listen_cnt|integer|true|none||none|
|»»» id|integer|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET default-play-list

GET /playlist

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|order|query|string| 是 |new 最新 hot 最热		默认 new|
|n|query|string| 是 |每页条数		默认 30|
|p|query|string| 是 |页数		默认 1|

> 返回示例

> default-play-list

```json
{
  "code": 0,
  "result": {
    "total_cnt": 9480,
    "data": [
      {
        "img_url": "https://img1.kuwo.cn/star/userpl2015/34/85/1651234597871_448366234_500.jpg",
        "user_name": "扶笙",
        "total_cnt": 58,
        "name": "学生党必备丨总结全科歌曲，逆袭成功",
        "listen_cnt": 23239,
        "id": 3346425332
      }
    ],
    "n": 1,
    "p": 2
  },
  "msg": "success",
  "details": null,
  "req_id": "4be5e72bcba2b78f6c230d9f520b29f9"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|default-play-list|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» total_cnt|integer|true|none||none|
|»» data|[object]|true|none||none|
|»»» img_url|string|false|none||none|
|»»» user_name|string|false|none||none|
|»»» total_cnt|integer|false|none||none|
|»»» name|string|false|none||none|
|»»» listen_cnt|integer|false|none||none|
|»»» id|integer|false|none||none|
|»» n|integer|true|none||none|
|»» p|integer|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET play-list-music

GET /playlist/3250004788

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|p|query|string| 是 |none|
|n|query|string| 是 |none|

> 返回示例

> play-list-music

```json
{
  "code": 0,
  "result": {
    "img_url": "https://img1.kuwo.cn/star/userpl2015/5/19/1643169744591_517355605_150.jpg",
    "user_img": "https://img1.kuwo.cn/star/userhead/5/19/1621528674963_517355605.jpg",
    "user_name": "莹火虫",
    "img_700": "https://img1.kuwo.cn/star/userpl2015/5/19/1643169744591_517355605_700.jpg",
    "img_300": "https://img1.kuwo.cn/star/userpl2015/5/19/1643169744591_517355605b.jpg",
    "img_500": "https://img1.kuwo.cn/star/userpl2015/5/19/1643169744591_517355605_500.jpg",
    "total_cnt": 26,
    "name": "恭喜发财，今年胜旧年",
    "listen_cnt": 57776,
    "id": 3250004788,
    "tag": "粤语,兴奋,开心",
    "list": [
      {
        "img_url": "https://img1.kuwo.cn/star/albumcover/500/82/55/4238176893.jpg",
        "name": "祝福你",
        "id": 40710748,
        "score": 41,
        "has_loss_less": true,
        "release_date": "2014-01-28",
        "img_120": "https://img1.kuwo.cn/star/albumcover/120/82/55/4238176893.jpg",
        "album": {
          "img_url": "https://img1.kuwo.cn/star/albumcover/500/82/55/4238176893.jpg",
          "name": "今年会更好",
          "id": 0
        },
        "artist": {
          "name": "华语群星",
          "id": 11068
        },
        "duration": {
          "duration": 198,
          "time_minutes": "03:18"
        },
        "has_mv": 1,
        "mv_id": 14559502
      },
      {
        "img_url": "https://img4.kuwo.cn/star/starheads/500/41/73/4157743756.jpg",
        "name": "祝新岁",
        "id": 536832,
        "score": 30,
        "has_loss_less": false,
        "release_date": "1970-01-01",
        "img_120": "https://img4.kuwo.cn/star/starheads/120/41/73/4157743756.jpg",
        "album": {
          "img_url": "https://img4.kuwo.cn/star/starheads/500/41/73/4157743756.jpg",
          "name": "",
          "id": 0
        },
        "artist": {
          "name": "群星",
          "id": 11135
        },
        "duration": {
          "duration": 121,
          "time_minutes": "02:01"
        },
        "has_mv": 1,
        "mv_id": 13873184
      },
      {
        "img_url": "https://img1.kuwo.cn/star/albumcover/500/64/87/3196871528.jpg",
        "name": "欢乐年年",
        "id": 40973561,
        "score": 28,
        "has_loss_less": false,
        "release_date": "1994-01-01",
        "img_120": "https://img1.kuwo.cn/star/albumcover/120/64/87/3196871528.jpg",
        "album": {
          "img_url": "https://img1.kuwo.cn/star/albumcover/500/64/87/3196871528.jpg",
          "name": "华纳群星贺岁金曲 喜上加喜+天天报喜",
          "id": 0
        },
        "artist": {
          "name": "杜德伟&叶倩文",
          "id": 713
        },
        "duration": {
          "duration": 161,
          "time_minutes": "02:41"
        },
        "has_mv": 1,
        "mv_id": 13458689
      }
    ]
  },
  "msg": "success",
  "details": null,
  "req_id": "c4ae7b0131db087d3385d218f1f578ae"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|play-list-music|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» img_url|string|true|none||none|
|»» user_img|string|true|none||none|
|»» user_name|string|true|none||none|
|»» img_700|string|true|none||none|
|»» img_300|string|true|none||none|
|»» img_500|string|true|none||none|
|»» total_cnt|integer|true|none||none|
|»» name|string|true|none||none|
|»» listen_cnt|integer|true|none||none|
|»» id|integer|true|none||none|
|»» tag|string|true|none||none|
|»» list|[object]|true|none||none|
|»»» img_url|string|true|none||none|
|»»» name|string|true|none||none|
|»»» id|integer|true|none||none|
|»»» score|integer|true|none||none|
|»»» has_loss_less|boolean|true|none||none|
|»»» release_date|string|true|none||none|
|»»» img_120|string|true|none||none|
|»»» album|object|true|none||none|
|»»»» img_url|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|»»» artist|object|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|»»» duration|object|true|none||none|
|»»»» duration|integer|true|none||none|
|»»»» time_minutes|string|true|none||none|
|»»» has_mv|integer|true|none||none|
|»»» mv_id|integer|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

# singer

## GET recommand-singer

GET /rec_singer

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|rn|query|string| 是 |每页数量		默认6|
|pn|query|string| 是 |页数		默认1|
|category|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET all-singer-by-tag

GET /singer

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|category|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET song-with-singer

GET /artist/music

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|
|p|query|string| 是 |none|
|n|query|string| 是 |none|

> 返回示例

> song-with-singer

```json
{
  "code": 0,
  "result": {
    "total_cnt": 675,
    "list": [
      {
        "img_url": "https://img4.kuwo.cn/star/albumcover/120/20/5/656820488.jpg",
        "name": "喜欢你",
        "id": 5245130,
        "score": 87,
        "has_loss_less": true,
        "release_date": "2014-08-15",
        "img_120": "https://img4.kuwo.cn/star/albumcover/120/20/5/656820488.jpg",
        "album": {
          "img_url": "https://img4.kuwo.cn/star/albumcover/120/20/5/656820488.jpg",
          "name": "喜欢你",
          "id": 0
        },
        "artist": {
          "name": "G.E.M.邓紫棋",
          "id": 5371
        },
        "duration": {
          "duration": 239,
          "time_minutes": "03:59"
        },
        "has_mv": 0,
        "mv_id": 12517527
      },
      {
        "img_url": "https://img4.kuwo.cn/star/albumcover/120/23/91/496164472.jpg",
        "name": "光年之外",
        "id": 15195332,
        "score": 88,
        "has_loss_less": true,
        "release_date": "2016-12-29",
        "img_120": "https://img4.kuwo.cn/star/albumcover/120/23/91/496164472.jpg",
        "album": {
          "img_url": "https://img4.kuwo.cn/star/albumcover/120/23/91/496164472.jpg",
          "name": "光年之外",
          "id": 0
        },
        "artist": {
          "name": "G.E.M.邓紫棋",
          "id": 5371
        },
        "duration": {
          "duration": 235,
          "time_minutes": "03:55"
        },
        "has_mv": 0,
        "mv_id": 13498159
      }
    ]
  },
  "msg": "success",
  "details": null,
  "req_id": "6a139fe84f9d83c7a3919588cfd7c55d"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|song-with-singer|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» total_cnt|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» img_url|string|true|none||none|
|»»» name|string|true|none||none|
|»»» id|integer|true|none||none|
|»»» score|integer|true|none||none|
|»»» has_loss_less|boolean|true|none||none|
|»»» release_date|string|true|none||none|
|»»» img_120|string|true|none||none|
|»»» album|object|true|none||none|
|»»»» img_url|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|»»» artist|object|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|»»» duration|object|true|none||none|
|»»»» duration|integer|true|none||none|
|»»»» time_minutes|string|true|none||none|
|»»» has_mv|integer|true|none||none|
|»»» mv_id|integer|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET album-with-singer

GET /artist/album

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |歌手id	required|
|n|query|string| 是 |每页数量		默认30|
|p|query|string| 是 |分页		默认 1|

> 返回示例

> album-with-singer

```json
{
  "code": 0,
  "result": {
    "total_cnt": 63,
    "list": [
      {
        "release_date": "2022-07-01",
        "album": {
          "info": "",
          "img_url": "https://img1.kuwo.cn/star/albumcover/300/8/70/2980150646.jpg",
          "name": "Bang&nbsp;Bang&nbsp;(From&nbsp;&apos;Minions:&nbsp;The&nbsp;Rise&nbsp;of&nbsp;Gru&apos;&nbsp;Soundtrack)",
          "id": 29593617
        },
        "artist": {
          "name": "G.E.M.邓紫棋",
          "id": 5371
        }
      },
      {
        "release_date": "2022-07-01",
        "album": {
          "info": "",
          "img_url": "https://img1.kuwo.cn/star/albumcover/300/8/70/2980150646.jpg",
          "name": "小黄人大眼萌：神偷奶爸前传&nbsp;电影原声专辑",
          "id": 29583970
        },
        "artist": {
          "name": "小黄人The&nbsp;Minions&王嘉尔&G.E.M.邓紫棋&St.&nbsp;Vincent&Diana&nbsp;Ross&Tame&nbsp;Impala&Phoebe&nbsp;Bridgers",
          "id": 254725
        }
      },
      {
        "release_date": "2021-12-21",
        "album": {
          "info": "如果时间可以倒流，你想回到谁的身边？&lt;br&gt;&lt;br&gt;G.E.M.邓紫棋再度参与作曲并担任填词，亲自制作年末最动人情歌《倒流时间》，并惊喜推出歌词版MV。细腻且多层次的嗓音，令人一秒开启回忆模式。&lt;br&gt;&nbsp;&lt;br&gt;已经进入2021年尾声，街道上充斥着节日的气息&lt;br&gt;想好今年的跨年计划了吗？&lt;br&gt;&nbsp;&lt;br&gt;此刻的你，是否也一样&lt;br&gt;望着窗外不属于自己的花天锦地，热烈繁华&lt;br&gt;不禁想起了过往，想起了有Ta陪伴的点点滴滴&lt;br&gt;&nbsp;&lt;br&gt;人生短暂，想倒流时间回到起点，从头再爱身边的人一遍&lt;br&gt;想倒流时间回到那个夏天，你送的花也曾开得漂亮&lt;br&gt;可如今，花瓣慢慢枯萎，就像我们在一起的时光，被你慢慢淡忘&lt;br&gt;我从回忆里走来，才发现，“原来世界最远，不是距离而是昨天”&lt;br&gt;&nbsp;&lt;br&gt;2021年的G.E.M.邓紫棋是充实且成长的，从写《孤独》时的迷茫感觉中走了出来，更加坚定自己存在的意义和目的。这些年来的经历使得她在创作歌曲时，如有神助，灵感爆棚。“人生难免会有后悔、遗憾的时候，只是我们既然尽力了，那就只能接受遗憾。”新的一年，好好生活，永远充满爱与希望！&lt;br&gt;&nbsp;&lt;br&gt;2022，一起期待净化了，且进化了的自己",
          "img_url": "https://img1.kuwo.cn/star/albumcover/300/56/64/934594559.jpg",
          "name": "倒流时间",
          "id": 25095736
        },
        "artist": {
          "name": "G.E.M.邓紫棋",
          "id": 5371
        }
      }
    ]
  },
  "msg": "success",
  "details": null,
  "req_id": "b03d309efde1897f91c394e32c9547e5"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|album-with-singer|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» total_cnt|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» release_date|string|true|none||none|
|»»» album|object|true|none||none|
|»»»» info|string|true|none||none|
|»»»» img_url|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|»»» artist|object|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

# music

## GET lyric

GET /music/lyric/162457325

> 返回示例

> lyric

```json
{
  "code": 0,
  "result": {
    "lyric_list": [
      {
        "lrc_line": "等你归来 - 程响",
        "lrc_time": "0.0"
      },
      {
        "lrc_line": "词：宁缺",
        "lrc_time": "0.99"
      },
      {
        "lrc_line": "曲：宁缺",
        "lrc_time": "1.4"
      },
      {
        "lrc_line": "编曲：曲比阿且",
        "lrc_time": "1.84"
      },
      {
        "lrc_line": "宣发：刘芳",
        "lrc_time": "2.7"
      },
      {
        "lrc_line": "混音：王朋",
        "lrc_time": "3.26"
      },
      {
        "lrc_line": "母带：全相彦",
        "lrc_time": "3.86"
      },
      {
        "lrc_line": "宣发：刘芳",
        "lrc_time": "4.6"
      },
      {
        "lrc_line": "文案：冷小婧",
        "lrc_time": "5.2"
      },
      {
        "lrc_line": "艺术效果：华玮轩",
        "lrc_time": "5.95"
      },
      {
        "lrc_line": "营销推广：银河方舟X壳壳菠萝",
        "lrc_time": "7.02"
      },
      {
        "lrc_line": "项目总监：庄有豪",
        "lrc_time": "9.05"
      },
      {
        "lrc_line": "出品人：潘鸿业",
        "lrc_time": "10.16"
      },
      {
        "lrc_line": "OP：北京金翼龙国际文化传媒有限公司",
        "lrc_time": "11.12"
      },
      {
        "lrc_line": "我就在这里等你披星戴月",
        "lrc_time": "20.71"
      },
      {
        "lrc_line": "乘着风而来",
        "lrc_time": "26.02"
      },
      {
        "lrc_line": "我就在这里埋好烈酒",
        "lrc_time": "30.1"
      },
      {
        "lrc_line": "候你故事开",
        "lrc_time": "34.35"
      },
      {
        "lrc_line": "千千万万人海灯火阑珊",
        "lrc_time": "37.38"
      },
      {
        "lrc_line": "你多少次不在",
        "lrc_time": "42.24"
      },
      {
        "lrc_line": "走遍高高低低一路辗转",
        "lrc_time": "45.96"
      },
      {
        "lrc_line": "朝暮青丝已白",
        "lrc_time": "50.81"
      },
      {
        "lrc_line": "我在红尘等你 人间等你",
        "lrc_time": "56.59"
      },
      {
        "lrc_line": "守繁华之外",
        "lrc_time": "61.52"
      },
      {
        "lrc_line": "揽尽星辰入怀 千川归来",
        "lrc_time": "64.71"
      },
      {
        "lrc_line": "化一片沧海",
        "lrc_time": "69.94"
      },
      {
        "lrc_line": "我在九幽等你 极乐等你",
        "lrc_time": "73.2"
      },
      {
        "lrc_line": "望彼岸花开",
        "lrc_time": "78.270004"
      },
      {
        "lrc_line": "长对三生浮白 不畏不改",
        "lrc_time": "81.42"
      },
      {
        "lrc_line": "渡过去将来",
        "lrc_time": "86.62"
      },
      {
        "lrc_line": "我就在这里等你跨山越海",
        "lrc_time": "110.8"
      },
      {
        "lrc_line": "踏着云烟来",
        "lrc_time": "115.990005"
      },
      {
        "lrc_line": "我就在这里望尽天涯",
        "lrc_time": "120.36"
      },
      {
        "lrc_line": "风雨也不改",
        "lrc_time": "124.36"
      },
      {
        "lrc_line": "安安静静岁月时光荏苒",
        "lrc_time": "127.34"
      },
      {
        "lrc_line": "你或许会徘徊",
        "lrc_time": "132.43"
      },
      {
        "lrc_line": "挥别近近远远一身尘埃",
        "lrc_time": "135.86"
      },
      {
        "lrc_line": "俯仰皆是无奈",
        "lrc_time": "140.81"
      },
      {
        "lrc_line": "我在红尘等你 人间等你",
        "lrc_time": "146.49"
      },
      {
        "lrc_line": "守繁华之外",
        "lrc_time": "151.48"
      },
      {
        "lrc_line": "揽尽星辰入怀 千川归来",
        "lrc_time": "154.70999"
      },
      {
        "lrc_line": "化一片沧海",
        "lrc_time": "160.0"
      },
      {
        "lrc_line": "我在九幽等你 极乐等你",
        "lrc_time": "163.27"
      },
      {
        "lrc_line": "望彼岸花开",
        "lrc_time": "168.25"
      },
      {
        "lrc_line": "长对三生浮白 不畏不改",
        "lrc_time": "171.4"
      },
      {
        "lrc_line": "渡过去将来",
        "lrc_time": "176.77"
      },
      {
        "lrc_line": "我在红尘等你 人间等你",
        "lrc_time": "179.73"
      },
      {
        "lrc_line": "守繁华之外",
        "lrc_time": "185.05"
      },
      {
        "lrc_line": "揽尽星辰入怀 千川归来",
        "lrc_time": "188.07"
      },
      {
        "lrc_line": "化一片沧海",
        "lrc_time": "193.47"
      },
      {
        "lrc_line": "我在九幽等你 极乐等你",
        "lrc_time": "196.65"
      },
      {
        "lrc_line": "望彼岸花开",
        "lrc_time": "201.84"
      },
      {
        "lrc_line": "长对三生浮白 不畏不改",
        "lrc_time": "204.88"
      },
      {
        "lrc_line": "渡过去将来",
        "lrc_time": "210.24"
      },
      {
        "lrc_line": "长对三生浮白 不畏不改",
        "lrc_time": "215.54001"
      },
      {
        "lrc_line": "渡过去将来",
        "lrc_time": "220.79001"
      }
    ]
  },
  "msg": "成功",
  "details": null,
  "req_id": "b3105d1cX8863X409dX816bX8c7c7bd8977c"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|lyric|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» lyric_list|[object]|true|none||none|
|»»» lrc_line|string|true|none||none|
|»»» lrc_time|string|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET song-url

GET /music/162457325

> 返回示例

> song-url

```json
{
  "code": 0,
  "result": {
    "url": "https://ci-sycdn.kuwo.cn/d34f6910e61061cfb2a38c5b0fdb8e9f/62e746b6/resource/n2/67/80/2690011765.mp3"
  },
  "msg": "success",
  "details": null,
  "req_id": "a2fd70a12121f4e076a0b9c86f40ae81"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|song-url|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» url|string|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET mv-url

GET /music/mv

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

> mv-url

```json
{
  "code": 0,
  "result": {
    "url": ""
  },
  "msg": "该歌曲为付费内容，请下载酷我音乐客户端后付费收听",
  "details": null,
  "req_id": "cce2a81ad66ccd019a46b7234d6e90c3"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|mv-url|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» url|string|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET song-info

GET /music/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|
|Content-Type|header|string| 是 |none|

> 返回示例

> song-info

```json
{
  "code": 0,
  "result": {
    "img_url": "https://img1.kuwo.cn/star/albumcover/500/78/79/3920698290.jpg",
    "name": "等你归来",
    "id": 162457325,
    "score": 87,
    "has_loss_less": false,
    "release_date": "2021-01-17",
    "img_120": "https://img1.kuwo.cn/star/albumcover/120/78/79/3920698290.jpg",
    "album": {
      "img_url": "https://img1.kuwo.cn/star/albumcover/500/78/79/3920698290.jpg",
      "name": "等你归来",
      "id": 0
    },
    "artist": {
      "name": "程响",
      "id": 9009
    },
    "duration": {
      "duration": 235,
      "time_minutes": "03:55"
    },
    "has_mv": 1,
    "mv_id": 9979740
  },
  "msg": "success",
  "details": null,
  "req_id": "9587340bad112da95e68e75f6594781b"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|song-info|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» img_url|string|true|none||none|
|»» name|string|true|none||none|
|»» id|integer|true|none||none|
|»» score|integer|true|none||none|
|»» has_loss_less|boolean|true|none||none|
|»» release_date|string|true|none||none|
|»» img_120|string|true|none||none|
|»» album|object|true|none||none|
|»»» img_url|string|true|none||none|
|»»» name|string|true|none||none|
|»»» id|integer|true|none||none|
|»» artist|object|true|none||none|
|»»» name|string|true|none||none|
|»»» id|integer|true|none||none|
|»» duration|object|true|none||none|
|»»» duration|integer|true|none||none|
|»»» time_minutes|string|true|none||none|
|»» has_mv|integer|true|none||none|
|»» mv_id|integer|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

# rank

## GET rank

GET /rank

> 返回示例

> rank

```json
{
  "code": 0,
  "result": [
    {
      "category": "官方",
      "list": [
        {
          "intro": "酷我每日搜索热度飙升最快的歌曲排行榜，按搜索播放数据对比前一天涨幅排序，每天更新",
          "name": "酷我飙升榜",
          "id": 93,
          "img_url": "https://img3.kuwo.cn/star/upload/2/5/1659322379.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放新歌（30天内)排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "酷我新歌榜",
          "id": 17,
          "img_url": "https://img3.kuwo.cn/star/upload/9/5/1659322362.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放热度最高的歌曲排行榜，按歌曲热度排序，每天更新",
          "name": "酷我热歌榜",
          "id": 16,
          "img_url": "https://img3.kuwo.cn/star/upload/0/2/1659322374.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放抖音歌曲排行榜，按歌曲热度排序，每天更新",
          "name": "抖音歌曲榜",
          "id": 158,
          "img_url": "https://img3.kuwo.cn/star/upload/3/6/1659322381.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放DJ歌曲排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "万物DJ榜",
          "id": 176,
          "img_url": "https://img3.kuwo.cn/star/upload/5/6/1659322402.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我会员用户最喜爱的付费歌曲排行榜，按歌曲热度、飙升幅度综合排序，每天更新",
          "name": "会员畅听榜",
          "id": 145,
          "img_url": "https://img3.kuwo.cn/star/upload/6/9/1659322390.png",
          "update_time": "今日更新"
        }
      ]
    },
    {
      "category": "特色",
      "list": [
        {
          "intro": "酷我用户每天播放电音歌曲排行榜，按歌曲热度排序，每天更新",
          "name": "极品电音榜",
          "id": 242,
          "img_url": "https://img3.kuwo.cn/star/upload/9/4/1659322318.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天讨论度最高的歌曲排行榜，按歌曲热度、评论综合排序，每天更新",
          "name": "酷我热评榜",
          "id": 284,
          "img_url": "https://img3.kuwo.cn/star/upload/2/3/1659322323.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我每日搜索热度飙升最快的歌曲排行榜，按歌曲热度、发行天数综合排序，每天更新",
          "name": "流行趋势榜",
          "id": 187,
          "img_url": "https://img3.kuwo.cn/star/upload/5/4/1659322241.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放综艺歌曲排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "酷我综艺榜",
          "id": 154,
          "img_url": "https://img3.kuwo.cn/star/upload/8/6/1659322321.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放网红新歌（90天内发行）排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "网红新歌榜",
          "id": 153,
          "img_url": "https://img3.kuwo.cn/star/upload/7/2/1659322273.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放怀旧歌曲排行榜，按歌曲热度排序，每天更新",
          "name": "经典怀旧榜",
          "id": 26,
          "img_url": "https://img3.kuwo.cn/star/upload/3/7/1659322277.png",
          "update_time": "07月26日更新"
        },
        {
          "intro": "酷我用户每天播放说唱歌曲排行榜，按歌曲热度排序，每天更新",
          "name": "酷我说唱榜",
          "id": 329,
          "img_url": "https://img3.kuwo.cn/star/upload/4/2/1659322342.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放古风歌曲排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "古风音乐榜",
          "id": 278,
          "img_url": "https://img3.kuwo.cn/star/upload/5/9/1659322314.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放翻唱歌曲排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "最强翻唱榜",
          "id": 185,
          "img_url": "https://img3.kuwo.cn/star/upload/6/4/1659322231.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放影视歌曲（5年内发行）排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "影视金曲榜",
          "id": 64,
          "img_url": "https://img3.kuwo.cn/star/upload/8/5/1659322275.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放佛乐排行榜，按歌曲热度排序，每天更新",
          "name": "禅心佛乐榜",
          "id": 294,
          "img_url": "https://img3.kuwo.cn/star/upload/6/6/1659322308.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放相声排行榜，按热度排序，每天更新",
          "name": "爆笑相声榜",
          "id": 291,
          "img_url": "https://img3.kuwo.cn/star/upload/9/5/1659322339.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放ACG新歌（180天内发行）排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "ACG新歌榜",
          "id": 290,
          "img_url": "https://img3.kuwo.cn/star/upload/4/5/1659322333.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放铃声排行榜，按歌曲热度排序，每天更新",
          "name": "酷我铃声榜",
          "id": 292,
          "img_url": "https://img3.kuwo.cn/star/upload/4/5/1659322337.png",
          "update_time": "今日更新"
        }
      ]
    },
    {
      "category": "场景",
      "list": [
        {
          "intro": "酷我用户每天播放适合春天听的歌曲排行榜，按歌曲热度排序，每天更新",
          "name": "夏日畅爽榜",
          "id": 279,
          "img_url": "https://img3.kuwo.cn/star/upload/9/2/1659322222.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放车载音乐排行榜，按歌曲热度排序，每天更新",
          "name": "车载歌曲榜",
          "id": 328,
          "img_url": "https://img3.kuwo.cn/star/upload/7/4/1659322227.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放跑步健身歌曲排行榜，按歌曲热度排序，每天更新",
          "name": "跑步健身榜",
          "id": 297,
          "img_url": "https://img3.kuwo.cn/star/upload/4/5/1659322220.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放宝宝哄睡歌曲排行榜，按歌曲热度排序，每天更新",
          "name": "宝宝哄睡榜",
          "id": 295,
          "img_url": "https://img3.kuwo.cn/star/upload/3/5/1659322217.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放放松轻音乐排行榜，按歌曲热度排序，每天更新",
          "name": "睡前放松榜",
          "id": 283,
          "img_url": "https://img3.kuwo.cn/star/upload/9/9/1659322203.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放熬夜提神歌曲排行榜，按歌曲热度排序，每天更新",
          "name": "熬夜修仙榜",
          "id": 282,
          "img_url": "https://img3.kuwo.cn/star/upload/7/6/1659322225.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放vlog音乐排行榜，按歌曲热度排序，每天更新",
          "name": "Vlog必备榜",
          "id": 264,
          "img_url": "https://img3.kuwo.cn/star/upload/5/8/1659322214.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放KTV歌曲排行榜，按歌曲热度排序，每天更新",
          "name": "KTV点唱榜",
          "id": 255,
          "img_url": "https://img3.kuwo.cn/star/upload/1/7/1659322212.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放通勤音乐排行榜，按歌曲热度排序，每天更新",
          "name": "通勤路上榜",
          "id": 281,
          "img_url": "https://img3.kuwo.cn/star/upload/6/1/1659322205.png",
          "update_time": "今日更新"
        }
      ]
    },
    {
      "category": "语言",
      "list": [
        {
          "intro": "酷我用户每天播放华语歌曲排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "酷我华语榜",
          "id": 104,
          "img_url": "https://img3.kuwo.cn/star/upload/6/6/1659322418.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放粤语歌曲排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "酷我粤语榜",
          "id": 182,
          "img_url": "https://img3.kuwo.cn/star/upload/8/0/1659322412.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放欧美歌曲排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "酷我欧美榜",
          "id": 22,
          "img_url": "https://img3.kuwo.cn/star/upload/7/1/1659322408.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放韩语歌曲排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "酷我韩语榜",
          "id": 184,
          "img_url": "https://img3.kuwo.cn/star/upload/5/4/1659322410.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我用户每天播放日语歌曲排行榜，按歌曲热度、发行时间综合排序，每天更新",
          "name": "酷我日语榜",
          "id": 183,
          "img_url": "https://img3.kuwo.cn/star/upload/5/6/1659322416.png",
          "update_time": "今日更新"
        }
      ]
    },
    {
      "category": "全球",
      "list": [
        {
          "intro": "美国公告牌Billboard Hot100单曲榜，每周三更新",
          "name": "美国The Billboard",
          "id": 12,
          "img_url": "https://img4.kuwo.cn/star/upload/10/10/1649672653402_.jpg",
          "update_time": "07月28日更新"
        },
        {
          "intro": "iTunes全美流行音乐综合排名，每周一更新",
          "name": "iTunes音乐榜",
          "id": 49,
          "img_url": "https://img4.kuwo.cn/star/upload/9/9/1649672678649_.jpg",
          "update_time": "07月26日更新"
        },
        {
          "intro": "beatport全球电子舞曲排行榜，每周二更新",
          "name": "beatport电音榜",
          "id": 180,
          "img_url": "https://img4.kuwo.cn/star/upload/0/0/1649672695520_.jpg",
          "update_time": "07月27日更新"
        },
        {
          "intro": "英国单曲排行榜（UK Singles Chart）是代表英国唱片工业的权威排行榜，每周六更新",
          "name": "英国UK Official",
          "id": 13,
          "img_url": "https://img4.kuwo.cn/star/upload/2/2/1649672715090_.jpg",
          "update_time": "07月31日更新"
        },
        {
          "intro": "酷我用户每周播放世界百大DJs歌曲TOP排行榜，每周二更新",
          "name": "百大DJ榜",
          "id": 164,
          "img_url": "https://img3.kuwo.cn/star/upload/10/10/1649672739226_.jpg",
          "update_time": "今日更新"
        },
        {
          "intro": "YouTube是一个视频网站，YouTube音乐排行榜是根据歌曲每周在YouTube的播放量进行排名，每周一更新。",
          "name": "YouTube音乐排行榜",
          "id": 246,
          "img_url": "https://img3.kuwo.cn/star/upload/7/7/1649672755016_.jpg",
          "update_time": "07月26日更新"
        },
        {
          "intro": "韩国Genie音乐榜，每周一更新。",
          "name": "韩国Genie榜",
          "id": 265,
          "img_url": "https://img4.kuwo.cn/star/upload/8/8/1649672772920_.jpg",
          "update_time": "07月26日更新"
        },
        {
          "intro": "日本公信榜也称为Oricon榜，是日本最具知名度的音乐排行榜，每周三更新",
          "name": "日本公信榜",
          "id": 15,
          "img_url": "https://img2.kuwo.cn/star/upload/14/14/1649672798206_.jpg",
          "update_time": "07月28日更新"
        },
        {
          "intro": "SPACE SHOWER TV设立于1989年，因通过MV和直播节目介绍了日本的众多艺人而备受好评。 本榜单根据频道上的播放次数计算，为你展现日本音乐的流行趋势。 每周三更新。",
          "name": "Space Shower",
          "id": 293,
          "img_url": "https://img3.kuwo.cn/star/upload/3/3/1649672816467_.jpg",
          "update_time": "07月28日更新"
        }
      ]
    },
    {
      "category": "更多",
      "list": [
        {
          "intro": " 酷我用户每天播放综艺歌曲TOP排行榜，为你展示时下最新的综艺热歌，每天更新",
          "name": "综艺热歌榜",
          "id": 154,
          "img_url": "https://img3.kuwo.cn/star/upload/3/4/1659322354.png",
          "update_time": "今日更新"
        },
        {
          "intro": "酷我聚星热歌TOP排行榜，为你展示真声音·好音乐，带给你不一样的感觉",
          "name": "酷我真声音",
          "id": 106,
          "img_url": "https://img3.kuwo.cn/star/upload/4/2/1659322348.png",
          "update_time": "今日更新"
        },
        {
          "intro": "集结最新最热的腾讯音乐人原创作品，每天更新",
          "name": "腾讯音乐人原创榜",
          "id": 151,
          "img_url": "https://img3.kuwo.cn/star/upload/5/7/1659322347.png",
          "update_time": "今日更新"
        },
        {
          "intro": "家务BGM排行榜，每天更新",
          "name": "家务进行曲",
          "id": 280,
          "img_url": "https://img3.kuwo.cn/star/upload/9/7/1659322344.png",
          "update_time": "今日更新"
        },
        {
          "intro": "",
          "name": "现场音乐榜",
          "id": 204,
          "img_url": "https://img3.kuwo.cn/star/upload/4/0/1659322351.png",
          "update_time": "今日更新"
        }
      ]
    }
  ],
  "msg": "success",
  "details": null,
  "req_id": "b6fbd434fc090032d4c388737dd06122"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|rank|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|[object]|true|none||none|
|»» category|string|true|none||none|
|»» list|[object]|true|none||none|
|»»» intro|string|true|none||none|
|»»» name|string|true|none||none|
|»»» id|integer|true|none||none|
|»»» img_url|string|true|none||none|
|»»» update_time|string|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

## GET rank-music

GET /rank/93

> 返回示例

> rank-music

```json
{
  "code": 0,
  "result": {
    "pl_img_url": "https://img1.kuwo.cn/star/upload/11/11/1539055092699_.png",
    "total_cnt": 300,
    "update_time": "2022-08-01",
    "list": [
      {
        "img_url": "https://img1.kuwo.cn/star/albumcover/500/85/57/351640245.jpg",
        "name": "美人鱼",
        "id": 93151,
        "score": 87,
        "has_loss_less": true,
        "release_date": "2004-06-04",
        "img_120": "https://img1.kuwo.cn/star/albumcover/120/85/57/351640245.jpg",
        "album": {
          "img_url": "https://img1.kuwo.cn/star/albumcover/500/85/57/351640245.jpg",
          "name": "第二天堂",
          "id": 0
        },
        "artist": {
          "name": "林俊杰",
          "id": 1062
        },
        "duration": {
          "duration": 254,
          "time_minutes": "04:14"
        },
        "has_mv": 1,
        "mv_id": 8084200
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/93/50/1282097401.jpg",
        "name": "Hug me (抱我)",
        "id": 228912860,
        "score": 98,
        "has_loss_less": true,
        "release_date": "2022-07-22",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/93/50/1282097401.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/93/50/1282097401.jpg",
          "name": "Hug me (抱我)",
          "id": 0
        },
        "artist": {
          "name": "蔡徐坤",
          "id": 908596
        },
        "duration": {
          "duration": 150,
          "time_minutes": "02:30"
        },
        "has_mv": 1,
        "mv_id": 15981876
      },
      {
        "img_url": "https://img2.kuwo.cn/star/albumcover/500/16/91/409859677.jpg",
        "name": "大眠",
        "id": 58240758,
        "score": 79,
        "has_loss_less": true,
        "release_date": "2018-12-07",
        "img_120": "https://img2.kuwo.cn/star/albumcover/120/16/91/409859677.jpg",
        "album": {
          "img_url": "https://img2.kuwo.cn/star/albumcover/500/16/91/409859677.jpg",
          "name": "CYNDILOVES2SING 爱。心凌",
          "id": 0
        },
        "artist": {
          "name": "王心凌",
          "id": 1306
        },
        "duration": {
          "duration": 239,
          "time_minutes": "03:59"
        },
        "has_mv": 1,
        "mv_id": 6646137
      },
      {
        "img_url": "https://img4.kuwo.cn/star/albumcover/500/54/53/624321623.jpg",
        "name": "当你",
        "id": 169088,
        "score": 86,
        "has_loss_less": true,
        "release_date": "2003-02-24",
        "img_120": "https://img4.kuwo.cn/star/albumcover/120/54/53/624321623.jpg",
        "album": {
          "img_url": "https://img4.kuwo.cn/star/albumcover/500/54/53/624321623.jpg",
          "name": "Begin...",
          "id": 0
        },
        "artist": {
          "name": "王心凌",
          "id": 1306
        },
        "duration": {
          "duration": 203,
          "time_minutes": "03:23"
        },
        "has_mv": 1,
        "mv_id": 14992537
      },
      {
        "img_url": "https://img1.kuwo.cn/star/albumcover/500/20/47/836307178.jpg",
        "name": "忘了你忘了我",
        "id": 77861,
        "score": 74,
        "has_loss_less": true,
        "release_date": "1988-07-11",
        "img_120": "https://img1.kuwo.cn/star/albumcover/120/20/47/836307178.jpg",
        "album": {
          "img_url": "https://img1.kuwo.cn/star/albumcover/500/20/47/836307178.jpg",
          "name": "忘了你‧忘了我",
          "id": 0
        },
        "artist": {
          "name": "王杰",
          "id": 1892
        },
        "duration": {
          "duration": 279,
          "time_minutes": "04:39"
        },
        "has_mv": 1,
        "mv_id": 8254666
      },
      {
        "img_url": "https://img1.kuwo.cn/star/albumcover/500/24/85/915166694.jpg",
        "name": "雨中的恋人们",
        "id": 886594,
        "score": 75,
        "has_loss_less": true,
        "release_date": "1992-01-01",
        "img_120": "https://img1.kuwo.cn/star/albumcover/120/24/85/915166694.jpg",
        "album": {
          "img_url": "https://img1.kuwo.cn/star/albumcover/500/24/85/915166694.jpg",
          "name": "Stay With Me",
          "id": 0
        },
        "artist": {
          "name": "黄凯芹",
          "id": 77
        },
        "duration": {
          "duration": 260,
          "time_minutes": "04:20"
        },
        "has_mv": 1,
        "mv_id": 157586
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/47/33/2949180979.jpg",
        "name": "我是你的格桑花",
        "id": 46566459,
        "score": 86,
        "has_loss_less": true,
        "release_date": "2018-06-30",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/47/33/2949180979.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/47/33/2949180979.jpg",
          "name": "我是你的格桑花",
          "id": 0
        },
        "artist": {
          "name": "欣宝儿",
          "id": 2765400
        },
        "duration": {
          "duration": 312,
          "time_minutes": "05:12"
        },
        "has_mv": 1,
        "mv_id": 8251509
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/4/41/1226507593.jpg",
        "name": "愿你",
        "id": 207527604,
        "score": 96,
        "has_loss_less": true,
        "release_date": "2022-01-24",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/4/41/1226507593.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/4/41/1226507593.jpg",
          "name": "愿你",
          "id": 0
        },
        "artist": {
          "name": "黄静美",
          "id": 3652009
        },
        "duration": {
          "duration": 183,
          "time_minutes": "03:03"
        },
        "has_mv": 1,
        "mv_id": 14341939
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/49/79/1011385325.jpg",
        "name": "孤勇者",
        "id": 198554068,
        "score": 105,
        "has_loss_less": true,
        "release_date": "2021-11-08",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/49/79/1011385325.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/49/79/1011385325.jpg",
          "name": "孤勇者",
          "id": 0
        },
        "artist": {
          "name": "陈奕迅",
          "id": 47
        },
        "duration": {
          "duration": 256,
          "time_minutes": "04:16"
        },
        "has_mv": 1,
        "mv_id": 11493948
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/76/68/4090861467.jpg",
        "name": "我们都是孤单的人(加速版|Remix)",
        "id": 224801308,
        "score": 87,
        "has_loss_less": true,
        "release_date": "2022-04-24",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/76/68/4090861467.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/76/68/4090861467.jpg",
          "name": "我们都是孤单的人",
          "id": 0
        },
        "artist": {
          "name": "王羽泽",
          "id": 4812
        },
        "duration": {
          "duration": 180,
          "time_minutes": "03:00"
        },
        "has_mv": 0,
        "mv_id": 16135975
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/60/81/812831951.jpg",
        "name": "那年夏天宁静的海",
        "id": 239411,
        "score": 74,
        "has_loss_less": true,
        "release_date": "2007-04-30",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/60/81/812831951.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/60/81/812831951.jpg",
          "name": "Magic Cyndi",
          "id": 0
        },
        "artist": {
          "name": "王心凌",
          "id": 1306
        },
        "duration": {
          "duration": 253,
          "time_minutes": "04:13"
        },
        "has_mv": 1,
        "mv_id": 14984187
      },
      {
        "img_url": "https://img2.kuwo.cn/star/albumcover/500/77/37/2775392033.jpg",
        "name": "用余生忘记你",
        "id": 171315993,
        "score": 77,
        "has_loss_less": true,
        "release_date": "2021-03-20",
        "img_120": "https://img2.kuwo.cn/star/albumcover/120/77/37/2775392033.jpg",
        "album": {
          "img_url": "https://img2.kuwo.cn/star/albumcover/500/77/37/2775392033.jpg",
          "name": "用余生忘记你",
          "id": 0
        },
        "artist": {
          "name": "魏佳艺",
          "id": 4543
        },
        "duration": {
          "duration": 287,
          "time_minutes": "04:47"
        },
        "has_mv": 1,
        "mv_id": 14797619
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/0/91/1985244086.jpg",
        "name": "爱你",
        "id": 89622,
        "score": 93,
        "has_loss_less": true,
        "release_date": "2004-03-26",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/0/91/1985244086.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/0/91/1985244086.jpg",
          "name": "爱你",
          "id": 0
        },
        "artist": {
          "name": "王心凌",
          "id": 1306
        },
        "duration": {
          "duration": 219,
          "time_minutes": "03:39"
        },
        "has_mv": 1,
        "mv_id": 14990940
      },
      {
        "img_url": "https://img2.kuwo.cn/star/albumcover/500/64/22/1839376303.jpg",
        "name": "半生雪",
        "id": 177504089,
        "score": 100,
        "has_loss_less": true,
        "release_date": "2021-05-14",
        "img_120": "https://img2.kuwo.cn/star/albumcover/120/64/22/1839376303.jpg",
        "album": {
          "img_url": "https://img2.kuwo.cn/star/albumcover/500/64/22/1839376303.jpg",
          "name": "半生雪",
          "id": 0
        },
        "artist": {
          "name": "七叔（叶泽浩）",
          "id": 4355654
        },
        "duration": {
          "duration": 176,
          "time_minutes": "02:56"
        },
        "has_mv": 1,
        "mv_id": 13525244
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/32/88/1567952196.jpg",
        "name": "夜曲",
        "id": 118980,
        "score": 90,
        "has_loss_less": true,
        "release_date": "2005-11-01",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/32/88/1567952196.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/32/88/1567952196.jpg",
          "name": "十一月的萧邦",
          "id": 0
        },
        "artist": {
          "name": "周杰伦",
          "id": 336
        },
        "duration": {
          "duration": 226,
          "time_minutes": "03:46"
        },
        "has_mv": 1,
        "mv_id": 8164146
      },
      {
        "img_url": "https://img1.kuwo.cn/star/albumcover/500/0/11/759002399.jpg",
        "name": "踏山河",
        "id": 156483846,
        "score": 100,
        "has_loss_less": true,
        "release_date": "2020-11-19",
        "img_120": "https://img1.kuwo.cn/star/albumcover/120/0/11/759002399.jpg",
        "album": {
          "img_url": "https://img1.kuwo.cn/star/albumcover/500/0/11/759002399.jpg",
          "name": "踏山河",
          "id": 0
        },
        "artist": {
          "name": "七叔（叶泽浩）",
          "id": 4355654
        },
        "duration": {
          "duration": 168,
          "time_minutes": "02:48"
        },
        "has_mv": 1,
        "mv_id": 13286493
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/16/13/869267863.jpg",
        "name": "樱花树下的约定（完整版）",
        "id": 218453668,
        "score": 96,
        "has_loss_less": true,
        "release_date": "2022-04-27",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/16/13/869267863.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/16/13/869267863.jpg",
          "name": "樱花树下的约定",
          "id": 0
        },
        "artist": {
          "name": "旺仔小乔",
          "id": 5237045
        },
        "duration": {
          "duration": 177,
          "time_minutes": "02:57"
        },
        "has_mv": 0,
        "mv_id": 16118374
      },
      {
        "img_url": "https://img1.kuwo.cn/star/albumcover/500/68/73/1467360956.jpg",
        "name": "眼中有泪心中有你",
        "id": 225631071,
        "score": 72,
        "has_loss_less": true,
        "release_date": "2022-07-04",
        "img_120": "https://img1.kuwo.cn/star/albumcover/120/68/73/1467360956.jpg",
        "album": {
          "img_url": "https://img1.kuwo.cn/star/albumcover/500/68/73/1467360956.jpg",
          "name": "眼中有泪心中有你",
          "id": 0
        },
        "artist": {
          "name": "李乐乐",
          "id": 687085
        },
        "duration": {
          "duration": 215,
          "time_minutes": "03:35"
        },
        "has_mv": 0,
        "mv_id": 16137438
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/99/73/3563978030.jpg",
        "name": "月老掉线(DJ阿卓版)",
        "id": 207231743,
        "score": 98,
        "has_loss_less": true,
        "release_date": "2022-01-16",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/99/73/3563978030.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/99/73/3563978030.jpg",
          "name": "月老掉线",
          "id": 0
        },
        "artist": {
          "name": "王不醒",
          "id": 9043816
        },
        "duration": {
          "duration": 206,
          "time_minutes": "03:26"
        },
        "has_mv": 0,
        "mv_id": 13982001
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/24/43/965300864.jpg",
        "name": "听我说谢谢你",
        "id": 70341882,
        "score": 88,
        "has_loss_less": true,
        "release_date": "2019-06-12",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/24/43/965300864.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/24/43/965300864.jpg",
          "name": "听我说谢谢你",
          "id": 0
        },
        "artist": {
          "name": "李昕融",
          "id": 1373230
        },
        "duration": {
          "duration": 187,
          "time_minutes": "03:07"
        },
        "has_mv": 1,
        "mv_id": 8254144
      },
      {
        "img_url": "https://img4.kuwo.cn/star/albumcover/500/65/89/2489271996.jpg",
        "name": "给你一瓶魔法药水",
        "id": 225785286,
        "score": 65,
        "has_loss_less": false,
        "release_date": "2022-07-04",
        "img_120": "https://img4.kuwo.cn/star/albumcover/120/65/89/2489271996.jpg",
        "album": {
          "img_url": "https://img4.kuwo.cn/star/albumcover/500/65/89/2489271996.jpg",
          "name": "玫瑰凭证",
          "id": 0
        },
        "artist": {
          "name": "告五人",
          "id": 1725643
        },
        "duration": {
          "duration": 258,
          "time_minutes": "04:18"
        },
        "has_mv": 0,
        "mv_id": 16269574
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/0/91/1985244086.jpg",
        "name": "第一次爱的人",
        "id": 89626,
        "score": 86,
        "has_loss_less": true,
        "release_date": "2004-03-26",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/0/91/1985244086.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/0/91/1985244086.jpg",
          "name": "爱你",
          "id": 0
        },
        "artist": {
          "name": "王心凌",
          "id": 1306
        },
        "duration": {
          "duration": 233,
          "time_minutes": "03:53"
        },
        "has_mv": 1,
        "mv_id": 14992540
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/15/10/12649691.jpg",
        "name": "破茧",
        "id": 140897945,
        "score": 97,
        "has_loss_less": true,
        "release_date": "2020-05-23",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/15/10/12649691.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/15/10/12649691.jpg",
          "name": "破茧",
          "id": 0
        },
        "artist": {
          "name": "张韶涵",
          "id": 492
        },
        "duration": {
          "duration": 211,
          "time_minutes": "03:31"
        },
        "has_mv": 1,
        "mv_id": 8779833
      },
      {
        "img_url": "https://img2.kuwo.cn/star/albumcover/500/64/96/2266534336.jpg",
        "name": "晴天",
        "id": 228908,
        "score": 93,
        "has_loss_less": true,
        "release_date": "2003-07-31",
        "img_120": "https://img2.kuwo.cn/star/albumcover/120/64/96/2266534336.jpg",
        "album": {
          "img_url": "https://img2.kuwo.cn/star/albumcover/500/64/96/2266534336.jpg",
          "name": "叶惠美",
          "id": 0
        },
        "artist": {
          "name": "周杰伦",
          "id": 336
        },
        "duration": {
          "duration": 269,
          "time_minutes": "04:29"
        },
        "has_mv": 1,
        "mv_id": 8132306
      },
      {
        "img_url": "https://img4.kuwo.cn/star/albumcover/500/58/33/3231164701.jpg",
        "name": "对你说",
        "id": 228770810,
        "score": 89,
        "has_loss_less": true,
        "release_date": "2022-07-19",
        "img_120": "https://img4.kuwo.cn/star/albumcover/120/58/33/3231164701.jpg",
        "album": {
          "img_url": "https://img4.kuwo.cn/star/albumcover/500/58/33/3231164701.jpg",
          "name": "对你说",
          "id": 0
        },
        "artist": {
          "name": "半吨兄弟",
          "id": 3670041
        },
        "duration": {
          "duration": 244,
          "time_minutes": "04:04"
        },
        "has_mv": 0,
        "mv_id": 16310531
      },
      {
        "img_url": "https://img1.kuwo.cn/star/albumcover/500/30/48/2563040593.jpg",
        "name": "你",
        "id": 1629262,
        "score": 76,
        "has_loss_less": true,
        "release_date": "2012-07-15",
        "img_120": "https://img1.kuwo.cn/star/albumcover/120/30/48/2563040593.jpg",
        "album": {
          "img_url": "https://img1.kuwo.cn/star/albumcover/500/30/48/2563040593.jpg",
          "name": "自己人",
          "id": 0
        },
        "artist": {
          "name": "屠洪刚",
          "id": 443
        },
        "duration": {
          "duration": 255,
          "time_minutes": "04:15"
        },
        "has_mv": 1,
        "mv_id": 3072
      },
      {
        "img_url": "https://img3.kuwo.cn/star/albumcover/500/60/8/1867172466.jpg",
        "name": "爱我的人和我爱的人",
        "id": 229744754,
        "score": 0,
        "has_loss_less": true,
        "release_date": "2022-07-27",
        "img_120": "https://img3.kuwo.cn/star/albumcover/120/60/8/1867172466.jpg",
        "album": {
          "img_url": "https://img3.kuwo.cn/star/albumcover/500/60/8/1867172466.jpg",
          "name": "爱我的人和我爱的人",
          "id": 0
        },
        "artist": {
          "name": "半吨兄弟",
          "id": 3670041
        },
        "duration": {
          "duration": 246,
          "time_minutes": "04:06"
        },
        "has_mv": 0,
        "mv_id": 16285857
      },
      {
        "img_url": "https://img4.kuwo.cn/star/albumcover/500/98/24/2430824662.jpg",
        "name": "旧梦(DJ默涵版)",
        "id": 217710574,
        "score": 80,
        "has_loss_less": true,
        "release_date": "2022-04-16",
        "img_120": "https://img4.kuwo.cn/star/albumcover/120/98/24/2430824662.jpg",
        "album": {
          "img_url": "https://img4.kuwo.cn/star/albumcover/500/98/24/2430824662.jpg",
          "name": "旧梦",
          "id": 0
        },
        "artist": {
          "name": "刘晓超",
          "id": 3499966
        },
        "duration": {
          "duration": 183,
          "time_minutes": "03:03"
        },
        "has_mv": 0,
        "mv_id": 16118450
      },
      {
        "img_url": "https://img1.kuwo.cn/star/albumcover/500/97/94/3750389106.jpg",
        "name": "秒针(Dj版)",
        "id": 198635187,
        "score": 97,
        "has_loss_less": true,
        "release_date": "2021-11-09",
        "img_120": "https://img1.kuwo.cn/star/albumcover/120/97/94/3750389106.jpg",
        "album": {
          "img_url": "https://img1.kuwo.cn/star/albumcover/500/97/94/3750389106.jpg",
          "name": "秒针",
          "id": 0
        },
        "artist": {
          "name": "阿梨粤&R7",
          "id": 4925534
        },
        "duration": {
          "duration": 178,
          "time_minutes": "02:58"
        },
        "has_mv": 0,
        "mv_id": 16277250
      },
      {
        "img_url": "https://img1.kuwo.cn/star/albumcover/500/67/0/2275517451.jpg",
        "name": "So Far Away",
        "id": 36290894,
        "score": 92,
        "has_loss_less": true,
        "release_date": "2017-12-01",
        "img_120": "https://img1.kuwo.cn/star/albumcover/120/67/0/2275517451.jpg",
        "album": {
          "img_url": "https://img1.kuwo.cn/star/albumcover/500/67/0/2275517451.jpg",
          "name": "So Far Away",
          "id": 0
        },
        "artist": {
          "name": "Martin Garrix&David Guetta&Jamie Scott&Romy Dya",
          "id": 116361
        },
        "duration": {
          "duration": 183,
          "time_minutes": "03:03"
        },
        "has_mv": 1,
        "mv_id": 5781546
      }
    ]
  },
  "msg": "success",
  "details": null,
  "req_id": "9f9192bb4570a4a55a37fb840c0e4187"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|rank-music|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» result|object|true|none||none|
|»» pl_img_url|string|true|none||none|
|»» total_cnt|integer|true|none||none|
|»» update_time|string|true|none||none|
|»» list|[object]|true|none||none|
|»»» img_url|string|true|none||none|
|»»» name|string|true|none||none|
|»»» id|integer|true|none||none|
|»»» score|integer|true|none||none|
|»»» has_loss_less|boolean|true|none||none|
|»»» release_date|string|true|none||none|
|»»» img_120|string|true|none||none|
|»»» album|object|true|none||none|
|»»»» img_url|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|»»» artist|object|true|none||none|
|»»»» name|string|true|none||none|
|»»»» id|integer|true|none||none|
|»»» duration|object|true|none||none|
|»»»» duration|integer|true|none||none|
|»»»» time_minutes|string|true|none||none|
|»»» has_mv|integer|true|none||none|
|»»» mv_id|integer|true|none||none|
|» msg|string|true|none||none|
|» details|null|true|none||none|
|» req_id|string|true|none||none|

# 数据模型

