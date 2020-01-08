# 功能

用户输入一个url，server会生成一个8位的短url，之后用户可用此短url查询到原本的url。

# 方法

对时间和ip端口做hash，将hash值转为8位字符的字符串，与原url一起存入数据库，过期时间默认为60分钟。

- POST  /generate：将url以json格式发送给server，server将会返回转换结果
- GET  /get-url?shortlink=11bsi6ZT：将要查询的短url作为路径参数传给server，server进入database中查询，并返回结果。 