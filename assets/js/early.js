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

function amy_cookieAlert() {
	alert("This site uses cookies to enhance the user experience.");
	return "1"
}

function amy_cookieAgree() {
	if (amy_getCookie("agreed") == "") {
		amy_setCookie("agreed", amy_cookieAlert(), 1);
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

var amy_onDOMReady() = function() { }

function amy_onDOMReadyTail() {
	if(location.hash.length >= 2) {
		var elem = document.getElementById(location.hash.slice(1))
		if(elem) {
			if(amy_hasClass(elem, "fragment-block")) {
				elem.style.display = "block";
			} else if(amy_hasClass(elem, "fragment-inline")) {
				elem.style.display = "inline";
			} else if(amy_hasClass(elem, "fragment-inline-block")) {
				elem.style.display = "inline-block";
			}
			if(!replaceState(location.href.split('#')[0])) {
				location.hash = '';
			}
		}
	} else if(location.href.slice(-1) == '#') {
		replaceState(location.href.slice(0, -1));
	}
}

function amy_onLightsChanged(event) {
	var lightSwitch = document.getElementById("lights-off");
	if(lightSwitch) {
		if(lightSwitch.checked) {
			amy_setCookie("lights", "off");
		} else {
			amy_setCookie("lights", "on");
		}
	}
}

var amy_onPageLoaded() = function() { }

function amy_onSubmit {
	if(amy_onWillSubmit === "function") {
		var ok = amy_onWillSubmit();
		if(typeof ok !== "undefined" && !ok) {
			return false;
		}
	}
	var elem = document.getElementById("form");
	if(typeof elem.submit === "function") {
		elem.submit();
	}
	if(amy_onSubmitted === "function") {
		return amy_onSubmitted();
	}
}

var amy_onSubmitted = function() { }
var amy_onWillSubmit = function() { }

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

function amy_onDOMReadyHead() {
	var content = document.getElementById("content");
	if(content) {
		content.style.display = "block";
	}
	var lightSwitch = document.getElementById("lights-off");
	if(lightSwitch) {
		var lights = amy_getCookie("lights");
		switch(lights) {
		case "off":
			lightSwitch.checked = true;
			break;
		case "on":
			lightSwitch.checked = false;
			break;
		default:
			lightSwitch.checked = false;
			amy_setCookie("lights", "on");
			break;
		}
		amy_addEventListener(lightSwitch, "change", amy_onLightsChanged);
	}
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
