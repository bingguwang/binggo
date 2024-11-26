

-- 方法一
-- 笛卡尔积

-- 先得出笛卡尔积
select a.* from Employee a, Employee b

-- 进行过滤
select a.* from Employee a, Employee b
where a.id != b.id and a.salary>=b.salary

-- 还要小于最高工资
select max(salary) from Employee

-- 结合一下
select a.* from Employee a, Employee b
where a.id != b.id and a.salary>=b.salary and a.salary < (select max(salary) from Employee)

-- 最后再去一下重
select distinct a.* from Employee a, Employee b
where a.id != b.id and a.salary>=b.salary and a.salary < (select max(salary) from Employee)

-- 解法二，排序

-- 先找出第二高的工资
SELECT DISTINCT salary FROM Employee ORDER BY salary DESC LIMIT 1 OFFSET 1

-- 然后工资是这个的人都返回就好了
SELECT *FROM Employee WHERE salary = (
    SELECT DISTINCT salary
    FROM Employee
    ORDER BY salary DESC
    LIMIT 1 OFFSET 1
)
