select em1.name as Employee from Employee as em1 
       left join Employee as em2 on 
       em1.ManagerId = em2.Id where em1.Salary > em2.Salary