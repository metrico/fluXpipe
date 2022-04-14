import "experimental/json"
import "experimental/array"

jsonStr = bytes(v:"{
      \"node\": {
          \"items\": [
              {
                  \"id\": \"15612462\",
                  \"color\": \"red\",
                  \"states\": [
                      {
                          \"name\": \"ready\",
                          \"duration\": 10
                      },
                      {
                          \"name\": \"closed\",
                          \"duration\": 13
                      },
                      {
                          \"name\": \"pending\",
                          \"duration\": 3
                      }
                  ]
              },
              {
                  \"id\": \"15612462\",
                  \"color\": \"blue\",
                  \"states\": [
                      {
                          \"name\": \"ready\",
                          \"duration\": 5
                      },
                      {
                          \"name\": \"closed\",
                          \"duration\": 0
                      },
                      {
                          \"name\": \"pending\",
                          \"duration\": 16
                      }
                  ]
              }
          ]
      }
 }")

data = json.parse(data: jsonStr)

// Map over all items in the JSON extracting
// the id, color and pending duration of each.
// Construct a table from the final records.
  
array.from(rows:
    data.node.items
        |> array.map(fn:(x) => {
            pendingState = x.states
                |> array.filter(fn: (x) => x.name == "pending")
            pendingDur = if length(arr: pendingState) == 1
                then
                    pendingState[0].duration
                else
                    0.0
            return {
                id: x.id,
                color: x.color,
                pendingDuration: pendingDur,
            }
        })
)
