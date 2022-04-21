<template>

<el-container>
  <el-header>
    <el-row>
      <el-col :span="24" align="center"><div class="grid-content bg-purple-dark">
            <schart class="wrapper" canvasId="canvas" :options="options"></schart>
      </div></el-col>
    </el-row>
  </el-header>

  <el-main>
    <el-col :span="5" offset="8">
      <el-input
        clearable
        v-model.number="input"
        maxlength="3"
        placeholder="请输入分段数"
        @keyup.enter="divData()"
      ></el-input>
    </el-col>
    <el-col :span="1">
      <el-button type="primary" plain @click="divData()">确定</el-button>
    </el-col>
    <el-col :span="2" offset="1">
      <el-button type="info" @click="getResult">查看方差和均值</el-button>
    </el-col>
  </el-main>
</el-container>
</template>

<script>
import Schart from "../chart/vue-schart";
const axios = require("axios");
export default {
    name: "Main",
    data() {
        return {
            options: {
                type: "bar",
                title: {
                    text: "信号分布图"
                },
                bgColor: "#f9eaf4",
                labels: [""],
                datasets: [  ],
                showValue: false,
                legend:{
                  display: false,
                } 
            },
            input:''
        };
    },
    created() {
        this.getData();
    },
    mounted(){
        document
        .querySelector("body")
        .setAttribute("style", "background-color:#f9eaf4");
    },
    components: {
        Schart
    },
    methods: {
        getColor(){
          return Math.round(Math.random()*250)
        },
        getData() {
            setTimeout(() => {
              this.$set(this.options, "datasets", data);
            }, 1000);
        },
        change(type) {
          this.options.type=type
        },
        getResult() {
          axios
          .get("api/result")
          .then((res) => {
            var result="<h4>均值:"+res.data.info.aver+"<\h4>";
            result+="<h4>方差:"+res.data.info.vari+"\<h4>";
            result+=res.data.info.create_time
            res.data.info.Average
            this.$alert(result, '统计方差和平均值', {
              confirmButtonText: '确定',
              dangerouslyUseHTMLString: true
            });
          })
          .catch((error) => {
            this.$alert('请求无响应', '网络错误', {
              confirmButtonText: '确定',
            });
            console.log(error)
          });
        },
        divData(){
          var str = "api/data" + "?cols=" + this.input
          var obj = this;
          axios
          .get(str)
          .then((res) => {
            console.log(res.data);
            obj.options.datasets=[];
            for(let i=0; i<res.data.info.length; i++){
              obj.options.datasets.push({
                label: res.data.info[i].lo+"<=x<"+res.data.info[i].hi,
                fillColor: "rgba("+obj.getColor()+", "+obj.getColor()+", 74, 0.5)",
                data: [res.data.info[i].value]
              })
            }
          })
          .catch((error) => {
            this.$alert('请求无响应', '网络错误', {
              confirmButtonText: '确定',
            });
            console.log(error)
          });
        }
    }
};
</script>
<style>


.wrapper {
    width: 500px;
    height: 400px;
}
.el-row {
    margin-bottom: 20px;
}

.el-rowlast-child {
    margin-bottom: 0;
}

.el-col {
    border-radius: 4px;
}
.bg-purple-dark {
    background: #f9eaf4;
}
.bg-purple {
    background: #f9eaf4;
}
.bg-purple-light {
    background: #e5e9f2;
}
.grid-content {
    border-radius: 4px;
    min-height: 36px;
}
.row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
}

  .el-header, .el-footer {
    background-color: #f9eaf4;
    color: #333;
    text-align: center;
    line-height: 60px;
  }

  .el-aside {
    background-color: #f9eaf4;
    color: #333;
    text-align: center;
    line-height: 200px;
  }
  
  .el-main {
    background-color: #f9eaf4;
    color: #333;
    text-align: center;
    line-height: 760px;
  }
  
  body > .el-container {
    margin-bottom: 40px;
  }
  
  .el-container:nth-child(5) .el-aside,
  .el-container:nth-child(6) .el-aside {
    line-height: 260px;
  }
  
  .el-container:nth-child(7) .el-aside {
    line-height: 320px;
  }
</style>