## web开发简介

golang提供了大量的web库函数，如果项目的路由在个位数、url固定且不通过url传递参数，则官方库函数足够了。但在复杂情况下，仍需框架，如：
````
GET    /api/:id
POST   /api/:id
DELETE /api/:id
···
````
golang web框架大致可分为两种：
````
1. mvc框架
2. Router框架
````