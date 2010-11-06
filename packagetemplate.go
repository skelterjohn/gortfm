package main

const packageTemplateStr = `<html>
<head>
	<title>Package {pkgname}</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<link rel="stylesheet" href="shared/gortfm.css" />
	<script type="text/javascript" src="{datafile}"></script>
	<script type="text/javascript" src="shared/jquery-1.4.2.min.js"></script>
	<script type="text/javascript" src="shared/gortfm-help.js"></script>
	<script type="text/javascript" src="shared/gortfm-fuzzy.js"></script>
	<script type="text/javascript" src="shared/gortfm.js"></script>
</head>
<body>

<div id="header">
	<div class="line">
		<input id="filter" class="inactive" placeholder="type ? for help" />
		<span>
			Package <a class="black" href="?">{pkgname}</a>
		</span>
	</div>
</div>

<div id="contents">
</div>

<div id="footer">
	<div>Powered by: 
		<a href="http://jquery.com">jQuery</a>
	</div>
</div>
</body>
</html>
`