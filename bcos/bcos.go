package bcos

import (
	"boutytask/dbs"
	"boutytask/wallet"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
)

var cli *client.Client
var taskInstance *Task
var tokenInstance *Token

type RespTaskInfo struct {
	TaskID    string `json:"task_id"`
	Issuer    string `json:"issuer"`
	Worker    string `json:"task_user"`
	Comment   string `json:"comment"`
	Status    uint8  `json:"task_status"`
	Desc      string `json:"task_name"`
	Bounty    int64  `json:"bonus"`
	IssueDate string `json:"timestampp"`
}

var task_contract string = "0x869a70af9f51dbdbc5a66280b098d5f567f30c17"

const token_contract string = "0xbfd40e6887248e9e77914bdab3324d144db5e74d"

//init会在被包含时自动运行
func init() {
	//1. 解析配置文件
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatalf("ParseConfigFile failed, err: %v", err)
	}
	//2. 连接节点
	client, err := client.Dial(&configs[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to fisco ok")
	cli = client
	//3. 创建合约实例 NewHello + 合约地址 + 节点信息
	instance, err := NewTask(common.HexToAddress(task_contract), client)
	if err != nil {
		log.Panic("failed to NewTask", err)
	}
	taskInstance = instance

	// token 初始化
	token, err := NewToken(common.HexToAddress(token_contract), client)
	if err != nil {
		log.Panic("failed to NewToken", err)
	}
	tokenInstance = token
}

func Register(name, pass string) error {
	// 检测用户是否存在
	if len(dbs.Query([]byte(name))) > 0 {
		fmt.Printf("user[%s] already exists\n", name)
		return fmt.Errorf("user[%s] already exists", name)
	}
	// 创建私钥account
	w := wallet.NewWallet("./accounts")

	addr, keyJson, err := w.NewAccount(pass)
	fmt.Println("new acc:", addr)

	if err != nil {
		return err
	}
	dbs.AddData([]byte(name), []byte(addr))
	dbs.AddData([]byte(addr), keyJson)
	auth, err := w.GetTransactOpts(keyJson, pass)
	// 存入数据库
	_, _, err = taskInstance.Register(auth)
	_, _, err = tokenInstance.Approve(auth, common.HexToAddress(task_contract), big.NewInt(999999999999999))
	return err
}

func Login(name, pass string) (string, error) {
	keyJson := dbs.Query(dbs.Query([]byte(name)))
	w := wallet.NewWallet("./accounts")
	addr, err := w.ImportAccount(keyJson, pass)
	if err != nil {
		return "", fmt.Errorf("user not exists or password err")
	}
	return addr, nil
}

func Issue(name, pass string, amount int64, desc string) (string, error) {
	keyJson := dbs.Query(dbs.Query([]byte(name)))
	w := wallet.NewWallet("./accounts")
	fmt.Println(string(keyJson), pass)
	auth, err := w.GetTransactOpts(keyJson, pass)
	if err != nil {
		fmt.Println("failed to GetTransactOpts:", err)
		return "", err
	}
	tx, _, err := taskInstance.Issue(auth, big.NewInt(amount), desc)
	if err != nil {
		fmt.Println("failed to Issue:", err)
		return "", err
	}
	//fmt.Println(account.Address.String())
	//w.DeleteAccount(account, pass)
	return tx.Hash().Hex(), err
}

func Tasklist() ([]RespTaskInfo, error) {
	var tklist []RespTaskInfo

	tasks, err := taskInstance.GetAllTasks(cli.GetCallOpts())
	if err != nil {
		return tklist, err
	}
	task := RespTaskInfo{}
	for key, val := range tasks {
		task.TaskID = strconv.FormatInt(int64(key), 10)
		task.Bounty = val.Bonus.Int64()
		task.Issuer = val.Issuer.Hex()
		task.Comment = val.Comment
		task.Desc = val.Desc
		task.Worker = val.Worker.String()
		task.Status = val.Status
		task.IssueDate = time.UnixMilli(val.Timestamp.Int64()).Format("2006-01-02 15:04:05")
		//task.IssueDate = val.Timestamp.String()
		tklist = append(tklist, task)
	}
	return tklist, nil
}

//任务更新接口
//1 - 接受； 2- 提交； 3 - 确认； 4- 退回。
func Update(name, pass, comment string, taskID int64, status uint8) (string, error) {
	keyJson := dbs.Query(dbs.Query([]byte(name)))
	w := wallet.NewWallet("./accounts")
	fmt.Println(string(keyJson), pass)
	auth, err := w.GetTransactOpts(keyJson, pass)
	if err != nil {
		fmt.Println("failed to GetTransactOpts:", err)
		return "", err
	}

	if status == 1 {
		//接受任务
		tx, _, err := taskInstance.Take(auth, big.NewInt(taskID))
		return tx.Hash().Hex(), err
	} else if status == 2 {
		tx, _, err := taskInstance.Commit(auth, big.NewInt(taskID))
		return tx.Hash().Hex(), err
	} else if status == 3 || status == 4 {
		tx, _, err := taskInstance.Confirm(auth, big.NewInt(taskID), status, comment)
		return tx.Hash().Hex(), err
	}

	return "", nil
}

func BalanceOf(addr string) (*big.Int, error) {

	return taskInstance.BalanceOf(cli.GetCallOpts(), common.HexToAddress(addr))

}

func Approve(name, pass string) (string, error) {
	keyJson := dbs.Query(dbs.Query([]byte(name)))
	w := wallet.NewWallet("./accounts")
	fmt.Println(string(keyJson), pass)
	auth, err := w.GetTransactOpts(keyJson, pass)
	if err != nil {
		fmt.Println("failed to GetTransactOpts:", err)
		return "", err
	}
	tx, _, err := tokenInstance.Approve(auth, common.HexToAddress(task_contract), big.NewInt(99999999999999999))
	if err != nil {
		fmt.Println("failed to approve", err)
		return "", err
	}
	return tx.Hash().String(), nil

}
