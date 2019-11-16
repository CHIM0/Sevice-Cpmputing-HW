博客API文档
HOST：http://api.myblog.cn/

获取认证[POST /auth]
- Request
    - Headers
        - Authorization: none
    - Body
    ```
    {
        "username": "用户名",
        "password": "密码"
    }
    ```
- Response 200 (application/json)
  ```
    {
        "ok": true,
        "data": {
            "token": "用户凭证",
            "time": "获取凭证时刻的时间戳"
        }
    }
  ```

注册用户[POST /users]
- Request
    - Headers
        - Content-Type：application/json
    - Body
    ```
    {
        "username": "用户名",
        "password": "密码",
        "email": "邮箱"
    }
    ```
- Response 200 (application/json)
  ```
    {
        "ok": true,
        "data": {
            "token": "用户凭证",
            "data":{
                username:"用户名",
                email:"邮箱地址"
            }
        }
    }
  ```

获取用户博客url列表[GET /article/{username}]
- Request
    - Headers
        - Authorization: token
- Response 200 (application/json)
  ```
    {
        "ok": true,
        "data": {
            articleUrllist:[
                {
                    title:"", // 文章标题
                    url:"" //文章url
                },
            ] //用户文章列表
        }
    }
  ```

用户在博客发表文章(md文件形式)[POST /{username}/articlefile]
- Request
    - Headers
        - Authorization: token
    - Body
    ```
    {
        title: "文章标题",
        file: "文件" //file格式(只接受.md文件)
    }
    ```
- Response 200 (application/json)
  ```
    {
        "ok": true,
        "data": {
            filename: "上传的文件名字"
        }
    }
  ```

用户取得文章的md文件 [GET /{username}/articlefile/{filename}]
- Request
    - Headers
        - Authorization: token
- Response 200 (application/json)
    - Headers
    ```
    {
        Content-Type：文件类型
    }
    ```
    - data
  ```
    {
        "ok": true,
        "data": {
            filename: "下载的文件名字"
        }
    }
  ```
  - 返回的文件