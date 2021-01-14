<template>
  <div class="container">
    <el-tabs v-model="activeName"  @tab-click="handleClick">
    <el-tab-pane label="All" name="all"></el-tab-pane>
    <el-tab-pane label="Info" name="info"></el-tab-pane>
    <el-tab-pane label="Error" name="error"></el-tab-pane>
  </el-tabs>
    <div v-for="item in logInfo"  v-bind:key="item.id">
        <span>{{item.createdAt}} [{{item.logType}}] {{item.content}}</span>
    </div>
  </div>
</template>

<script>
import { logListReq } from '../../api/log';
export default {
  data() {
      return {
        token:'',
        activeName: 'all',
        logInfo:[],
        query: {
            page: 1,
            size: 20,
            type: "",
        },
      };
  },
  created() {
        this.token = localStorage.getItem("token");
        if (this.$route.params) {
            console.log("params", this.$route.params.id);
            this.query.job_id = this.$route.params.id;
        }
        this.getLog();
    },
  methods: {
      getLog() {
        logListReq(this.query,this.token).then(res => {
            if (res.success) {
                this.logInfo = res.data.jobLogs;
                this.pageTotal = res.data.total || 50;
            }else{
                if (res.code == 401) {   
                    localStorage.removeItem('ms_username');
                    this.$router.push("/login")
                }
                this.$message.error(res.msg || "unkown err");
            }  
        });
      },
      handleClick(tab, event) {
        if (tab.name == "all") {
            this.query.type = ""
        }else if (tab.name == "info") {
            this.query.type = "INFO"
        }else if (tab.name == "error") {
            this.query.type = "Error"
        }
        this.getLog();
      },
  }
}
</script>


<style scoped>
.container {
    padding: 10px 20px;
    height: 100%;
    font-size: 14px;
    font-family: inherit;
    color: #1E2736;
}
</style>
