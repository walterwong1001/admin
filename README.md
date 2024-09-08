说明

# MySQL

## Docker 安装

```bash
docker run --name mysql -p 3306:3306 --restart=always -v D:/docker/mysql/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=root2024 \
-d mysql:latest \
--character-set-server=utf8mb4 \
--collation-server=utf8mb4_unicode_ci
```
Petal ID System：花瓣 ID 系统

解释：每片花瓣的形态和纹理都是独特的，象征着算法生成的 ID 的美丽和独特性。

Aurora ID Generator：极光 ID 生成器

解释：极光的形态和颜色是独一无二的，每次出现都不同，象征着算法生成的 ID 的独特性和美丽。

Crystal ID Composer：水晶 ID 作曲者

解释：每块天然水晶的内部结构和形态都是独特的，象征着算法生成的 ID 的纯净和唯一性。

Ripple ID Engine：波纹 ID 引擎

解释：每个水波的形状和传播方式都是独特的，象征着算法生成的 ID 的动态和变化性。
Flame ID Generator：火焰 ID 生成器

解释：每束火焰的形态和燃烧方式都是独特的，象征着算法生成的 ID 的动态和能量。

Raindrop ID Generator：雨滴 ID 生成器

解释：每滴雨水的形态和落下轨迹都是独特的，象征着算法生成的 ID 的多样性和独特性。

Frost ID Composer：霜花 ID 作曲者

解释：每朵霜花的形态和结晶方式都是独特的，象征着算法生成的 ID 的精致和独特性。

Plume ID Composer：羽状 ID 作曲者

解释：每片羽毛的形态和花纹都是独特的，象征着算法生成的 ID 的轻盈和独特性。

Iris ID Generator：虹膜 ID 生成器

解释：每个人的虹膜图案都是独特的，象征着算法生成的 ID 的唯一性和个体特征。