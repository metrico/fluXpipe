import "sampledata"

sampledata.int()
    |> group(columns: ["_time", "tag"])
