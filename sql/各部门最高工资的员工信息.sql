https://leetcode-cn.com/problems/department-highest-salary/

    主要考察子查询

-- 解法一，子查询

-- 先找出每个部门的最高
SELECT department_id ,max(salary)  FROM test.employee group by department_id;

-- 利用上面的子查询查出各部门里最高工资的员工信息
select * from  test.employee where (department_id, salary) in (
    SELECT department_id ,max(salary)  FROM test.employee group by department_id
);

-- 补全员工信息
select a.*, b.name as dpName from  test.employee a
left join test.department b
on b.id = a.department_id
where (department_id, salary) in (
    SELECT department_id ,max(salary)  FROM test.employee group by department_id
)

-- 解法二

-- 仔细看，上面的子查询是可以优化的

select a.*, b.name as dpName from  test.employee a
left join test.department b
on b.id = a.department_id
where salary in (
    SELECT max(salary)  FROM test.employee where department_id = b.id
);

