function mainFunc(romero,data){

	var size = graph.length,
	newData = undefined,
	poblacionActual = [],
	mejores = [],
	aleatorios = [],
	nuevaPoblacion = [],
	contador = 0;

	romero.asyncRequest();

	poblacionActual = iniciaPoblacion(size,40);

	while(true){
		//Check if there is new data from server.
		newData = romero.newData();
		if(typeof newData != "undefined"){

			if(newData == "finished"){break;}
			else{poblacion.push(newData.data);}
			
			//Prepare a new request
			romero.asyncRequest();
		}

		mejores = seleccionaMejores(poblacionActual, 5);
		aleatorios = seleccionaAleatorios(poblacionActual, 5);
		nuevaPoblacion = cruza(mejores.concat(aleatorios), 20);
		nuevaPoblacion = muta(nuevaPoblacion, 0.1);

		poblacionActual = seleccionaMejores(poblacionActual.concat(nuevaPoblacion), 40);


		//Cada cinco iteraciones, manda el mejor cromosoma.
		contador +=1;

		if(contador == 5){
			contador = 0;

			romero.asyncRequest(seleccionaMejores(poblacionActual, 1));
		}
		
	}
	
	romero.finish();
}



function iniciaPoblacion(size, muestra){
	var poblacion = [];
	var primerCromosoma = [];
	var auxCromosoma = [];
	var auxFitness = 0;

	//genera primer cromosoma
	for(var i = 0; i<size; i++){
		primerCromosoma.push(i);
	}

	//Genera tantos cromosomas como indique el parametro muestra
	for(var i = 0; i<muestra; i++){
		auxCromosoma = shuffle(primerCromosoma);
		auxFitness = fitness(auxCromosoma);

		poblacion.push({"cromosoma":auxCromosoma, "fitness":auxFitness});
	}

	return poblacion;
}


function seleccionaMejores(poblacion, muestra){

	poblacion.sort(function(a,b){b.fitness - a.fitness});

	return poblacion.slice(0,muestra);

}

function seleccionaAleatorios(poblacion, muestra){
	
	var auxPoblacion = poblacion.slice(0);
	var res = [];
	var temporal = undefined;
	var randomIndex = 0;

	for(var i = 0; i < muestra; i++){
		randomIndex = Math.floor(Math.random() * auxPoblacion.length);
		temporal = auxPoblacion.splice(randomIndex,1);
		res.push(temporal[0]);
	}

	return res;
}

//Cruce basado en orden
function cruza(poblacion, numHijos){

	var res = [];
	var iPadre1 = 0;
	var iPadre2 = 0;
	var padre1 = [];
	var padre2 = [];
	var primerPunto = 0;
	var segundoPunto = 0;
	var auxCruza = undefined;
	

	for(var i = 0; i<numHijos; i++){
		iPadre1 = Math.floor(Math.random() * poblacion.length);

		//Para asegurar que ambos padres no son el mismo
		while(true){
			iPadre2 = Math.floor(Math.random() * poblacion.length);
			if(iPadre1 != iPadre2){
				break;
			}
		}

		padre1 = poblacion[iPadre1].cromosoma;
		padre2 = poblacion[iPadre2].cromosoma;

		auxCruza = function(p1, p2){

			var hijo = [];
			var auxFitness = 0;

			primerPunto = Math.floor(Math.random() * p1.length);

			//asegurar que segundo punto es mayor
			segundoPunto = Math.floor(Math.random() *
			p1.length-primerPunto)+primerPunto+1;

			hijo = p1.slice(primerPunto,segundoPunto);

			for(var e = 0; e<p2.length; e++){
				if(hijo.indexOf(p2[e]) == (-1)){
					hijo.push(p2[e])
				}
			}

			auxFitness = fitness(hijo);
			return {"cromosoma":hijo, "fitness":auxFitness};
			
		}

		res.push(auxCruza(padre1,padre2));
		res.push(auxCruza(padre2,padre1));

	}

	return res;

}


function muta(poblacion, probabilidad){

	var res = [];
	var randomIndex = undefined;
	var crom = [];
	var auxFitness = 0;
	var aux1 = 0;
	var aux2 = 0;

	for(var i = 0; i<poblacion.length; i++){
		
		crom = poblacion[i].cromosoma;
		
		for(var e = 0; e<crom.length; e++){
			if(Math.random()<probabilidad){
				rIndex = Math.floor(Math.random() 
					* (crom.length-e-1)) + e;
				aux1 = crom[e];
				aux2 = crom[rIndex];
				crom[e] = aux2;
				crom[rIndex] = aux1;
			}		
		}
		auxFitness = fitness(crom);
		res.push({"cromosoma":crom, "fitness":auxFitness});
	}

	return res;
}

function fitness(cromosoma){
	var valor = 0;

	for(var i = 0; i < cromosoma.length - 1; i++){
		valor += graph[cromosoma[i]][cromosoma[i+1]]
	}

	return valor;
}

function shuffle(array) {
	var currentIndex = array.length
    , temporaryValue
    , randomIndex,
    auxArray = array.slice(0);
    ;

  // While there remain elements to shuffle...
  while (0 !== currentIndex) {

    // Pick a remaining element...
    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex -= 1;

    // And swap it with the current element.
    temporaryValue = auxArray[currentIndex];
    auxArray[currentIndex] = auxArray[randomIndex];
    auxArray[randomIndex] = temporaryValue;
  }

  return auxArray;
}
