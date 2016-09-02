# street name

## Installation
`go run main.go`

## Usage
```bash
curl /city
# [
#   "花蓮縣",
#   "台北縣",
#   "新北市"
#   "..."
# ]

curl /city_area/花蓮縣
# [
#   "吉安鄉",
#   "花蓮市",
#   "..."
# ]

curl /street_name/花蓮縣/吉安鄉
# {
#   name: "吉安鄉",
#   zip: 973,
#   street_name: [
#     "文化四街",
#     "一心街",
#     "..."
#   ]
# }
```

```
BenchmarkCity-4                     5000            364521 ns/op
BenchmarkCityArea-4                 1000           1035943 ns/op
BenchmarkStreetName-4               1000           1521675 ns/op
PASS
ok      github.com/nukr/street_name     4.820s
```
