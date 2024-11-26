
--笛卡尔积来解决
-- 先生成笛卡尔积
SELECT a.*, b.* from Employee a, Employee b;

-- 然后看下前面的salary大于后面的有哪些
SELECT a.*, b.* from Employee a, Employee b where a.salary > b.salary ;

-- 最后的结果就是这样
SELECT a.*, b.* from Employee a, Employee b where a.salary > b.salary and a.managerId = b.id;


-- 其实还有种子查询的解法，但是效率低
SELECT
    a.name Employee
FROM
    test.Employee a -- a是员工
WHERE
salary > (SELECT -- 子查询结果是领导的工资
       salary
FROM
    Employee b
WHERE
    a.managerId = b.id);