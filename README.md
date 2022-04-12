# ZJU每日健（zha）康（gen）打卡程序
打卡不能自选打卡地点成为了润的一大障碍，为此写了使用fake location打卡的东东

-----

## 食用方法

### 获取Cookie
在输入账户名和密码[登录](https://zjuam.zju.edu.cn/cas/login?service=https%3A%2F%2Fhealthreport.zju.edu.cn%2Fa_zju%2Fapi%2Fsso%2Findex%3Fredirect%3Dhttps%253A%252F%252Fhealthreport.zju.edu.cn%252Fncov%252Fwap%252Fdefault%252Findex)后，打开万能的f12获取你的cookie

![获取cookie](images/%E8%8E%B7%E5%8F%96cookie.png)

拿到cookie后将cookie复制保存在根目录中的`.cookie`文件中

![保存位置](images/%E4%BF%9D%E5%AD%98cookie.png)

### 修改配置
配置文件如下

![config.ini](images/%E9%85%8D%E7%BD%AE.png)

`注：配置文件会根据配置的校区所在地读取geoInfos目录下对应的fake定位信息（目前仅支持宁波校区，后续有需要再更新）`

### 打卡
配置完成后直接运行checkIn.exe即可打卡

`warning：本程序仅用于防扎根，后续可能会加入自动化（概率咕咕咕），目前还需手动。现疫情形势严重，伪造行程可能会造成严重的不良后果，一切后果由使用者承担，请谨慎使用`

-----
## 更新日志
 - 2022-04-01：鉴于最近疫情比较严重，打卡地杭甬两地反复横跳有概率导致蓝码失效，谨慎使用
