package models

type AccessLog struct {
	ID                uint64 `json:"id" mapstructure:"id"`                                     // 日志记录的唯一标识符
	RemoteAddr        string `json:"remote_addr" mapstructure:"remote_addr"`                   // 客户端IP地址
	RemoteUser        uint64 `json:"remote_user" mapstructure:"remote_user"`                   // 用户ID（从Token中解析得出）
	RequestMethod     string `json:"request_method" mapstructure:"request_method"`             // HTTP请求方法（如GET, POST等）
	RequestURI        string `json:"request_uri" mapstructure:"request_uri"`                   // 请求的URI，包括路径和查询参数
	ServerProtocol    string `json:"server_protocol" mapstructure:"server_protocol"`           // 请求使用的协议版本（如HTTP/1.1）
	Status            uint16 `json:"status" mapstructure:"status"`                             // HTTP响应状态码（如200, 404等）
	BodyBytesSent     uint64 `json:"body_bytes_sent" mapstructure:"body_bytes_sent"`           // 发送给客户端的字节数
	HttpReferer       string `json:"http_referer" mapstructure:"http_referer"`                 // HTTP Referer头信息，表示请求来源
	HttpUserAgent     string `json:"http_user_agent" mapstructure:"http_user_agent"`           // 客户端的User-Agent字符串
	HttpXForwardedFor string `json:"http_x_forwarded_for" mapstructure:"http_x_forwarded_for"` // X-Forwarded-For头信息，表示代理或负载均衡器传递的客户端IP
	RequestTime       uint64 `json:"request_time" mapstructure:"request_time"`                 // 请求处理的耗时（毫秒）
	RequestBody       string `json:"request_body" mapstructure:"request_body"`                 // 请求体内容（如POST数据）
	Timestamp         uint64 `json:"timestamp" mapstructure:"timestamp"`                       // 日志记录时间戳（UNIX时间戳，毫秒）
}
