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
                <el-table-column prop="name" label="Name"></el-table-column>
                <el-table-column prop="description" label="Description"></el-table-column>
                <el-table-column prop="running" label="Running">
                     <template slot-scope="scope">
                        <el-button v-if="scope.row.running"
                            size="mini"
                            type="danger"
                            @click.stop="handleClickRun(scope.$index, scope.row)">Stop</el-button>
                        <el-button v-else
                            size="mini"
                            @click.stop="handleClickRun(scope.$index, scope.row)">Run</el-button>
                    </template>
                </el-table-column>
                <el-table-column prop="createdAt" label="CreatedAt" :formatter="dateFormat"></el-table-column>
                <el-table-column prop="updatedAt" label="UpdatedAt" :formatter="dateFormat"></el-table-column>

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
        // 编辑操作
        handleClickRun(index,row) {
            switchJob({id: row.id},this.token).then(res => {
                if (res.success) {    
                    row.running = !row.running
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
            this.dialogTitle = 'Job - '+ row.algorithm;
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
</style>
