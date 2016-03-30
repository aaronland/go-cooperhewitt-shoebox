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

				mapzen.whosonfirst.yesnofix.enabled(false);
				self.get_index();
			},

			'toggle': function(){
				self.feedback('toggle');
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

				// self.feedback(w + "," + h);

				var wh = window.innerHeight;
				wh = wh * .85;

				el.setAttribute("height", wh);
				el.setAttribute("data-toggle-state", "scaled");
			},

			'expand': function(){

			},

			'get_index': function(){
				
				var on_success = function(rsp){
					var isa = rsp['refers_to_a'];
					var uid = rsp['refers_to_uid'];

					if (isa == "object"){
						self.get_object(uid);
					}

					else {
						self.feedback("I don't know how to " + isa);
					}

					mapzen.whosonfirst.yesnofix.makeitso(rsp, "details-item");
				};

				var on_error = function(){
					self.feedback("error");
				};
				
				self.xhr('index.json', on_success, on_error);
			},

			'get_object': function(id){

				var url = id + ".json";

				var on_success = function(rsp){

					var object = rsp['object'];
					mapzen.whosonfirst.yesnofix.makeitso(object, "details-refers-to");					
				};

				var on_error = function(){
					self.feedback("error");
				};
				
				self.xhr(url, on_success, on_error);
			},
			
			'xhr': function(url, on_success, on_error){
			
				var req = new XMLHttpRequest();

				req.onload = function(){

					try {
						var rsp = JSON.parse(this.responseText);
					}

					catch (e){
						self.feedback("failed to parse " + url + ", because " + e);

						if (on_error){
							on_error();
						}

						return false;
					}

					if (on_success){
						on_success(rsp);
					}
				};

				try {
					req.open("get", url, true);
					req.send();
				}

				catch(e){
					self.feedback("failed to fetch " + url + ", because ");
					self.feedback(e);

					if (on_error){
						on_fail();
					}
				}
			},

			'feedback': function(msg){
				console.log(msg);
			}
		};

		return self;
})();
