<template>
    <div>
        <div class="crumbs">
        </div>
        <div class="container">
            <div class="handle-box">
                <el-button
                    type="primary"
                    class="handle-del mr10"
                    @click="reloadData"
                    size="mini"
                >Reload</el-button>
                <el-button
                    class="handle-del mr10"
                    @click="addJob"
                    size="mini"
                >Add</el-button>
                <el-button
                    class="handle-del mr10"
                    @click="handleDelete"
                    size="mini"
                >Delete</el-button>
            </div>
            <el-table
                :data="tableData"
                class="table"
                ref="multipleTable"
                empty-text="No Data"
                header-cell-class-name="table-header"
                @selection-change="handleSelectionChange"
                @row-click="handleEdit"
            >
                <el-table-column type="selection" width="55" align="center"></el-table-column>
                <el-table-column v-if="false" prop="id" label="ID" width="100" align="center"></el-table-column>
                <el-table-column prop="name" label="Name">
                     <template slot-scope="scope">
                       <div class="name">{{scope.row.name}}</div>
                       <div class="description">{{scope.row.description|| "no description"}}</div>
                    </template>
                </el-table-column>
                <el-table-column prop="running" label="Status">
                     <template slot-scope="scope">
                         <i class="el-icon-video-play" v-if="scope.row.status == 1"></i>
                         <i class="el-icon-video-pause" v-else></i>
                        <span>{{tranformStatus(scope.row.status)}}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="createdAt" label="CreatedAt" :formatter="dateFormat"></el-table-column>
                <el-table-column label="">
                    <template slot-scope="scope">
                        <div @click.stop>
                        <el-dropdown trigger="click" size="medium" @command="handleCommand">
                            <i class="el-icon-more"></i>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item icon="el-icon-switch-button" :command="composeValue('stop', scope.row)" v-if="scope.row.running">Job Stop</el-dropdown-item>
                                <el-dropdown-item icon="el-icon-video-play" :command="composeValue('run', scope.row)" v-else>Job Run</el-dropdown-item>
                                <el-dropdown-item icon="el-icon-refresh" :command="composeValue('restart', scope.row)">Job Restart</el-dropdown-item>
                                <el-dropdown-item icon="el-icon-document" :command="composeValue('viewlog', scope.row)">View Log</el-dropdown-item>
                            </el-dropdown-menu>
                            </el-dropdown>
                        </div>
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination
                    background
                    layout="total, prev, pager, next"
                    :current-page="query.page"
                    :page-size="query.size"
                    :total="pageTotal"
                    @current-change="handlePageChange"
                ></el-pagination>
            </div>
        </div>

        <!-- 编辑弹出框 -->
        <el-dialog :title="dialogTitle" :visible.sync="editVisible" width="30%">
            <el-form ref="form" :model="form" label-width="100px" label-position="left">
                <el-form-item label="Name:">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="Description:">
                    <el-input v-model="form.description"></el-input>
                </el-form-item>
                <el-form-item label="Exchange:" prop="exchange" >
                    <el-select v-model="form.exchangeId" placeholder="Select Exchanges" :disabled="edit">
                    <el-option v-for="item in exchanges" :label="item.name" :value="item.id" :key="item.id"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="Algorithm:" prop="algorithm" >
                    <el-select v-model="form.algorithmId" placeholder="Select Algorithm" :disabled="edit">
                    <el-option v-for="item in algorithms" :label="item.name" :value="item.id" :key="item.id"></el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">Cancel</el-button>
                <el-button type="primary" @click="saveEdit">OK</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
import { formatDateTime } from "../../utils/date";
import { jobListReq, addJobReq, delJobReq, switchJob } from '../../api/job';
import { exchangeListReq } from '../../api/exchange';
import { algorithmListReq } from '../../api/algorithm';
export default {
    name: 'basetable',
    data() {
        return {
            query: {
                page: 1,
                size: 10
            },
            token:"",
            // exchangeTypes:[],
            exchanges:[],
            algorithms:[],
            tableData: [],
            multipleSelection: [],
            editVisible: false,
            edit:false,
            pageTotal: 0,
            form: {},
            idx: -1,
            id: -1,
            dialogTitle:'',
        };
    },
    created() {
        this.token = localStorage.getItem("token");
        this.getData();
        this.getExchanges();
        this.getAlgorithms();
    },
    methods: {
        dateFormat(row, column, cellValue, index){
            if (!cellValue) {
                return ''
            }
            var date = new Date(cellValue);
            return formatDateTime(date);
        },
        getData() {
            jobListReq(this.query,this.token).then(res => {
                if (res.success) {
                    console.log(res.data.jobs)
                    this.tableData = res.data.jobs;
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
        reloadData() {
            this.getData();
        },
        addJob() {
            this.form = {};
            this.edit = false;
            this.dialogTitle = 'Add Job'
            this.editVisible = true;

            this.getExchanges();
            this.getAlgorithms();
        },
        getExchanges(){
            exchangeListReq(this.query,this.token).then(res => {
                if (res.success) {
                    this.exchanges = res.data.exchanges;
                }else{
                    if (res.code == 401) {   
                        localStorage.removeItem('ms_username');
                        this.$router.push("/login")
                    }
                    this.$message.error(res.msg || "unkown err");
                }  
            });
        },
        getAlgorithms(){
            algorithmListReq(this.query,this.token).then(res => {
                if (res.success) {
                    this.algorithms = res.data.algorithms;
                }else{
                    if (res.code == 401) {
                        localStorage.removeItem('ms_username');
                        this.$router.push("/login")
                    }
                    this.$message.error(res.msg || "unkown err");
                }  
            });
        },
        handleSwitch(row) {
            switchJob({id: row.id},this.token).then(res => {
                if (res.success) {    
                    row.status = row.status == 1? 0:1
                }else { 
                    this.$message.error(res.msg || "unkown err");
                }
            })
        },
        // 触发搜索按钮
        handleSearch() {
            this.$set(this.query, 'pageIndex', 1);
            this.getData();
        },
        // 删除操作
        handleDelete(index, row) {
            // 二次确认删除
            this.$confirm('Are you sure to DELETE ? ', '', {
                type: 'warning'
            }).then(() => {
                   this.delAllSelection();
            }).catch(() => {});
        },
        // 多选操作
        handleSelectionChange(val) {
            this.multipleSelection = val;
        },
        delAllSelection() {
            const length = this.multipleSelection.length;
            if (length <= 0) {
                return
            }

            let data = {
                ids:[],
            };

            for (let i=0; i< length; i++) {
               data.ids.push(this.multipleSelection[i].id)
            }

            delJobReq(data,this.token).then(res => {
                if (res.success) {
                    this.multipleSelection = [];
                }else {
                    this.$message.error(res.msg || "unkown err");
                }
            })

            this.getData();
        },
        // 编辑操作
        handleEdit(row,column,event) {
            this.form = row;
            this.edit = true;
            this.dialogTitle = 'Job - '+ row.name;
            this.editVisible = true;
        },
        // 保存编辑
        saveEdit() {
            if (!this.form.exchangeId || !this.form.algorithmId) {
                this.$message.error("please select exchange and algorithm!");
                return 
            } 
            addJobReq(this.form, this.token).then(res => {
                if (res.success) {
                    this.editVisible = false;
                    this.$message.success(`Success`);
                    // refesh
                    this.getData();
                }else{
                    this.$message.error(res.msg || "unkown err");
                }  
            })
        },
        // 分页导航
        handlePageChange(val) {
            this.$set(this.query, 'pageIndex', val);
            this.getData();
        },
        composeValue(name, row) {
            return {
                name: name,
                row: row,
            }
        },
        handleCommand(cmd) {
            if (cmd.name == "run" || cmd.name == "stop") {
                this.handleSwitch(cmd.row);
            }else if (cmd.name == "restart") {
                if (cmd.row.status == 1) {
                    this.handleSwitch(cmd.row);
                    setTimeout(() => {
                        this.handleSwitch(cmd.row);
                    }, 2000);
                }else{
                    this.handleSwitch(cmd.row);
                }
            }
        },
        tranformStatus(code) {
            if (code == 1) {
                return 'Running';
            }else if (code == 0) {
                return 'Stop';
            }else if (code == 2) {
                return 'Error';
            }else {
                return 'UnKnow';
            }
        }
    }
};
</script>

<style scoped>
.handle-box {
    margin: 10px 10px;
}

.handle-select {
    width: 120px;
}

.handle-input {
    width: 300px;
    display: inline-block;
}
.table {
    width: 100%;
    font-size: 12px;
}
.red {
    color: #ff0000;
}
.mr10 {
    margin-right: 10px;
}
.table-td-thumb {
    display: block;
    margin: auto;
    width: 40px;
    height: 40px;
}

span {
    margin-left: 5px;
}

.el-icon-video-pause {
    font-size: 16px;
    color: red;
}

.el-icon-video-play {
    font-size: 16px;
    color: rgb(10, 212, 70);
}

.name {
    color: #1E2736;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}

.description {
    font-size:8px;
    color: grey;
}

.el-icon-more {
    cursor: pointer;
    color: #8492a6;
    font-size: 20px;
}
</style>
