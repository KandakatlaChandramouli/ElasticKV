import os
import re
import json
import pandas as pd
import plotly.express as px

rows = []

pattern = re.compile(r'([0-9.]+)\s+ns/op')

for root, _, files in os.walk("results"):

    for file in files:

        if not file.endswith(".json"):
            continue

        path = os.path.join(root, file)

        category = root.split("/")[-1]

        with open(path) as f:

            for line in f:

                try:

                    obj = json.loads(line)

                    if "Output" not in obj:
                        continue

                    output = obj["Output"]

                    match = pattern.search(output)

                    if match:

                        rows.append({
                            "category": category,
                            "benchmark": file.replace(".json",""),
                            "ns_per_op": float(match.group(1)),
                        })

                        break

                except:
                    pass

df = pd.DataFrame(rows)

print(df)

fig = px.bar(
    df,
    x="benchmark",
    y="ns_per_op",
    color="category",
    log_y=True,
    template="plotly_dark",
    title="ElasticKV Research Benchmark Suite"
)

fig.update_layout(
    width=2000,
    height=1000,
    font=dict(size=18)
)

fig.write_image(
    "research_dashboard.png",
    scale=4
)

fig.write_html(
    "research_dashboard.html"
)

print("generated research_dashboard.png")
