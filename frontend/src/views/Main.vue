<template>
<el-row>
  <el-col :span="24" align="center"><div class="grid-content bg-purple-dark">
      <schart class="wrapper" canvasId="canvas" :options="options"></schart>
  </div></el-col>
</el-row>
    <!-- <div>
        <schart class="wrapper" canvasId="canvas" :options="options"></schart>
        <button @click="change('bar')">change to bar</button>
        <button @click="change('line')">change to line</button>
        <button @click="change('pie')">change to pie</button>
        <button @click="change('ring')">change to ring</button>
    </div> -->

</template>

<script>
import Schart from "../chart/vue-schart";
export default {
    name: "Main",
    data() {
        return {
            example: {
                label: "x<=100",
                fillColor: "rgba(241, 49, 74, 0.5)",
                data: [234]
            },
            options: {
                type: "bar",
                title: {
                    text: "信号分布图"
                },
                bgColor: "#f9eaf4",
                labels: ["x 为信号量"],
                datasets: [  ]
            }
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
        getData() {
            setTimeout(() => {
                this.options.datasets.push(this.example)
                this.options.datasets.push(this.example)
                this.$set(this.options, "datasets", data);
            }, 1000);
        },
        change(type) {
          this.options.type=type
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
    background: #d3dce6;
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
</style>