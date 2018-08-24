'use strict';

function amy_onChangedFull(event) {
	var captcha = document.getElementById("captcha-onchange");
	if(captcha) {
		grecaptcha.ready(function() {
			grecaptcha.execute(amy_reCAPTCHAv3SiteKey, {action: "change"}).then(function(token) {
				captcha.value = token;
				amy_onChangedFull2();
			});
		});
	} else {
		amy_onChangedFull2();
	}
}

function amy_onChangedFull2() {
	amy_onChangedHead();
	if(typeof amy_onChanged === "function") {
		var ok = amy_onChanged();
		if(typeof ok !== "undefined" && !ok) {
			return false;
		}
	}
	amy_onChangedTail();
}

function amy_onChangedHead() {
}

function amy_onChangedTail() {
}

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
	var submit1 = document.getElementById("submit1");
	if(submit1) {
		amy_addEventListener(window, "submit", amy_onSubmitFull)
		submit1.disabled = false;
		submit1.value = "Submit";
	}
	var ageClassAdult = document.getElementById("age-class-adult");
	if(ageClassAdult) {
		amy_addEventListener(ageClassAdult, "click", amy_onChangedFull)
	}
	var ageClassMinor = document.getElementById("age-class-minor");
	if(ageClassMinor) {
		amy_addEventListener(ageClassMinor, "click", amy_onChangedFull)
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

function amy_onPageLoadedFull(event) {
	var captcha = document.getElementById("captcha-onload");
	if(captcha) {
		grecaptcha.ready(function() {
			grecaptcha.execute(amy_reCAPTCHAv3SiteKey, {action: "load"}).then(function(token) {
				captcha.value = token;
				amy_onPageLoadedFull2();
			});
		});
	} else {
		amy_onPageLoadedFull2();
	}
}

function amy_onPageLoadedFull2() {
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

function amy_onSubmitFull(event) {
	event.preventDefault(); // don't submit the form
	var captcha = document.getElementById("captcha-onsubmit");
	if(captcha) {
		grecaptcha.ready(function() {
			grecaptcha.execute(amy_reCAPTCHAv3SiteKey, {action: "submit"}).then(function(token) {
				captcha.value = token;
				amy_onSubmitFull2();
			});
		});
	} else {
		amy_onSubmitFull2();
	}
}

function amy_onSubmitFull2() {
	var lightSwitch = document.getElementById("lights-off");
	if(lightSwitch) {
		var lightAnalytics = document.getElementById("analytics-lights");
		if(lightAnalytics) {
			if(lightSwitch.checked) {
				lightAnalytics.value = "off";
			} else {
				lightAnalytics.value = "on";
			}
		}
	}
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
