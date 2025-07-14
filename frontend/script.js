// ------ do not touch
function get(url, done) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) done(xhr.responseText);
    };
    xhr.send(null);
}
// ------