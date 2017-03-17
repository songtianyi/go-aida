<!DOCTYPE html>

<html>
<head>
  <title>Go-Aida</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    .logo {
      background-repeat: no-repeat;
      -webkit-background-size: 100px 100px;
      background-size: 100px 100px;
      background-position: center center;
      text-align: center;
      font-size: 42px;
      padding: 100px 0 70px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }

    header {
      padding: 100px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
  </style>
</head>

<body>
  <header>
    <h1 class="logo">Welcome, I'm Aida</h1>
    <div class="description">
      Aida是一个功能型wechat机器人, 名称来源于美剧《神盾局特工》里的类人机器人Aida.
	 </div>
  </header>

	 <div class="description">
	 功能列表
	  <ul>
	  	<li>自动回复</li>
	  	<li>追剧提醒</li>
	  	<li>聊天自带gif</li>
	  	<li>头像性别/年龄识别</li>
		<li>美女图片</li>
		</ul>
    </div>
	<div class="description">
	指令列表
	<ul>
	  	<li>fache</li>
	</ul>
	</div>

  <footer>
    <div class="author">
      Project website:
      <a href="http://github.com/songtianyi/{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
