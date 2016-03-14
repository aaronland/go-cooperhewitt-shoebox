var shoebox = shoebox || {};

shoebox.common = (function(){

		var self = {

			'init': function(){
				window.onkeyup = self.onkeyup;
			},

			'onkeyup': function(e){

				var key = e.keyCode || e.which;
				var keychar = String.fromCharCode(key);

				if (key == 39){
					self.goto_next();
				}

				if (key == 37){
					self.goto_previous();
				}
			},

			'goto_next': function(){

				// please to check next/prev rel link
				var el = document.getElementById("rel-next");

				if (! el){
					return;
				}

				var href = el.getAttribute("href");
				location.href = href;
			},

			'goto_previous': function(){

				// please to check next/prev rel link
				var el = document.getElementById("rel-prev");

				if (! el){
					return;
				}

				var href = el.getAttribute("href");
				location.href = href;
			}
		};

		return self;
})();