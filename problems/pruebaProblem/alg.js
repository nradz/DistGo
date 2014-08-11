function problem(romero){
			this.mainFunc = function(data){
				romero.result(["6"]);
				romero.request();
				var up = romero.newUpdate();
				self.postMessage({'cmd':'log','message':up});
				romero.finish();
				}
			}
