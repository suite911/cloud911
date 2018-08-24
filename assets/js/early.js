'use strict';

function amy_addEventListener(elem, on, cb, useCapture) {
	if(!elem) {
		return false;
	}
	if(typeof elem.addEventListener === "function") {
		if(useCapture === "undefined") {
			useCapture = false;
		}
		elem.addEventListener(on, cb, useCapture);
	} else if(typeof elem.attachEvent === "function") {
		elem.attachEvent("on" + on, cb);
	} else {
		elem["on" + on] = cb;
	}
	return true;
}

var amy_cookieAlert = function() {
	alert("This site uses cookies to enhance the user experience.");
	return "1"
}

var amy_cookieAgree = function() {
	if (amy_getCookie("agreed") == "") {
		var value = "1";
		if(typeof amy_cookieAlert === "function") {
			value = amy_cookieAlert();
		}
		amy_setCookie("agreed", value, 1);
	}
}

function amy_defer(obj, mthd) {
	if(typeof obj[mthd] === "function") {
		obj[mthd].apply(obj, [].slice.call(arguments).slice(2));
		return true;
	}
	return false;
}

function amy_getCookie(name) {
	var n = name + "=";
	var a = document.cookie.split(';');
	for(var i = 0; i < a.length; i++) {
		var c = a[i];
		while(c.charAt(0) == ' ') {
			c = c.substring(1);
		}
		if (c.indexOf(n) == 0) {
			return c.substring(n.length, c.length);
		}
	}
	return "";
}

function amy_hasClass(elem, className) {
	return (' ' + elem.className + ' ').indexOf(' ' + className + ' ') > -1;
}

var amy_onChanged = function() { }
var amy_onDOMReady = function() { }
var amy_onPageLoaded = function() { }
var amy_onSubmitted = function() { }
var amy_onWillSubmit = function() { }
var amy_reCAPTCHAv3SiteKey = ""

function amy_removeEventListener(elem, on, cb, useCapture) {
	if(!elem) {
		return false;
	}
	if(typeof elem.removeEventListener === "function") {
		if(useCapture === "undefined") {
			useCapture = false;
		}
		elem.removeEventListener(on, cb, useCapture);
	} else if(typeof elem.detachEvent === "function") {
		elem.detachEvent("on" + on, cb);
	} else {
		elem["on" + on] = null;
	}
	return true;
}

function amy_replaceState(url) {
	if(typeof history.replaceState === "function") {
		history.replaceState(null, null, url);
		return true;
	}
	return false;
}

function amy_setCookie(name, value, hours) {
	var nv = name + "=" + value;
	var p = ";path=/";
	if(hours === undefined) {
		document.cookie = nv + p;
	} else {
		var d = new Date();
		d.setTime(d.getTime() + (hours * 3600000 ));
		var x = ";expires=" + d.toUTCString();
		document.cookie = nv + x + p;
	}
}
