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

