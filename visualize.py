import json
import os
import matplotlib.pyplot as plt

names = []
times = []

for file in os.listdir("final_results"):

    if not file.endswith(".json"):
        continue

    path = os.path.join("final_results", file)

    with open(path) as f:

        for line in f:

            try:

                obj = json.loads(line)

                if "NsPerOp" in obj:

                    names.append(
                        file.replace(".json","")
                    )

                    times.append(
                        obj["NsPerOp"]
                    )

                    break

            except:
                pass

plt.figure(figsize=(18,8))

plt.bar(names, times)

plt.xticks(rotation=70)

plt.ylabel("ns/op")

plt.title("ElasticKV Benchmark Runtime")

plt.tight_layout()

plt.savefig("elastickv_benchmarks.png")

print("generated elastickv_benchmarks.png")
