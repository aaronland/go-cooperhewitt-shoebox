var shoebox = shoebox || {};

shoebox.item = (function(){

		var self = {

			'init': function(){

				var el = document.getElementById("item-image");

				if (! el){
					return;
				}

				el.onclick = self.toggle;

				self.scale();

				window.onresize = self.resize;
			},

			'toggle': function(){
				console.log('toggle');
			},

			'resize': function(){
				// am I expanded?

				self.scale();
			},

			'scale': function(){

				var el = document.getElementById("item-image");

				if (! el){
					return false;
				}

				var h = el.height;
				var w = el.width;

				// console.log(w + "," + h);
				
				if (w > h){

					var wh = window.innerHeight;
					wh = wh * .85;
					
					el.setAttribute("height", wh);
					el.setAttribute("data-toggle-state", "scaled");
				}

				else {

					var ww = window.innerWidth;
					ww = ww * .85;
					
					el.setAttribute("width", ww);
					el.setAttribute("data-toggle-state", "scaled");
				}
			},

			'expand': function(){

			}
		};

		return self;
})();
