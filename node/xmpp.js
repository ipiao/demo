
var ws = require("nodejs-websocket");
console.log("开始建立连接...")

var socket = new WebSocket('wss://xmpp.ickapay.com/websocket');


// 建立 websocket 连接成功触发事件
ws.onopen = function () {
  console.log("连接成功");
  
   var  from={"serialNo":"xxx","idNo":"xxx"}
   ws.send(JSON.stringify(from));//发送参数
};

// 当服务端处理完成后会将数据发送回来
ws.onmessage = function (evt) {
  var result= evt.data;
  var data = JSON.parse(event.data)
  console.log(data)    
};
