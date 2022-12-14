package routes

import (
	"boutytask/bcos"
	"boutytask/dbs"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

// 通用化响应消息格式
type RespMsg struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//code提前编码
const (
	RESPCODE_OK       = "0"
	RESPCODE_LOGINERR = "1"
	RESPCODE_PARAMERR = "2"
	RESPCODE_BLKERR   = "3"
	RESPCODE_UNKNOWN  = "4"
)

var respMapMsg = map[string]string{
	RESPCODE_OK:       "正常",
	RESPCODE_LOGINERR: "登陆失败",
	RESPCODE_PARAMERR: "参数错误",
	RESPCODE_BLKERR:   "区块链访问错误",
	RESPCODE_UNKNOWN:  "系统正在升级",
}

//统一的响应消息
func CoderspMsg(resp *RespMsg, c *gin.Context) {
	resp.Msg = respMapMsg[resp.Code]
	c.JSON(http.StatusOK, resp)
}

//需要一个请求消息内容的结构体
type UserInfo struct {
	UserName string `json:"username"`
	Passwd   string `json:"password"` // 严格遵循接口
}

func Register(c *gin.Context) {
	//1. 组织响应消息
	resp := RespMsg{
		Code: "0",
	}

	defer CoderspMsg(&resp, c) //当Register退出时，CoderspMsg执行，使用resp信息
	//2. 解析请求消息
	var ui UserInfo
	err := c.ShouldBind(&ui)
	if err != nil {

		resp.Code = RESPCODE_PARAMERR
		log.Panic("failed to ShouldBind", err)
	}
	fmt.Printf("Register:%+v\n", ui)
	//3. 调用区块链合约
	err = bcos.Register(ui.UserName, ui.Passwd)
	if err != nil {
		resp.Code = RESPCODE_BLKERR
		log.Println("failed to bcos.Register", err)
	}
}

//curl  -H "Content-type: application/json" -X POST -d '{"username":"yekai","password":"123"}' "http://localhost:9090/login"

func Login(c *gin.Context) {
	//1. 组织响应消息
	resp := RespMsg{
		Code: "0",
	}

	defer CoderspMsg(&resp, c) //当Login退出时，CoderspMsg执行，使用resp信息
	//2. 解析请求消息
	var ui UserInfo
	err := c.ShouldBind(&ui)
	if err != nil {

		resp.Code = RESPCODE_PARAMERR
		log.Panic("failed to ShouldBind", err)
	}
	fmt.Printf("Login:%+v\n", ui)
	//3. 调用区块链合约
	addr, err := bcos.Login(ui.UserName, ui.Passwd)
	if err != nil {
		resp.Code = RESPCODE_BLKERR
		log.Panic("failed to bcos.Login", err)
	}
	if addr == "" {
		resp.Code = RESPCODE_LOGINERR
		fmt.Println("User or password err")
	}

	//4. session记录
	store := ginsession.FromContext(c)
	store.Set("username", ui.UserName)
	store.Set("password", ui.Passwd)
	store.Set("address", addr)
	err = store.Save()
	if err != nil {
		//c.AbortWithError(500, err)
		resp.Code = RESPCODE_UNKNOWN
		fmt.Println("session err", err)
		return
	}
}

func Issue(c *gin.Context) {
	//1. 组织响应消息
	resp := RespMsg{
		Code: "0",
	}

	defer CoderspMsg(&resp, c)
	//2. 解析请求消息
	var taskinfo bcos.RespTaskInfo
	err := c.ShouldBind(&taskinfo)
	if err != nil {

		resp.Code = RESPCODE_PARAMERR
		log.Panic("failed to ShouldBind", err)
	}
	fmt.Printf("Issue:%+v\n", taskinfo)
	//3. 从session获取用户和密码信息
	store := ginsession.FromContext(c)
	username, ok1 := store.Get("username")
	password, ok2 := store.Get("password")
	if !ok1 || !ok2 {
		resp.Code = RESPCODE_UNKNOWN
		fmt.Println("Failed to get session")
		return
	}
	//4. 调用智能合约
	fmt.Println("user,pass:", username, password)
	txhash, err := bcos.Issue(username.(string), password.(string), taskinfo.Bounty, taskinfo.Desc)
	if err != nil {
		resp.Code = RESPCODE_BLKERR
		log.Panic("failed to bcos.Issue:", err)
	}
	fmt.Println("issue, hash=", txhash)

}

//  /tasklist?page=1
func Tasklist(c *gin.Context) {
	//1. 组织响应消息(data字段使用)
	resp := RespMsg{
		Code: "0",
	}

	defer CoderspMsg(&resp, c)

	//2. 获取前端请求页(任务列表可以有很多，前端请求分页显示，一行显示10条，显示第几页)
	page := c.Query("page")
	ipage, _ := strconv.Atoi(page)
	fmt.Println("Tasklist:page = ", ipage)

	//3. 获取任务列表信息
	tasks, err := bcos.Tasklist()
	if err != nil {
		resp.Code = RESPCODE_BLKERR
		log.Panic("failed to Tasklist", err)
	}

	//4. 响应消息填写
	//需要根据page计算要返回的列表是什么
	if ipage <= 0 {
		ipage = 1
	}
	ibegin := (ipage - 1) * 10
	iend := ipage * 10
	if iend > len(tasks) {
		iend = len(tasks)
	}
	//data:{total:8, data:[]list}
	data := struct {
		Total int         `json:"total"`
		Data  interface{} `json:"data"`
	}{
		Total: len(tasks),
		Data:  tasks[ibegin:iend],
	}

	resp.Data = data

}

//任务更新接口
func Update(c *gin.Context) {
	//1. 组织响应消息
	resp := RespMsg{
		Code: "0",
	}

	defer CoderspMsg(&resp, c)
	//2. 解析请求消息
	var taskinfo bcos.RespTaskInfo
	err := c.ShouldBind(&taskinfo)
	if err != nil {

		resp.Code = RESPCODE_PARAMERR
		log.Panic("failed to ShouldBind", err)
	}
	fmt.Printf("Update:%+v\n", taskinfo)
	//3. 从session获取用户和密码信息
	store := ginsession.FromContext(c)
	username, ok1 := store.Get("username")
	password, ok2 := store.Get("password")
	if !ok1 || !ok2 {
		resp.Code = RESPCODE_UNKNOWN
		fmt.Println("Failed to get session")
		return
	}
	//4. 调用智能合约
	taskID, _ := strconv.Atoi(taskinfo.TaskID)
	fmt.Println(username, password)
	txhash, err := bcos.Update(username.(string), password.(string), taskinfo.Comment, int64(taskID), taskinfo.Status)
	if err != nil {
		resp.Code = RESPCODE_BLKERR
		log.Panic("failed to bcos.Update", err)
	}
	hashData := struct {
		Txhash string `json:"txhash"`
	}{
		txhash,
	}
	resp.Data = hashData

}

type RespBalance struct {
	UserName string `json:"username"`
	Amount   int64  `json:"amount"`
}

//余额查询

func BalanceOf(c *gin.Context) {
	// 1. 响应消息
	resp := RespMsg{
		Code: RESPCODE_OK,
		Data: nil,
	}
	//延迟执行
	defer CoderspMsg(&resp, c)
	// 2. 获取请求消息
	username, ok := c.GetQuery("username")
	if !ok {
		fmt.Println("request must have username")
		resp.Code = RESPCODE_PARAMERR
		return
	}
	addr := dbs.Query([]byte(username))
	fmt.Println("BalanceOf-addr:", string(addr))
	// 3. 调用链码
	value, err := bcos.BalanceOf(string(addr))
	if err != nil {
		resp.Code = RESPCODE_BLKERR
		return
	}
	// 4. 响应数据
	blResp := RespBalance{username, value.Int64()}
	resp.Data = blResp
}

// 退出
func Logout(c *gin.Context) {
	// 响应消息
	resp := RespMsg{
		Code: RESPCODE_OK,
		Data: nil,
	}
	//延迟执行
	defer CoderspMsg(&resp, c)
	// 删除session
	err := ginsession.Destroy(c)
	if err != nil {
		fmt.Println("failed to logout destory session", err)
		resp.Code = RESPCODE_UNKNOWN
		return
	}
	return
}

// 授权
func Approve(c *gin.Context) {
	// 响应消息
	resp := RespMsg{
		Code: RESPCODE_OK,
		Data: nil,
	}
	//延迟执行
	defer CoderspMsg(&resp, c)
	//3. 从session获取用户和密码信息
	store := ginsession.FromContext(c)
	username, ok1 := store.Get("username")
	password, ok2 := store.Get("password")
	if !ok1 || !ok2 {
		resp.Code = RESPCODE_UNKNOWN
		fmt.Println("Failed to get session")
		return
	}
	// 合约调用
	txhash, err := bcos.Approve(username.(string), password.(string))
	if err != nil {
		resp.Code = RESPCODE_BLKERR
		fmt.Println("Failed to Approve")
		return
	}
	resp.Data = struct {
		Txhash string `json:"txhash"`
	}{
		txhash,
	}
	return
}
