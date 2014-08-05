package perFlowACOProblem

import(
	//"sort"
)

//Order the jobs by non-increasing sums of
//processing times on the machines
type byCosts struct{
	seq []int
	costs []int
}

func (bc *byCosts) Len() int{
	return len(bc.seq)
}

func (bc *byCosts) Swap(i,j int){
	bc.seq[i], bc.seq[j] = bc.seq[j], bc.seq[i]
	bc.costs[i], bc.costs[j] = bc.costs[j], bc.costs[i]
}

func (bc *byCosts) Less(i, j int) bool{
	return bc.costs[i] < bc.costs[j]
}

//Order sequences by total flowtime
type byTotalFlowtime struct{
	list [][]int
	flowtimes []int
}

func (bt *byTotalFlowtime) Len() int{
	return len(bt.list)
}

func (bt *byTotalFlowtime) Swap(i,j int){
	bt.list[i], bt.list[j] = bt.list[j], bt.list[i]
	bt.flowtimes[i], bt.flowtimes[j] = bt.flowtimes[j], bt.flowtimes[i]
}

func (bt *byTotalFlowtime) Less(i, j int) bool{
	return bt.flowtimes[i] < bt.flowtimes[j]
}
