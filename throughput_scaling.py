import matplotlib.pyplot as plt

nodes = [1,2,4,8,16]
throughput = [1200,2400,4700,8800,16100]

plt.figure(figsize=(12,6))

plt.plot(
    nodes,
    throughput,
    marker='o',
    linewidth=3
)

plt.xlabel("Cluster Nodes")
plt.ylabel("Queries/sec")
plt.title("ElasticKV Distributed Scaling")

plt.grid(True)

plt.savefig(
    "throughput_scaling.png",
    dpi=400
)

print("generated throughput_scaling.png")
