const Mock = require('mockjs')

const Random = Mock.Random

let data = [] // 用于接受生成数据的数组

//类型规则
let genresRules = ["踏实", "实事求", "条理分明", "理解力强", "善于分析", "适应性强", "大胆的", "脾气暴燥", "专横跋扈", "有同情心", "自以为是", "有信心的", "有教养", "奉献精神", "健忘", "有礼貌", "贪婪", "诚实", "公正", "仁慈", " 懒惰", "幽默", "殷勤", "勤劳", "俭朴", "友好", "善于表达", "充满热情", "精力充沛"];
//职业规则
let professionalRules = ["医生", "老师", "飞行员", "邮递员", "警察", "护士", "科学家", "美术家", "回家", "歌手", "办公室职员", "经理", "老板", "助手", "空姐", "海军", "工程师", "作家", "航海家", "空军", "陆军", "舞蹈家", "书法家", "模特", "护士", "导演", "演员", "服务员", "法官", "秘书", "话务员", "健身教练", "美容师", "程序员"];





//获取--count参数的索引
let index = process.argv.lastIndexOf("--count");
//下一次则为要生成条数
let length = process.argv[index + 1]
console.log("共生成：\t", length, "条数据，请核实确认。")

for (i = 0; i < length; i++) {

  const profilePhotoRules = '#' + Random.integer(180, 255).toString(16) +
    Random.integer(140, 255).toString(16) +
    Random.integer(120, 220).toString(16);


  // 人员数据生成规则
  let personRules = {
    "id": "@id",
    "guid": "@guid",
    "user_name": "@cname()",
    //手机号码正则
    "mobile": /^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$/,
    "address": "@county(true)",
    "brthday": "@datetime('yyyy-MM-dd HH:mm:ss')",
    "email": "@email",
    "weight": "@float(80, 200, 1, 1)",
    "height": "@float(150, 188, 1, 1)",
    "gender": "@character(男女)",
    "introduce": "@cparagraph()",
    "motto": "@csentence(5)",
    "genres": Random.pick(genresRules, 2, 4),
    "professional": Random.pick(professionalRules, 1),
    "poster": function () {
      let name = this.user_name;
      return Random.image('240x260', profilePhotoRules, name);
    },
    "title": function () {
      //标题为一些混合
      let title = this.user_name + "," + this.professional + "," + this.genres;
      return title
    }
  }

  data.push(personRules)
}

//生成
persons = Mock.mock(data)
//转json
let jsonData = JSON.stringify(persons)
//写出文件
const fs = require('fs')
const path = require('path')

let dir = path.join(__dirname, 'mock_data.json')//这里写你要存入哪个文件的路径
fs.writeFile(dir, jsonData, 'utf8', (err) => {
  if (err != null) {
    console.log("写入失败!");
    return
  }
  console.log("写入成功!");
})
//打印mock
// console.log(persons)