# golang_session_module
基于redis的session验证模块


### 测试接口:
# 127.0.0.1:8080/v1/register
# 127.0.0.1:8080/v1/login
# 127.0.0.1:8080/v1/auth

### auth为基于登录状态下的会话请求 登录后redis与本地cookie池中的session_id均保持60秒

### 使用前注意redis及mysql数据库配置
