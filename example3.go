package main

import (
	"os"
	"fmt"
	"time"
	"github.com/DataIntelligenceCrew/go-faiss"
)

const (
	DIM = 10
	NRESULT = 5
	INDEX_FILE = "/tmp/faiss.idx"
	SAMPLE_INDEX = 7000
	NSEARCH = 3000
)

func DumpIndexFileSize(index faiss.Index) {
	faiss.WriteIndex(index, INDEX_FILE)
	file, _ := os.Open(INDEX_FILE)
	defer file.Close()
	stat, _ := file.Stat()
	fmt.Printf("Index size: %v bytes\n", stat.Size())
}

func IndexPerformance(description string, efSearch float64, nprobe float64, train bool) {
	fmt.Println("*** ", fmt.Sprintf("%s, efSearch: %v, nprobe: %v", description, efSearch, nprobe))
	searches := *GetData("random.csv", DIM)
	data := *GetData("data.csv", DIM)
	index, _ := faiss.IndexFactory(DIM, description, faiss.MetricInnerProduct)
	parameterSpace, _ := faiss. NewParameterSpace()
	if efSearch > 0 {
		parameterSpace.SetIndexParameter(index, "efSearch", efSearch)
	}
	if nprobe > 0 {
		parameterSpace.SetIndexParameter(index, "nprobe", nprobe)
	}
	if train {
		startAt := time.Now().UnixNano()
		index.Train(data)
		endAt := time.Now().UnixNano()
		fmt.Printf("Train() Elapsed: %v ms\n", (endAt - startAt) / 1000000)
	}
	{
		startAt := time.Now().UnixNano()
		index.Add(data)
		endAt := time.Now().UnixNano()
		fmt.Printf("Add() Elapsed: %v ms\n", (endAt - startAt) / 1000000)
	}
	fmt.Printf("Ntotal: %v\n", index.Ntotal())
	distances, labels, _ := index.Search(searches[SAMPLE_INDEX*DIM:(SAMPLE_INDEX+1)*DIM], NRESULT)
	fmt.Printf("Sample Distances: %v, Labels: %v\n", distances, labels)
	DumpIndexFileSize(index)
	startAt := time.Now().UnixNano()
	for i := 0; i < NSEARCH; i++ {
		index.Search(searches[i*DIM:(i+1)*DIM], NRESULT)
	}
	endAt := time.Now().UnixNano()
	fmt.Printf("Search Elapsed: %v ms\n", (endAt - startAt) / 1000000)
}

func main() {
	IndexPerformance("Flat", 0, 0, false)
	IndexPerformance("PQ2x8", 0, 0, true)
	IndexPerformance("PQ2x8,RFlat", 0, 0, true)
	IndexPerformance("PQ5x8", 0, 0, true)
	IndexPerformance("OPQ2,PQ2x8", 0, 0, true)
	IndexPerformance("OPQ5,PQ5x8", 0, 0, true)
	IndexPerformance("SQ8", 0, 0, true)
	IndexPerformance("HNSW2", 2, 0, false)
	IndexPerformance("HNSW4", 2, 0, false)
	IndexPerformance("HNSW8", 2, 0, false)
	IndexPerformance("HNSW8_PQ5x8", 2, 0, true)
	IndexPerformance("OPQ5,HNSW8_PQ5x8", 2, 0, true)
	IndexPerformance("IVF100,Flat", 0, 1, true)
	IndexPerformance("IVF100,Flat", 0, 2, true)
	IndexPerformance("IVF100,Flat", 0, 100, true)
	IndexPerformance("OPQ5,IVF100,PQ5x8", 0, 1, true)
	IndexPerformance("OPQ5,IVF100,PQ5x8", 0, 2, true)

	IndexPerformance("OPQ1_2,PQ2x8", 0, 0, true)
	IndexPerformance("OPQ1_5,PQ5x8", 0, 0, true)
	IndexPerformance("OPQ1_5,HNSW8_PQ5x8", 2, 0, true)
	IndexPerformance("OPQ1_5,IVF100,PQ5x8", 0, 1, true)
	IndexPerformance("OPQ1_5,IVF100,PQ5x8", 0, 2, true)

	IndexPerformance("OPQ1_3,PQ3x8", 0, 0, true)
	IndexPerformance("OPQ1_3,HNSW8_PQ3x8", 2, 0, true)
	IndexPerformance("OPQ1_3,IVF100,PQ3x8", 0, 1, true)
	IndexPerformance("OPQ1_3,IVF100,PQ3x8", 0, 2, true)
}
