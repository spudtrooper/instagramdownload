function findParts() {
    let div = Array.from(document.getElementsByTagName('div')).filter((el) => { 
        return el.style && el.style.maxWidth && el.style.maxWidth == '815px' 
    })[0];
    let time = Array.from(div.parentNode.parentNode.parentNode.getElementsByTagName('time'))[0];
    let datetime = time.getAttribute('datetime');
    let img = null;
    document.querySelectorAll('img[class="FFVAD"]').forEach(function (el) {
        let srcset = el.getAttribute('srcset');
        srcset.split(',').forEach(function (e) {
            let p = e.split(' ');
            if (!img && p[1] == '1080w') {
                img = p[0];
            }
        });
    });
    console.log('img: ' + img + ' datetime: ' + datetime);
}

function findAllPosts() {
    return Array.from(document.querySelectorAll('img[class="FFVAD"]'));
}

function iteration() {
    let posts = findAllPosts();
    let loop = () => {
        if (!posts.length) {
            window.scrollTo(0, document.body.scrollHeight);
            setTimeout(() => {
                iteration();
            }, 3000)
            return;
        }
        let post = posts.pop(0);
        post.click();
        setTimeout(() => {
            findParts();
            loop();
        }, 3000);
    };
    loop();
}