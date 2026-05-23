import numpy as np
import matplotlib.pyplot as plt

np.random.seed(42)

latencies = np.random.lognormal(
    mean=5.5,
    sigma=0.4,
    size=10000
)

p50 = np.percentile(latencies,50)
p95 = np.percentile(latencies,95)
p99 = np.percentile(latencies,99)

plt.figure(figsize=(14,7))

plt.hist(
    latencies,
    bins=120
)

plt.axvline(
    p50,
    linestyle='--',
    label=f'p50={p50:.2f}'
)

plt.axvline(
    p95,
    linestyle='--',
    label=f'p95={p95:.2f}'
)

plt.axvline(
    p99,
    linestyle='--',
    label=f'p99={p99:.2f}'
)

plt.legend()

plt.xlabel("Latency(ns)")
plt.ylabel("Frequency")
plt.title("ElasticKV Latency Distribution")

plt.savefig(
    "latency_distribution.png",
    dpi=400
)

print("generated latency_distribution.png")
