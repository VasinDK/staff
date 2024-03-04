## Тестовое задание: добавление, удаление, получение и обновление сотрудников
Документация в swagger и тесты будут позже. 

А сейчас для проверки можно воспользоваться следующими методами:

### Получает сотрудников по id компании и id департамента
GET localhost:8080/v1/staff/1?department_id=1

### Добавляет сотрудника
POST localhost:8080/v1/staff

    {
        "name": "Кирил",
        "surname": "Ручников",
        "phone": "877799977",
        "companyId": 1,
        "passportNumber": 1122331,
        "passportTypeId": 2,
        "departmentId": 1
    }

### Обновляет поля сотрудника по id
PUT localhost:8080/v1/staff

    {
        "id": 1,
        "name": "Максим",
        "surname": "Ручников",
        "phone": "012255",
        "companyId": 1,
        "passportNumber": 1122303,
        "passportTypeId": 20,
        "departmentId": 2
    }

### Удаляет сотрудника по id
DELETE localhost:8080/v1/staff/5