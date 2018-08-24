'use strict';

function amy_onDOMReadyFull() {
	amy_onDOMReadyHead();
	if(typeof amy_onDOMReady === "function") {
		var ok = amy_onDOMReady();
		if(typeof ok !== "undefined" && !ok) {
			return false;
		}
	}
	amy_onDOMReadyTail();
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

function amy_onPageLoadedFull() {
	amy_onPageLoadedHead();
	if(typeof amy_onPageLoaded === "function") {
		var ok = amy_onPageLoaded();
		if(typeof ok !== "undefined" && !ok) {
			return false;
		}
	}
	amy_onPageLoadedTail();
}

function amy_onPageLoadedHead() {
}

function amy_onPageLoadedTail() {
	if(typeof amy_cookieAgree === "function") {
		amy_cookieAgree();
	}
}

if (typeof document.addEventListener === "function") {
	document.addEventListener("DOMContentLoaded", amy_onDOMReadyFull, false);
} else if (typeof document.attachEvent === "function") {
	document.attachEvent("onreadystatechange", amy_onDOMReadyFull);
} else {
	window.onload = amy_onDOMReadyFull;
}

if (typeof window.addEventListener === "function") {
	window.addEventListener("load", amy_onPageLoadedFull, false);
} else if (typeof window.attachEvent === "function") {
	window.attachEvent("onload", amy_onPageLoadedFull);
} else {
	window.onload = amy_onPageLoadedFull;
}
