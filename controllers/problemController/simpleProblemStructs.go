package problemController

import(
"github.com/nradz/DistGo/channels"
)

type data interface{}

type simpleProblemState struct{
	Alg string //The actual algorithm that is executing in the clients
	LastUpdate data //The last update available
	Clients map[uint32]*clientState //Map of the clients with their state
}

type clientState struct{
	Ready bool //A bool variable that indicate if the client is ready for a update
	Updated bool //It indicate if the client have received the last update
	ResChan chan channels.ProblemControlResponse /*The channel where the problemController
	receive the request*/
}