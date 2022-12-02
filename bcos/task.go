// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bcos

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// TaskInfo is an auto generated low-level Go binding around an user-defined struct.
type TaskInfo struct {
	Issuer    common.Address
	Worker    common.Address
	Bonus     *big.Int
	Desc      string
	Status    uint8
	Comment   string
	Timestamp *big.Int
}

// TaskABI is the input ABI used to generate the binding from.
const TaskABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_comment\",\"type\":\"string\"}],\"name\":\"confirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTasks\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structTaskInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOneTask\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structTaskInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bonus\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_desc\",\"type\":\"string\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"take\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Task is an auto generated Go binding around a Solidity contract.
type Task struct {
	TaskCaller     // Read-only binding to the contract
	TaskTransactor // Write-only binding to the contract
	TaskFilterer   // Log filterer for contract events
}

// TaskCaller is an auto generated read-only Go binding around a Solidity contract.
type TaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskTransactor is an auto generated write-only Go binding around a Solidity contract.
type TaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TaskSession struct {
	Contract     *Task             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TaskCallerSession struct {
	Contract *TaskCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TaskTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TaskTransactorSession struct {
	Contract     *TaskTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskRaw is an auto generated low-level Go binding around a Solidity contract.
type TaskRaw struct {
	Contract *Task // Generic contract binding to access the raw methods on
}

// TaskCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TaskCallerRaw struct {
	Contract *TaskCaller // Generic read-only contract binding to access the raw methods on
}

// TaskTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TaskTransactorRaw struct {
	Contract *TaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTask creates a new instance of Task, bound to a specific deployed contract.
func NewTask(address common.Address, backend bind.ContractBackend) (*Task, error) {
	contract, err := bindTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// NewTaskCaller creates a new read-only instance of Task, bound to a specific deployed contract.
func NewTaskCaller(address common.Address, caller bind.ContractCaller) (*TaskCaller, error) {
	contract, err := bindTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskCaller{contract: contract}, nil
}

// NewTaskTransactor creates a new write-only instance of Task, bound to a specific deployed contract.
func NewTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskTransactor, error) {
	contract, err := bindTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskTransactor{contract: contract}, nil
}

// NewTaskFilterer creates a new log filterer instance of Task, bound to a specific deployed contract.
func NewTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskFilterer, error) {
	contract, err := bindTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskFilterer{contract: contract}, nil
}

// bindTask binds a generic wrapper to an already deployed contract.
func bindTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Task.Contract.TaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.TaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.TaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Task.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Task *TaskCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Task *TaskSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Task *TaskCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, account)
}

// GetAllTasks is a free data retrieval call binding the contract method 0x779900b4.
//
// Solidity: function getAllTasks() constant returns([]TaskInfo)
func (_Task *TaskCaller) GetAllTasks(opts *bind.CallOpts) ([]TaskInfo, error) {
	var (
		ret0 = new([]TaskInfo)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "getAllTasks")
	return *ret0, err
}

// GetAllTasks is a free data retrieval call binding the contract method 0x779900b4.
//
// Solidity: function getAllTasks() constant returns([]TaskInfo)
func (_Task *TaskSession) GetAllTasks() ([]TaskInfo, error) {
	return _Task.Contract.GetAllTasks(&_Task.CallOpts)
}

// GetAllTasks is a free data retrieval call binding the contract method 0x779900b4.
//
// Solidity: function getAllTasks() constant returns([]TaskInfo)
func (_Task *TaskCallerSession) GetAllTasks() ([]TaskInfo, error) {
	return _Task.Contract.GetAllTasks(&_Task.CallOpts)
}

// GetOneTask is a free data retrieval call binding the contract method 0x9c657a38.
//
// Solidity: function getOneTask(uint256 _index) constant returns(TaskInfo)
func (_Task *TaskCaller) GetOneTask(opts *bind.CallOpts, _index *big.Int) (TaskInfo, error) {
	var (
		ret0 = new(TaskInfo)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "getOneTask", _index)
	return *ret0, err
}

// GetOneTask is a free data retrieval call binding the contract method 0x9c657a38.
//
// Solidity: function getOneTask(uint256 _index) constant returns(TaskInfo)
func (_Task *TaskSession) GetOneTask(_index *big.Int) (TaskInfo, error) {
	return _Task.Contract.GetOneTask(&_Task.CallOpts, _index)
}

// GetOneTask is a free data retrieval call binding the contract method 0x9c657a38.
//
// Solidity: function getOneTask(uint256 _index) constant returns(TaskInfo)
func (_Task *TaskCallerSession) GetOneTask(_index *big.Int) (TaskInfo, error) {
	return _Task.Contract.GetOneTask(&_Task.CallOpts, _index)
}

// Commit is a paid mutator transaction binding the contract method 0xf4f98ad5.
//
// Solidity: function commit(uint256 _index) returns()
func (_Task *TaskTransactor) Commit(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "commit", _index)
}

func (_Task *TaskTransactor) AsyncCommit(handler func(*types.Receipt, error), opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "commit", _index)
}

// Commit is a paid mutator transaction binding the contract method 0xf4f98ad5.
//
// Solidity: function commit(uint256 _index) returns()
func (_Task *TaskSession) Commit(_index *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, _index)
}

func (_Task *TaskSession) AsyncCommit(handler func(*types.Receipt, error), _index *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncCommit(handler, &_Task.TransactOpts, _index)
}

// Commit is a paid mutator transaction binding the contract method 0xf4f98ad5.
//
// Solidity: function commit(uint256 _index) returns()
func (_Task *TaskTransactorSession) Commit(_index *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, _index)
}

func (_Task *TaskTransactorSession) AsyncCommit(handler func(*types.Receipt, error), _index *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncCommit(handler, &_Task.TransactOpts, _index)
}

// Confirm is a paid mutator transaction binding the contract method 0xd17fae07.
//
// Solidity: function confirm(uint256 _index, uint8 _status, string _comment) returns()
func (_Task *TaskTransactor) Confirm(opts *bind.TransactOpts, _index *big.Int, _status uint8, _comment string) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "confirm", _index, _status, _comment)
}

func (_Task *TaskTransactor) AsyncConfirm(handler func(*types.Receipt, error), opts *bind.TransactOpts, _index *big.Int, _status uint8, _comment string) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "confirm", _index, _status, _comment)
}

// Confirm is a paid mutator transaction binding the contract method 0xd17fae07.
//
// Solidity: function confirm(uint256 _index, uint8 _status, string _comment) returns()
func (_Task *TaskSession) Confirm(_index *big.Int, _status uint8, _comment string) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, _index, _status, _comment)
}

func (_Task *TaskSession) AsyncConfirm(handler func(*types.Receipt, error), _index *big.Int, _status uint8, _comment string) (*types.Transaction, error) {
	return _Task.Contract.AsyncConfirm(handler, &_Task.TransactOpts, _index, _status, _comment)
}

// Confirm is a paid mutator transaction binding the contract method 0xd17fae07.
//
// Solidity: function confirm(uint256 _index, uint8 _status, string _comment) returns()
func (_Task *TaskTransactorSession) Confirm(_index *big.Int, _status uint8, _comment string) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, _index, _status, _comment)
}

func (_Task *TaskTransactorSession) AsyncConfirm(handler func(*types.Receipt, error), _index *big.Int, _status uint8, _comment string) (*types.Transaction, error) {
	return _Task.Contract.AsyncConfirm(handler, &_Task.TransactOpts, _index, _status, _comment)
}

// Issue is a paid mutator transaction binding the contract method 0x9169d937.
//
// Solidity: function issue(uint256 _bonus, string _desc) returns()
func (_Task *TaskTransactor) Issue(opts *bind.TransactOpts, _bonus *big.Int, _desc string) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "issue", _bonus, _desc)
}

func (_Task *TaskTransactor) AsyncIssue(handler func(*types.Receipt, error), opts *bind.TransactOpts, _bonus *big.Int, _desc string) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "issue", _bonus, _desc)
}

// Issue is a paid mutator transaction binding the contract method 0x9169d937.
//
// Solidity: function issue(uint256 _bonus, string _desc) returns()
func (_Task *TaskSession) Issue(_bonus *big.Int, _desc string) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, _bonus, _desc)
}

func (_Task *TaskSession) AsyncIssue(handler func(*types.Receipt, error), _bonus *big.Int, _desc string) (*types.Transaction, error) {
	return _Task.Contract.AsyncIssue(handler, &_Task.TransactOpts, _bonus, _desc)
}

// Issue is a paid mutator transaction binding the contract method 0x9169d937.
//
// Solidity: function issue(uint256 _bonus, string _desc) returns()
func (_Task *TaskTransactorSession) Issue(_bonus *big.Int, _desc string) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, _bonus, _desc)
}

func (_Task *TaskTransactorSession) AsyncIssue(handler func(*types.Receipt, error), _bonus *big.Int, _desc string) (*types.Transaction, error) {
	return _Task.Contract.AsyncIssue(handler, &_Task.TransactOpts, _bonus, _desc)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Task *TaskTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "register")
}

func (_Task *TaskTransactor) AsyncRegister(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Task *TaskSession) Register() (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Register(&_Task.TransactOpts)
}

func (_Task *TaskSession) AsyncRegister(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Task.Contract.AsyncRegister(handler, &_Task.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Task *TaskTransactorSession) Register() (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Register(&_Task.TransactOpts)
}

func (_Task *TaskTransactorSession) AsyncRegister(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Task.Contract.AsyncRegister(handler, &_Task.TransactOpts)
}

// Take is a paid mutator transaction binding the contract method 0x4fd9efc4.
//
// Solidity: function take(uint256 _index) returns()
func (_Task *TaskTransactor) Take(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "take", _index)
}

func (_Task *TaskTransactor) AsyncTake(handler func(*types.Receipt, error), opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "take", _index)
}

// Take is a paid mutator transaction binding the contract method 0x4fd9efc4.
//
// Solidity: function take(uint256 _index) returns()
func (_Task *TaskSession) Take(_index *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, _index)
}

func (_Task *TaskSession) AsyncTake(handler func(*types.Receipt, error), _index *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncTake(handler, &_Task.TransactOpts, _index)
}

// Take is a paid mutator transaction binding the contract method 0x4fd9efc4.
//
// Solidity: function take(uint256 _index) returns()
func (_Task *TaskTransactorSession) Take(_index *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, _index)
}

func (_Task *TaskTransactorSession) AsyncTake(handler func(*types.Receipt, error), _index *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncTake(handler, &_Task.TransactOpts, _index)
}
