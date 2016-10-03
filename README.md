# myagent
agent for any json

go get -u github.com/beego/bee
go get -u github.com/astaxie/beego

# 通过bee api 进行自动编译
进入到你的GOPATH/src目录，执行命令bee api bapi,
进入目录cd bapi,执行命令bee run -downdoc=true -gendoc=true