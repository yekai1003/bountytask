// SPDX-License-Identifier: MIT
pragma solidity ^0.6.10;
pragma experimental ABIEncoderV2;

/*
    一个任务系统，带有奖励机制，任务状态变化，完成后，任务的执行这可以自动获取奖励
    角色： 发行方，执行方
    功能： 任务发行、接受、提交、确认
    结构设计： 
        任务信息
            任务id、发行方、接受方、奖励、描述、状态、时间戳、评价
*/

interface IERC20 {
    function balanceOf(address account) external view returns (uint256);
    function transfer(address to, uint256 amount) external returns (bool);
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) external returns (bool);
}

struct TaskInfo {
    address issuer;
    address worker;
    uint    bonus;
    string  desc;
    uint8   status;
    string  comment;
    uint    timestamp;
}

contract Task {
    // 定义任务数组
    TaskInfo[] tasks;
    uint256 constant FAUCETS = 100;
    mapping(address=>bool) faucets;
    address token;
    uint8 constant TASK_BEGIN   = 0;
    uint8 constant TASK_TAKE    = 1;
    uint8 constant TASK_COMMIT  = 2;
    uint8 constant TASK_CONFIRM = 3;

    constructor(address _token) public {
        token = _token;
    }

    function register() public {
        require(!faucets[msg.sender], "user already request faucets");
        faucets[msg.sender] = true;
        IERC20(token).transfer(msg.sender, FAUCETS);
    }

    // 任务发行
    function issue(uint _bonus, string memory _desc) public  {
        require(_bonus > 0, "task's bonus <= 0");
        // 判断用户余额充足
        require(IERC20(token).balanceOf(msg.sender) >= _bonus, "user's balance not enough");
        TaskInfo memory taskinfo = TaskInfo(msg.sender, address(0), _bonus, _desc, TASK_BEGIN, "", block.timestamp);
        tasks.push(taskinfo);
    }

    // 任务接受
    function take(uint _index) public  {
        require(_index < tasks.length, "index out of ranges");
        require(tasks[_index].status == TASK_BEGIN, "task's status error");
        require(tasks[_index].worker == address(0), "task's worker error");
        TaskInfo storage taskinfo = tasks[_index];
        taskinfo.status = TASK_TAKE;
        taskinfo.worker = msg.sender;
    }
    // 任务提交
    function commit(uint _index) public {
        require(_index < tasks.length, "index out of ranges");
        require(tasks[_index].status == TASK_TAKE, "task's status error");
        require(tasks[_index].worker == msg.sender, "task's worker error"); 
        TaskInfo storage taskinfo = tasks[_index];
        taskinfo.status = TASK_COMMIT;
    }
    // 任务确认
    function confirm(uint _index, uint8 _status, string memory _comment) public {
        require(_index < tasks.length, "index out of ranges");
        require(tasks[_index].status == TASK_COMMIT, "task's status error");
        require(tasks[_index].issuer == msg.sender, "task's issuer error"); 
        TaskInfo storage taskinfo = tasks[_index];
        taskinfo.comment = _comment;
        if(_status == TASK_CONFIRM) {
            // 任务通过
            taskinfo.status = TASK_CONFIRM;
            // 转账的操作
            // msg.sender -> (task.confirm)->(token.transfer)
            // 目标 issuer - > worker
            //IERC20(token).transfer(tasks[_index].worker, tasks[_index].bonus);
            // issuser 已经针对token合约对task合约进行了 approve
            IERC20(token).transferFrom(tasks[_index].issuer, tasks[_index].worker, tasks[_index].bonus);
        } else {
            // 任务没通过
            taskinfo.status = TASK_TAKE;
        }
    }

    function getOneTask(uint _index) public  view returns (TaskInfo memory) {
        return tasks[_index];
    }

    function getAllTasks() public  view  returns (TaskInfo[] memory) {
        return tasks;
    }
    
    function balanceOf(address account) public view returns (uint256) {
        return IERC20(token).balanceOf(account);
    }
}