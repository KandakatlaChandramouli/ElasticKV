from graphviz import Digraph

g = Digraph(
    "ElasticKV",
    format="png"
)

g.attr(rankdir="LR")
g.attr(bgcolor="black")
g.attr("node",
       shape="box",
       style="filled",
       color="lightblue",
       fontcolor="black")

systems = [
    "SQL Planner",
    "Query Runtime",
    "HNSW Engine",
    "Vector Search",
    "GPU Runtime",
    "CUDA Bridge",
    "RDMA Layer",
    "Raft Consensus",
    "SSTable",
    "RAG Pipeline",
    "Distributed Runtime",
]

for s in systems:
    g.node(s)

edges = [
    ("SQL Planner", "Query Runtime"),
    ("Query Runtime", "Vector Search"),
    ("Vector Search", "HNSW Engine"),
    ("Vector Search", "GPU Runtime"),
    ("GPU Runtime", "CUDA Bridge"),
    ("Distributed Runtime", "Raft Consensus"),
    ("Distributed Runtime", "RDMA Layer"),
    ("Query Runtime", "SSTable"),
    ("Query Runtime", "RAG Pipeline"),
]

for a, b in edges:
    g.edge(a, b)

g.render("elastickv_architecture")

print("generated elastickv_architecture.png")
