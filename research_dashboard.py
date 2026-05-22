import os
import re
import json
import pandas as pd
import plotly.express as px

benchmarks = []
times = []

folder = "benchmark_results"

pattern = re.compile(r'([0-9.]+)\s+ns/op')

for file in os.listdir(folder):

    if not file.endswith(".json"):
        continue

    path = os.path.join(folder, file)

    with open(path) as f:

        for line in f:

            try:

                obj = json.loads(line)

                if "Output" not in obj:
                    continue

                output = obj["Output"]

                match = pattern.search(output)

                if match:

                    benchmarks.append(
                        file.replace(".json","")
                    )

                    times.append(
                        float(match.group(1))
                    )

                    break

            except:
                pass

df = pd.DataFrame({
    "Subsystem": benchmarks,
    "Latency(ns/op)": times
})

fig = px.bar(
    df,
    x="Subsystem",
    y="Latency(ns/op)",
    color="Latency(ns/op)",
    title="ElasticKV Research Benchmark Dashboard",
    template="plotly_dark",
)

fig.update_layout(
    width=1600,
    height=900,
    font=dict(size=18),
)

fig.write_image(
    "research_dashboard.png",
    scale=3,
)

fig.write_html(
    "research_dashboard.html"
)

print(df)

print("generated research_dashboard.png")
print("generated research_dashboard.html")
