# grpc_stream

GRPC 四种模式

● 简单模式
  ○ 跟普通的rpc没什么区别，一次请求，一次响应。
  
● 服务端数据流
  ○ 客户端一次请求，服务端返回一段连续的数据流。如：客户端发送股票代码，服务端把该数据的实时数据不断返回给客户端。
  
● 客户端流模式
  ○ 如：物联网终端想服务端发送数据-传感器
  
● 双向流模式
  ○ 如：聊天机器人
