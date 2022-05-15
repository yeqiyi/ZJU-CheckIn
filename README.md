# ZJU每日健（zha）康（gen）打卡程序
打卡不能自选打卡地点成为了润的一大障碍，为此写了使用fake location打卡的东东

-----

## 食用方法

### 获取Cookie
在输入账户名和密码[登录](https://zjuam.zju.edu.cn/cas/login?service=https%3A%2F%2Fhealthreport.zju.edu.cn%2Fa_zju%2Fapi%2Fsso%2Findex%3Fredirect%3Dhttps%253A%252F%252Fhealthreport.zju.edu.cn%252Fncov%252Fwap%252Fdefault%252Findex)后，打开万能的f12获取你的cookie

![获取cookie](https://s1.ax1x.com/2022/04/19/L08FC8.png)

拿到cookie后将cookie复制保存在根目录中的`.cookie`文件中

![保存位置](https://s1.ax1x.com/2022/04/19/L089Et.png)

### 修改配置
配置文件如下

![config.ini](https://s1.ax1x.com/2022/04/19/L08SHI.png)

`注：配置文件会根据配置的校区所在地读取geoInfos目录下对应的fake定位信息（目前仅支持宁波校区，后续有需要再更新）`

### 添加钉钉机器人（可选）
可添加钉钉机器人用于打卡后的提醒，使用方法如下：
 1. 创建单人群聊（钉钉右上角加号->添加联系人->面对面建群（随便输入一串数字即可））
 2. 添加群机器人（注意保存acces_token和密签secret）
 3. 在`config.ini`中加进保存的token和密签

![获取access_token](https://s1.ax1x.com/2022/04/19/L08P4f.png)

![获取secret](https://s1.ax1x.com/2022/04/19/L08CUP.png)

![配置token](https://s1.ax1x.com/2022/04/19/L08k8S.png)

![签到结果](https://s1.ax1x.com/2022/04/19/L08Agg.png)

### 打卡
配置完成后直接运行checkIn.exe即可打卡

`warning：本程序仅用于防扎根，后续可能会加入自动化（概率咕咕咕），目前还需手动。现疫情形势严重，伪造行程可能会造成严重的不良后果，一切后果由使用者承担，请谨慎使用`

-----
## 更新日志
 - 2022-04-01：鉴于最近疫情比较严重，打卡地杭甬两地反复横跳有概率导致蓝码失效，谨慎使用
 - 2022-05-15：现打卡为了反脚本需要验证码了，虽然还是有办法但由于懒不想搞了，此脚本目前无法使用
