**Simple CSV processing service**

**Для запуска необходимо склонировать репозиторий, а затем запустить main.go указав путь к сsv файлу с данными:**

![image](https://user-images.githubusercontent.com/106326324/216029068-3a7441b6-3905-47f4-b5b1-e5500603fbfb.png)

**CSV файлы для тестирования находятся в testdata, базовый файл с данными в internal/data**

**Примеры тестов**

*Проверка на пустой файл*

![image](https://user-images.githubusercontent.com/106326324/216029208-0a9b9191-8cf8-46b8-bd5a-32c141dcdce6.png)

*Проверка на бесконечную рекурсию*

![image](https://user-images.githubusercontent.com/106326324/216029250-3d795133-d627-4a2f-af63-2a2cee13792f.png)

*Проверка на клетку с неверным типом данных*

![image](https://user-images.githubusercontent.com/106326324/216029293-f25dd72f-84bf-4d94-8567-fd76a0c4da70.png)

*Проверка на клетку с несуществующим индексом*

![image](https://user-images.githubusercontent.com/106326324/216029320-73023847-ca08-4960-948c-f0bf55388663.png)

*Проверка на неправильную индексацию*

![image](https://user-images.githubusercontent.com/106326324/216029346-af4844bf-69ac-4c84-b3fd-2769b2f54e36.png)

*Проверка на деление на 0*

![image](https://user-images.githubusercontent.com/106326324/216029374-f45c4c69-dc4a-4678-9130-308bc192ad41.png)
