说明

# MySQL

## Docker 安装

```bash
docker run --name mysql -p 3306:3306 -v D:/docker/mysql/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=root2024 \
-d mysql:latest \
--character-set-server=utf8mb4 \
--collation-server=utf8mb4_unicode_ci
```
