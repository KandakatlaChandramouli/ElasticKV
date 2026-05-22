import json
import os
import re
import matplotlib.pyplot as plt

names = []
times = []

folder = "benchmark_results"

pattern = re.compile(r'([0-9.]+)\s+ns/op')

for file in os.listdir(folder):

    if not file.endswith(".json"):
        continue

    filepath = os.path.join(folder, file)

    with open(filepath) as f:

        for line in f:

            try:

                obj = json.loads(line)

                if "Output" not in obj:
                    continue

                output = obj["Output"]

                match = pattern.search(output)

                if match:

                    value = float(match.group(1))

                    names.append(
                        file.replace(".json", "")
                    )

                    times.append(value)

                    break

            except:
                pass

plt.figure(figsize=(18,8))

plt.bar(names, times)

plt.xticks(rotation=45)

plt.ylabel("ns/op")

plt.title("ElasticKV Benchmark Runtime")

plt.tight_layout()

plt.savefig("elastickv_benchmarks_fixed.png")

print("generated elastickv_benchmarks_fixed.png")
