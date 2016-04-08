var shoebox = shoebox || {};

shoebox.index = (function(){

		var self = {
			
			'init': function(class_name){
				
				var els = document.getElementsByClassName(class_name);
				var count = els.length;

				for (var i=0; i < count; i++){
					self.init_thumb(els[i]);
				}
			},

			'init_thumb': function(el){

				el.onmouseover = self.onmouseover;
				el.onmouseout = self.onmouseout;				
			},

			'onmouseover': function(ev){

				var img = ev.target;
				self.toggle_image(img);
			},

			// grrrrrrrnnnnnnnnnnnn.... (20160324/thisisaaronland)

			'onmouseout': function(ev){

				var el = ev.target;

				if (el.nodeName != 'DIV'){
					return;
				}

				var imgs = el.getElementsByTagName("img");
				var img = imgs[0];

				self.toggle_image(img);
			},

			'toggle_image': function(img){

				var src = img.getAttribute("src");
				var alt = img.getAttribute("data-alt-src");

				if ((! src) || (! alt)){
					return;
				}

				img.setAttribute("src", alt);
				img.setAttribute("data-alt-src", src);
			}
		};

		return self;
})();