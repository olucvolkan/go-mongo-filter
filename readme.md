# Golang Test Case

## How to access the app?
```
curl --location --request POST 'https://afternoon-eyrie-09851.herokuapp.com/mongo' \
--header 'Content-Type: application/json' \
--data-raw '{
"startDate": "2016-01-26",
"endDate": "2018-02-02",
"minCount": 2700,
"maxCount": 3000
}'| jq
```

```
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100  3699    0  3609  100    90   9860    245 --:--:-- --:--:-- --:--:-- 10106
{
"code": 0,
"msg": "success",
"records": [
{
"key": "wtSjVcpg",
"createdAt": "2016-02-22T11:13:43.165Z",
"totalCount": 2888
},
{
"key": "VAbAYJgn",
"createdAt": "2016-10-24T05:57:07.37Z",
"totalCount": 2971
},
{
"key": "XCiSazeS",
"createdAt": "2016-12-13T18:58:33.864Z",
"totalCount": 2906
},
{
"key": "UlDFSFPv",
"createdAt": "2016-03-02T09:30:11.209Z",
"totalCount": 2735
},
{
"key": "UYOsBBSI",
"createdAt": "2016-02-14T15:31:29.518Z",
"totalCount": 2948
},
{
"key": "HYiwsPjw",
"createdAt": "2016-04-05T17:07:46.062Z",
"totalCount": 2977
},
{
"key": "LRgJxDop",
"createdAt": "2016-06-10T21:40:36.359Z",
"totalCount": 2863
},
{
"key": "kzSqsBrJ",
"createdAt": "2016-12-02T15:07:30.465Z",
"totalCount": 2803
},
{
"key": "HJGWkdmD",
"createdAt": "2016-06-08T13:28:10.965Z",
"totalCount": 2718
},
{
"key": "xwqjExMK",
"createdAt": "2016-03-27T09:36:31.788Z",
"totalCount": 2783
},
{
"key": "eaeVCokN",
"createdAt": "2016-07-14T01:45:13.255Z",
"totalCount": 2920
},
{
"key": "NOdGNUDn",
"createdAt": "2016-01-28T07:10:33.558Z",
"totalCount": 2813
},
{
"key": "bxoQiSKL",
"createdAt": "2016-01-29T01:59:53.494Z",
"totalCount": 2991
},
{
"key": "ibfRLaFT",
"createdAt": "2016-12-25T16:43:27.909Z",
"totalCount": 2892
},
{
"key": "mVHGbEap",
"createdAt": "2016-04-11T03:16:28.581Z",
"totalCount": 2853
},
{
"key": "udZfCkvB",
"createdAt": "2016-05-15T00:36:34.126Z",
"totalCount": 2701
},
{
"key": "MhXsNtaT",
"createdAt": "2016-04-17T01:06:48.972Z",
"totalCount": 2942
},
{
"key": "rtKEKONd",
"createdAt": "2016-11-04T17:16:08.542Z",
"totalCount": 2723
},
{
"key": "brcWXLkc",
"createdAt": "2016-08-22T09:16:07.431Z",
"totalCount": 2780
},
{
"key": "zZBGKskQ",
"createdAt": "2016-06-30T01:39:35.456Z",
"totalCount": 2730
},
{
"key": "DIalotCF",
"createdAt": "2016-06-06T02:00:18.387Z",
"totalCount": 2744
},
{
"key": "uLQICSdH",
"createdAt": "2016-11-23T23:55:44.449Z",
"totalCount": 2716
},
{
"key": "UFYsJHDM",
"createdAt": "2016-09-06T04:08:44.393Z",
"totalCount": 2951
},
{
"key": "plSuXweN",
"createdAt": "2016-10-25T11:36:28.069Z",
"totalCount": 2970
},
{
"key": "KYKAKxDr",
"createdAt": "2016-11-27T00:30:34.725Z",
"totalCount": 2713
},
{
"key": "jOjBYTLV",
"createdAt": "2016-11-13T19:54:23.677Z",
"totalCount": 2954
},
{
"key": "nfzwhnJm",
"createdAt": "2016-05-15T23:21:00.153Z",
"totalCount": 2719
},
{
"key": "fEWwrjug",
"createdAt": "2016-10-30T22:49:27.236Z",
"totalCount": 2935
},
{
"key": "rwghjfLQ",
"createdAt": "2016-03-17T11:07:46.355Z",
"totalCount": 2907
},
{
"key": "yoDNIfdV",
"createdAt": "2016-09-02T22:47:37.049Z",
"totalCount": 2780
},
{
"key": "IAuxBQIS",
"createdAt": "2016-11-25T17:55:12.341Z",
"totalCount": 2769
},
{
"key": "ohsXfpHi",
"createdAt": "2016-03-30T20:01:01.948Z",
"totalCount": 2894
},
{
"key": "HmsYvNTB",
"createdAt": "2016-06-12T21:50:44.088Z",
"totalCount": 2917
},
{
"key": "buCwWkpp",
"createdAt": "2016-04-19T11:00:36.397Z",
"totalCount": 2731
},
{
"key": "gtOhweII",
"createdAt": "2016-09-27T16:51:55.223Z",
"totalCount": 2878
},
{
"key": "dcJUSDLR",
"createdAt": "2016-02-27T16:12:30.813Z",
"totalCount": 2780
},
{
"key": "kkxEdhft",
"createdAt": "2016-02-19T06:35:39.409Z",
"totalCount": 2980
},
{
"key": "dNzXijip",
"createdAt": "2016-10-05T21:39:15.288Z",
"totalCount": 2963
},
{
"key": "tyqnxHZh",
"createdAt": "2016-11-25T14:14:46.048Z",
"totalCount": 2872
},
{
"key": "cUZMtDFd",
"createdAt": "2016-08-22T07:54:11.729Z",
"totalCount": 2759
},
{
"key": "vZZOIiPi",
"createdAt": "2016-03-02T09:30:26.664Z",
"totalCount": 2701
},
{
"key": "YUhMrgmc",
"createdAt": "2016-09-25T09:55:20.813Z",
"totalCount": 2862
},
{
"key": "kxMfldnX",
"createdAt": "2016-09-21T04:41:32.835Z",
"totalCount": 2971
},
{
"key": "wIFZewQA",
"createdAt": "2016-03-18T23:32:55.236Z",
"totalCount": 2863
},
{
"key": "plaqeWiK",
"createdAt": "2016-11-20T07:45:28.618Z",
"totalCount": 2773
},
{
"key": "wPpaIkGA",
"createdAt": "2016-06-04T08:07:22.109Z",
"totalCount": 2987
},
{
"key": "pxClAvll",
"createdAt": "2016-12-19T10:00:40.05Z",
"totalCount": 2772
}
]
}
```



```
curl --location --request POST 'https://afternoon-eyrie-09851.herokuapp.com/in-memory/' \
--header 'Content-Type: application/json' \
--data-raw '{
        "key": "active-tabs",
        "value": "getir"
}' | jq
```

```

```

```
curl --location --request GET 'https://afternoon-eyrie-09851.herokuapp.com/in-memory?key=active-tab'| jq
```




## How to run locally
```
 ./start.sh
```

## How to test locally
```
 go test -v
```