import matplotlib.pyplot as plt

k = [1,5,10,20,50,100]
recall = [0.61,0.72,0.81,0.88,0.94,0.98]

plt.figure(figsize=(12,6))

plt.plot(
    k,
    recall,
    marker='o',
    linewidth=3
)

plt.xlabel("Top-K")
plt.ylabel("Recall")
plt.title("ElasticKV HNSW Recall Curve")

plt.grid(True)

plt.savefig(
    "recall_curve.png",
    dpi=400
)

print("generated recall_curve.png")
