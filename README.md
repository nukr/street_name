# street name

地址來源：http://download.post.gov.tw/post/download/Xml_10510.xml

## Installation
go get


## Usage
```bash
curl http://yourdomain/list
# [
#   "台灣",
#   "韓國",
#   "中國"
#   "..."
# ]

curl http://yourdomain/list -H "Accept-Language: en-us"
# [
#   "台灣",
#   "Korea",
#   "China"
#   "..."
# ]

curl http://yourdomain/list/台灣
# [
#   "花蓮縣",
#   "台北縣",
#   "新北市"
#   "..."
# ]

curl http://yourdomain/list/花蓮縣
# [
#   "吉安鄉",
#   "花蓮市",
#   "..."
# ]

curl http://yourdomain/list/花蓮縣/吉安鄉
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
