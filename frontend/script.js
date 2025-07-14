// helper functions
function get(url, token, done) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.setRequestHeader('X-Custom-Token', token || 'default'); // send simple string
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) done(xhr.responseText);
    };
    xhr.send(null);
}

function getByClass(c) {
    var all = document.getElementsByTagName('*'), out = [];
    for (var i = 0; i < all.length; i++) {
        if ((' ' + all[i].className + ' ').indexOf(' ' + c + ' ') > -1)
        out.push(all[i]);
    }
    return out;
}

function timeSince(t) {
    var secs = Math.floor((Date.now() + 3000) / 1000) - t; // +3 seconds because kindles are weird
    var h = Math.floor(secs / 3600),
        m = Math.floor((secs % 3600) / 60),
        s = secs % 60;
    return [h, m, s].map(function(v) {
        return v < 10 ? '0' + v : v;
    }).join(':');
}

// webkit 534 is WEIRD. Date() is UTC-only (on my kindle at least)
function getLocalTimeString(offsetHours) {
    var d = new Date(Date.now() + 3000); // +3 seconds because kindles are weird (again)
    var h = (d.getHours() + offsetHours) + 24 % 24;
    var m = d.getMinutes();
    var s = d.getSeconds();

    return [h, m, s].map(function (v) {
        return v < 10 ? '0' + v : v;
    }).join(':');
}

function elementPing(ip, el) {
    get('/ping', ip, function(res) {
        el.textContent = Number(res) + '% dropped (0s ago)';
        var count = 0;

        if (el.pingInterval) clearInterval(el.pingInterval)

        el.pingInterval = setInterval(function() {
            count++;
            el.textContent = Number(res) + '% dropped (' + count + 's ago)';
        }, 1000)
    });
}

// initial value fill & clock
try {
    setInterval(function() {
        document.getElementById('clock').textContent = getLocalTimeString(-7);
    }, 1000)

    getByClass('right').forEach(function(el) {
        var req = el.textContent.split('@');

        switch(req[0]) {
            case 'hostname':
                get('/hostname', null, function(res) {
                    el.textContent = res;
                });
                break;
            case 'uptime':
                get('/boottime', null, function(res) {
                    setInterval(function() {
                        el.innerText = timeSince(Number(res));
                    }, 1000);
                });
                break;
            case 'ping':
                elementPing(req[1], el)
                setInterval(function() {
                    elementPing(req[1], el)
                }, Number(req[2]) * 1000);
                break;
        }
    });
} catch(e) {
    document.getElementById('err').textContent = e;
}