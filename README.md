# Gogolook-Assignment

## Usage
Build
```
docker build -t james-task-api .
```

Run
```
docker run -p 8080:8080 james-task-api
```

## API Document
運行後可以在本機連線至 http://127.0.0.1:8080/swagger/index.html

## 說明
此作業整體需求不困難，單論需求能用簡易的架構即可完成，並且時間相當充裕。  
考量時間充足，假定此為大型專案，採用整潔架構來進行，在此架構下，雖有近半時間在建制架構，但臨時有想要更換資料庫或是業務邏輯的狀況，基本上僅需針對每層專責範圍進行擴充或是修改並調整依賴注入即可。
