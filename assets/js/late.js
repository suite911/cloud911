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

function amy_onPageLoadedFull() {
	// amy_onPageLoadedHead();
	if(typeof amy_onPageLoaded === "function") {
		var ok = amy_onPageLoaded();
		if(typeof ok !== "undefined" && !ok) {
			return false;
		}
	}
	// amy_onPageLoadedTail();
	amy_cookieAgree();
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
