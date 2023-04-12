## 快速部署在服务器上

### 部署前的tips:

- 安装docker

- 需要用到nginx，mysql

- 确保端口3306，80，3000未被占用（lsof -i:端口号查询）

- 打开对应端口防火墙，开启443的https连接

  ### 开始部署

  1.  进入主目录（含dockerfile），执行以下命令获取镜像

     ```go
      docker build -t project .
     ```

     2. 安装mysql，这里直接一个命令获取

        ```go
        docker run --name mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -v /docker/mysql:/var/lib/mysql mysql:5.7
        ```

        3. 安装nginx，这里直接一个命令获取

           ```go
           docker run -d --restart=always \ 
           -p 80:80 -p 443:443 \ 
           --name nginx \
           -v /docker/nginx/nginx.conf:/etc/nginx/nginx.conf \
           -v /docker/nginx/conf.d:/etc/nginx/conf.d \
           -v /docker/nginx/logs:/var/log/nginx \
           -v /docker/nginx/cert:/etc/nginx/cert \
           nginx:1.19.4 
           
           //tips:
           //在挂载数据卷之前得创建，如上示例，并在给定卷中填写配置文件，如反向代理、https等
           ```

           4. 回到得到的项目镜像，把它启动

              ```go
              docker run -p 3000:3000 --link mysql:mysql project
              
              // 利用--link，项目中指定 ip:port为 mysql:3306 即可连接到mysql
              ```

              5. 如果nginx配置好了域名，便可以访问了。

  ###  

  

  