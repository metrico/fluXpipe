import "csv"

data = "
#datatype,string,long,long,string
#group,false,false,false,true
#default,_result,,,
,result,table,value,tag
,,0,10,a
,,0,10,a
,,1,20,b
,,1,20,b
,,2,30,c
,,2,30,c
,,3,40,d
,,3,50,e
"

csv.from(csv: data)  |> filter(fn: (r) => r["value"] >= 40 or  r["value"] <= 10) |> yield(name: "res")
