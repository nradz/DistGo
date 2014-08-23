function problem(rom){

	this.machines = 0;
	this.jobs = 0;
	this.bestSeq = [];
	this.bestValue = 0;
	this.trails = [];
	this.p = 0.75;
	this.initialized = false;


	this.mainFunc = function(data){
		var newData = rom.newUpdate();

		//Initialize the problem 
		if(!this.initialized){
			if(newData != null){
				this.initialized = true;
				this.jobs = costs[0].length;
				this.machines = costs.length;
				this.bestSeq = newData;
				this.bestValue = this.evaluate(this.bestSeq);
				this.trails = this.initTrails();				
			}
		}
		//Normal execution
		else{
			if(newData != null){

				var dataValue = this.evaluate(newData);

				if(dataValue < this.bestValue){
					this.bestSeq = newData;
					this.bestValue = this.evaluate(this.bestSeq);
				}
				
				rom.asyncRequest();
			}

			var currentSeq = this.antSequence();
			currentSeq = this.jobIndexBased(currentSeq);
			currentSeq = this.jobIndexBased(currentSeq);
			currentSeq = this.jobIndexBased(currentSeq);

			this.updateTrails(currentSeq);

			var currentSeqValue = this.evaluate(currentSeq);

			if(currentSeqValue < this.bestValue){
				this.bestSeq = currentSeq;
				this.bestValue = currentSeqValue;

				var result = this.bestSeq.map(function(num){return num.toString();});
				//self.postMessage({'cmd':'log','message':this.bestSeq.length});
				rom.result(result);
			}


		}



	};


	this.evaluate = function(seq){
		var flowSeq = new Array(seq.length);

		for(var i = 0; i < seq.length; i++){
			flowSeq[i] = 0;
		}

		for(var m = 0; m < this.machines; m++){

			for(var j = 0; j< seq.length; j++){

				if(j == 0 && m == 0){
					flowSeq[0] = costs[0][seq[j]];
				}
				else if(j == 0){
					flowSeq[j] = flowSeq[j] + costs[m][seq[j]];
				}
				else{
					flowSeq[j] = Math.max(flowSeq[j-1], flowSeq[j]) + costs[m][seq[j]];
				}
			}
		}

		var total = 0;

		for(var j = 0; j < flowSeq.length; j++){
			total += flowSeq[j];
		}
		
		return total;
	};

	
	this.antSequence = function(){
		var seq = new Array(this.jobs);
		var unscheduled = this.bestSeq.slice();

		

		for(var pos = 0; pos < this.jobs; pos++){
			
			var values = new Array(this.jobs);

			for(var job = 0; job < this.jobs; job++){

				values[job] = this.trails[job].slice(0,pos+1).reduce(function(prev, curr){
					return prev + curr;
				});

			}

			var rand = Math.random();

			if(rand <= 0.4){
				seq[pos] = unscheduled[0];
				unscheduled = unscheduled.slice(1);				
			}
			else if(rand <= 0.8){
				var selected = unscheduled.slice(0,5);
				this.sortBySumWeights(selected,values);
				seq[pos] = selected[0];
				//Delete the job from unscheduled array
				var index = unscheduled.indexOf(selected[0]);
				unscheduled = unscheduled.slice(0,index).concat(unscheduled.slice(index+1));
			}
			else{
				var selected = unscheduled.slice(0,5);
				var res = this.roulette(selected, values);
				seq[pos] = res;
				//Delete the job from unscheduled array
				var index = unscheduled.indexOf(res);
				unscheduled = unscheduled.slice(0,index).concat(unscheduled.slice(index+1));
				
			}

		}

		return seq;
	};


	this.jobIndexBased = function(seq){
		var res = seq.slice();

		for(var elem = 0; elem < seq.length; elem++){
			var index = res.indexOf(seq[elem]);

			var list = [];
			for(var pos = 0; pos < seq.length; pos++){
				actual = res.slice();

				//Delete the current element
				if(index < (seq.length-1)){
					actual = actual.slice(0,index).concat(actual.slice(index+1));
				}
				else{
					actual = actual.slice(0,index);
				}
			

				if(pos == index){} //Nothing to do
				else if(pos == 0){
					actual = [seq[elem]].concat(actual);
					list.push(actual);
				}
				else if(pos == (seq.length-1)){
					actual.push(seq[elem]);
					list.push(actual);
				}
				else{
					actual = actual.slice(0, pos).concat([seq[elem]],
						actual.slice(pos));
					list.push(actual);
				}

			}

			list.push(res);
			res = this.sortByFlowtime(list)[0];
		}

		return res;
	};


	this.initTrails = function(){
		
		//Create the matrix
		var trails = new Array(this.jobs);
		for(var i = 0; i < this.jobs; i++){
			trails[i] = new Array(this.jobs);
		}

		for(var job = 0; job < this.jobs; job++){
			for(var pos = 0; pos < this.jobs; pos++){
				var currentIndex = this.bestSeq.indexOf(job);

				if((Math.abs(currentIndex - pos)+1) <= this.jobs/4){
					trails[job][pos] = 1/this.bestValue;
				}
				else if((Math.abs(currentIndex - pos)+1) <= this.jobs/2){
					trails[job][pos] = 1/(2*this.bestValue);
				}
				else{
					trails[job][pos] = 1/(4*this.bestValue);
				}
			}
		}

		return trails;
	};


	this.updateTrails = function(currentSeq){
		
		for(var job = 0; job < this.jobs; job++){
			for(var pos = 0; pos < this.jobs; pos++){
				
				var diff = Math.sqrt(Math.abs(this.bestSeq.indexOf(job)-pos) + 1);

				if(this.jobs < 40){
					
					if(Math.abs(currentSeq[job]-pos)<=1){
						this.trails[job][pos] = this.p*this.trails[job][pos]+(1/(diff*this.bestValue)); 
					}
					else{
						this.trails[job][pos] = this.p*this.trails[job][pos];
					}
					
				}
				else{
					
					if(Math.abs(currentSeq[job]-pos)<=2){
						this.trails[job][pos] = this.p*this.trails[job][pos]+(1/(diff*this.bestValue)); 
					}
					else{
						this.trails[job][pos] = this.p*this.trails[job][pos];
					}

				}
			}
		}

		var pos = this.trails.slice();
	};


	this.sortByFlowtime = function(list){
		var aux = new Array(list.length);

		for(var i = 0; i < list.length; i++){
			aux[i] = {"seq": list[i],
			 "flowtime":this.evaluate(list[i])};
		}
		
		aux.sort(function(a, b){return a.flowtime - b.flowtime});

		var res = new Array(list.length);

		for(var i = 0; i < list.length; i++){
			res[i] = aux[i].seq;
		}

		return res;
	};


	this.sortBySumWeights = function(selected, values){
		selected.sort(function(a,b){
			return values[b] - values[a];
		});
	};

	this.roulette = function(selected, values){
		var total = 0;

		for(var pos = 0; pos < selected.length; pos++){
			total += values[selected[pos]];
		}

		var norm = new Array(selected.length);
		for(var pos = 0; pos < selected.length; pos++){
			norm[pos] = values[selected[pos]]/total;
		}

		var rand = Math.random();
		var chosen = 0;
		var acum = 0;

		for(chosen = 0; chosen < selected.length; chosen++){			
			acum += norm[chosen];			
			if(rand < acum){
				break;
			}
		}
		
		return selected[chosen];
	};

	rom.asyncRequest();

}