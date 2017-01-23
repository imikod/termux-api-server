package main

var index = `
<html>
<head>
    <title>termux-api-server</title>
<style>
html,body {
	height: 100%;
	margin: 0;
	padding: 0;
	font-family: sans-serif;
	background-color: #fdfdfd;
}
#main {
	margin: auto;
	max-width: 40em;
	padding: 5em 1em;
	line-height: 1.7em;
	color: #336;
}
pre {
	font-family: monospace;
	padding: 0em 1em;
	background-color: #eee;
	color: #336;
	overflow: auto;
}
a {
	color: #2a7ae2;
	text-decoration: none;
}
a:visited {
	color: #1756a9;
}
a:hover {
	color: #111;
	text-decoration: underline;
}
</style>
</head>
<body>
<div id="main">

<h1>Termux Services</h1>
Simple server to expose the
<a href="https://termux.com/add-on-api.html">termux-api</a>
with CORS
<br><br>

<h3><a href="/battery-status">termux-battery-status</a></h3>
<pre>GET /battery-status</pre>
Get the status of the device battery.
<br><br>

<h3><a href="/camera-info">termux-camera-info</a></h3>
<pre>GET /camera-info</pre>
Get information about device camera(s).
<br><br>

<h3><a href="/camera-photo">termux-camera-photo</a></h3>
<pre>GET /camera-photo?c={cameraID}</pre>
Take a photo and view it in JPEG format.
<br><br>

<h3><a href="/contact-list">termux-contact-list</a></h3>
<pre>GET /contact-list</pre>
List all contacts.
<br><br>

<h3><a href="/infrared-frequencies">termux-infrared-frequencies</a></h3>
<pre>GET /infrared-frequencies</pre>
Query the infrared transmitter's supported carrier frequencies.
<br><br>

<h3><a href="/location">termux-location</a></h3>
<pre>GET /location?p={provider}</pre>
Get the device location.
<br><br>

<h3><a href="/notification?c=Hello!">termux-notification</a></h3>
<pre>GET /notification?c={content}&i={id}&t={title}&u={url}</pre>
Display a system notification.
<br><br>

<h3><a href="/sms-inbox">termux-sms-inbox</a></h3>
<pre>GET /sms-inbox?d={true}&l={limit}&n={true}&o={offset}</pre>
List received SMS messages.
<br><br>

<h3><a href="/sms-send">termux-sms-send</a></h3>
<pre>GET /sms-send?n={number1}&n={number2}&n={number3}&t={text}</pre>
Send a SMS message to the specified recipient number(s).
<br><br>

<h3><a href="/telephony-cellinfo">termux-telephony-cellinfo</a></h3>
<pre>GET /telephony-cellinfo</pre>
Get information about all observed cell information from all radios
on the device including the primary and neighboring cells.
<br><br>

<h3><a href="/telephony-deviceinfo">termux-telephony-deviceinfo</a></h3>
<pre>GET /telephony-deviceinfo</pre>
Get information about the telephony device.
<br><br>

<h3><a href="/tts-engines">termux-tts-engines</a></h3>
<pre>GET /tts-engines</pre>
Get information about the available text-to-speech (TTS) engines.
The name of an engine may be given to the termux-tts-speak command using the -e option.
<br><br>

<h3><a href="/tts-speak?t=Hello">termux-tts-speak</a></h3>
<pre>GET /tts-speak?e={engine}&l={language}&p={pitch}&r={rate}&s={stream}&t={text}</pre>
Speak text with a system text-to-speech (TTS) engine.
<br><br>

<h3><a href="/vibrate">termux-vibrate</a></h3>
<pre>GET /vibrate?d={duration}&f={true}</pre>
Vibrate the device.
<br><br>

</div>
</body>
</html>
`
