var express = require('express');
var app = express();
var bodyParser = require('body-parser');
 

 
// 创建 application/x-www-form-urlencoded 编码解析
// var urlencodedParser = bodyParser.urlencoded({ extended: false })
//var urlencodedParser = bodyParser.json({ strict: false })
var urlencodedParser = bodyParser.raw({type:"application/json"})


app.get('/', function (req, res) {
   res.send('Hello World');
})


app.post('/a', urlencodedParser, function (req, res) {
   // 输出 JSON 格式
   response = {
       first_name:req.body,
       last_name:req.body
   };
  res.setHeader("Content-Type","application/json")
  res.send(req.body)
})
 


var server = app.listen(8081, function () {
 
  var host = server.address().address
  var port = server.address().port
 
  console.log("应用实例，访问地址为 http://%s:%s", host, port)
 
})