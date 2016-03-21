var shoebox = shoebox || {};

shoebox.index = (function(){

		var self = {
			
			'init': function(){

				
				var els = document.getElementsByClassName("shoebox-thumb");
				var count = els.length;

				for (var i=0; i < count; i++){

					self.init_thumb(els[i]);
				}
			},

			'init_thumb': function(el){

				el.onmouseover = self.onmouse;
				el.onmouseout = self.onmouse;
				
			},

			'onmouse': function(e){

				var el = e.target;
				var src = el.getAttribute("src");
				var alt = el.getAttribute("data-alt-src");

				el.setAttribute("src", alt);
				el.setAttribute("data-alt-src", src);
			},
		};

		return self;
})();