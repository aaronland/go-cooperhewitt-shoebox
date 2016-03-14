var shoebox = shoebox || {};

shoebox.item = (function(){

		var self = {

			'init': function(){

				var el = document.getElementById("item-image");

				if (! el){
					return;
				}

				el.onclick = self.toggle;

				self.collapse();
			},

			'toggle': function(){
				console.log('toggle');
			},

			'collapse': function(){

				var el = document.getElementById("item-image");

				if (! el){
					return false;
				}

				var h = el.height;
				var w = el.width;

				// console.log(w + "," + h);

				var wh = window.innerHeight;
				wh = wh * .85;

				el.setAttribute("height", wh);
				el.setAttribute("data-toggle-state", "collapse");
			},

			'expand': function(){

			}
		};

		return self;
})();