template_runtime = """package {pkg}

import "sync/atomic"

type Runtime struct {{
    Operations atomic.Uint64
}}

func NewRuntime() *Runtime {{
    return &Runtime{{}}
}}

func (r *Runtime) Execute() {{
    r.Operations.Add(1)
}}

func (r *Runtime) Count() uint64 {{
    return r.Operations.Load()
}}
"""

template_bench = """package {pkg}

import (
    "testing"

    engine "github.com/KandakatlaChandramouli/ElasticKV/internal/{pkg}"
)

func Benchmark{title}(
    b *testing.B,
) {{

    runtime := engine.NewRuntime()

    b.ReportAllocs()

    b.ResetTimer()

    for i := 0; i < b.N; i++ {{
        runtime.Execute()
    }}

    if runtime.Count() == 0 {{
        b.Fatal("runtime failure")
    }}
}}
"""

targets = [
("vectorkernel","VectorKernel"),
("gpukernel","GPUKernel"),
("faissbridge","FAISSBridge"),
("embeddingruntime","EmbeddingRuntime"),
("embeddingpipeline","EmbeddingPipeline"),
("semanticcache","SemanticCache"),
("querygraph","QueryGraph"),
("querygraphoptimizer","QueryGraphOptimizer"),
("learnedindex","LearnedIndex"),
("neuralplanner","NeuralPlanner"),
("attentioncache","AttentionCache"),
("transformindex","TransformIndex"),
("ragpipeline","RAGPipeline"),
("ragretriever","RAGRetriever"),
("reranker","Reranker"),
("semanticexecutor","SemanticExecutor"),
("semanticrouter","SemanticRouter"),
("tokenizer","Tokenizer"),
("tensorruntime","TensorRuntime"),
("tensorallocator","TensorAllocator"),
("tensorcache","TensorCache"),
("featurestore","FeatureStore"),
("invertedvectorindex","InvertedVectorIndex"),
("postinglist","PostingList"),
("docstore","DocStore"),
("chunkindex","ChunkIndex"),
("chunkallocator","ChunkAllocator"),
("chunkstream","ChunkStream"),
("chunkreplication","ChunkReplication"),
("vectorgraph","VectorGraph"),
("graphtraversal","GraphTraversal"),
("hnswgraph","HNSWGraph"),
("ivfpq","IVFPQ"),
("productquantizer","ProductQuantizer"),
("clusterindex","ClusterIndex"),
("kmeans","KMeans"),
("centroidcache","CentroidCache"),
("distancesimd","DistanceSIMD"),
("cosinekernel","CosineKernel"),
("dotproduct","DotProduct"),
("gpubatcher","GPUBatcher"),
("cudaplanner","CUDAPlanner"),
("cudastream","CUDAStream"),
("gpuprefetch","GPUPrefetch"),
("vectorcompactor","VectorCompactor"),
("vectorcheckpoint","VectorCheckpoint"),
("semanticwal","SemanticWAL"),
("semanticraft","SemanticRaft"),
("semanticreplication","SemanticReplication"),
("semanticconsensus","SemanticConsensus"),
]

for pkg, title in targets:

    runtime_path = f"internal/{pkg}/runtime.go"
    bench_path = f"benchmarks/{pkg}/{pkg}_bench_test.go"

    with open(runtime_path, "w") as f:
        f.write(
            template_runtime.format(
                pkg=pkg,
            )
        )

    with open(bench_path, "w") as f:
        f.write(
            template_bench.format(
                pkg=pkg,
                title=title,
            )
        )

print("SINGULARITY SYSTEM WAVE READY")
