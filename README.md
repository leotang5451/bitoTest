# bitoTest

## 編譯docker並執行
1. docker build -t bitotest .
2. docker run -p 8080:8080 --name bitotest bitotest

## Question 1 : Check out this golang program. What happens when this program runs?
Answer:透過goroutine呼叫transfer裡面的from、to進行lock互相等待，會發生死鎖。

## Question 2 : You are required to implement an API that queries a user's recent 100 purchased products. The API's RTT time should be lower than 50ms, so you need to use Redis as the data store. How would you store the data in Redis? How would you minimize memory usage?
Answer:使用Sorted Set記錄100個產品資料，使用user:product:user_id做為key, source存產品資料的timestamp，member存產品資料，新增資料做去從

## Question 3 : Please explain the difference between rolling upgrade and re-create Kubernetes deployment strategies, and the relationship between rolling upgrade and readiness probe.
Answer:
rolling upgrade:保持服務可用性，逐步將舊pod換成新pod，更新過程可以保持系統可用
re-create:將所有的pod一次刪除並重新建立，會導致服務在重新建立的過程中不可使用
Readiness Probe和rolling upgrade會結合使用，更新好的pod會先經過Readiness Probe確認就緒後才會開始接受流量

## Question 4 : Check out the following SQL. Of index A or B, which has better performance and why? SELECT * FROM orders WHERE user_id = ? AND created_at >= ? AND status = ? index A : idx_user_id_status_created_at(user_id, status, created_at) index B : idx_user_id_created_at_status(user_id, created_at, status) index C : idx_user_id_created_at(user_id, created_at)
Answer:index C的效能比較好，user_id是等值條件，created_at是範圍條件，通常status的範圍比較小，不在index也影響不大

## Question 5 : In the Kafka architecture design, how does kafka scale consumer-side performance? Does its solution have any drawbacks? Is there any counterpart to this drawback?
1. 增加Consumer Instances，需要調整kafka集群的負載，也會增加成本費用
solution:透過監控系統、自動化工具動態調整Consumer數量
2. 增加Partitions，會導致資料重分佈、再平衡
solution:使用自動化工具、腳本來管理Partitions的調整
3. 調整Consumer配置

## Question 6 : Please follow the following requirements to implement an HTTP server and post your GitHub repo link.
Design an HTTP server for the Tinder matching system. The HTTP server must support the following three APIs:
1. AddSinglePersonAndMatch : Add a new user to the matching system and find any possible matches for the new user.
2. RemoveSinglePerson : Remove a user from the matching system so that the user cannot be matched anymore.
3. QuerySinglePeople : Find the most N possible matched single people, where N is a request parameter.
Here is the matching rule:
- A single person has four input parameters: name, height, gender, and number of
wanted dates.
- Boys can only match girls who have lower height. Conversely, girls match boys who
are taller.
- Once the girl and boy match, they both use up one date. When their number of dates
becomes zero, they should be removed from the matching system.


Note : Please do not use other databases such as MySQL or Redis, just use in-memory data structure which in application to store your data.
Other requirements :
- Unit test
- Docker image
- Structured project layout
- API documentation
- System design documentation that also explains the time complexity of your API
- You can list TBD tasks.



