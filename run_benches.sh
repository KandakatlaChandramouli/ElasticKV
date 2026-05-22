#!/bin/bash

export PATH=$PATH:/usr/local/go/bin
export GOMAXPROCS=$(nproc)

runbench () {

    NAME=$1
    PKG=$2

    echo "================================="
    echo "RUNNING: $NAME"
    echo "================================="

    go test \
        -run=^$ \
        -bench=. \
        -benchtime=1x \
        -json \
        $PKG \
        > benchmark_results/$NAME.json 2>&1

    echo "================================="
    echo "DONE: $NAME"
    echo "================================="
}

runbench hnsw "./benchmarks/hnsw/..."
runbench raft "./benchmarks/raft/..."
runbench rdma "./benchmarks/rdma/..."
runbench querysql "./benchmarks/querysql/..."
runbench lsm "./benchmarks/lsm/..."
runbench execution "./benchmarks/execution/..."
runbench vectorsearch "./benchmarks/vectorsearch/..."
runbench multigpu "./benchmarks/multigpu/..."
runbench rag "./benchmarks/ragpipeline/..."
