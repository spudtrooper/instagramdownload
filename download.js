setInterval(function() {
  document.querySelectorAll('img[class="FFVAD"]').forEach(function(el) {
    console.log(el.getAttribute('alt'));
    let srcset = el.getAttribute('srcset');
    srcset
      .split(',')
      .forEach(function(e) {
	let p = e.split(' ');
	console.log(p[1] + ' ' + p[0]);
      });
  });
  window.scrollTo(0,document.body.scrollHeight);
}, 3000);
