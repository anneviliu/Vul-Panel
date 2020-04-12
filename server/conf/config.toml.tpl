[base]
Port=":8888"
BaseURL="http://xray/"
WebRoot="./"

[mysql]
DBHost="localhost:33060"
DBUsername="root"
DBPassword="root"
DBName="Xray"

[wechat]
CorpID=""
CorpSecret=""
AccessTokenUrl="https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
SendMsgUrl="https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
AgentID=

[admin]
InviteCode = ""