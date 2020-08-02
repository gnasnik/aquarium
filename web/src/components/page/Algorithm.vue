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
                    @click="addAlgorithm"
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
                @expand-change="expandChange"
                @row-click="handleEdit"
            >
                <el-table-column type="expand">
                <template slot-scope="props">
                    <el-form label-position="left" inline class="demo-table-expand">
                    <el-table
                        :data="props.row.traders"
                        class="table"
                        ref="multipleTable"
                        style="width: 100%"
                        empty-text="No Data">
                        <el-table-column v-if="false" prop="id" label="ID"></el-table-column>
                        <el-table-column prop="name" label="Name"></el-table-column>
                         <!-- <el-table-column prop="status" label="Status" :formatter="statusFormat"> -->
                        <el-table-column label="Status">
                           <template slot-scope="scope">
                            <i  v-if="scope.row.status != 0" class="el-icon-loading"></i>
                            {{ statusFormat(scope.row.status) }}
                          </template>
                         </el-table-column>
                        <el-table-column prop="createdAt" label="CreatedAt" :formatter="dateFormat"></el-table-column>
                        <el-table-column prop="updatedAt" label="UpdatedAt" :formatter="dateFormat"></el-table-column>
                        <el-table-column label="Action">
                        <template slot-scope="scope">
                            <el-button class="handle-del mr10"  ref="button11" size="mini" @click="handleClickRun(scope.$index, scope.row)">{{runOrStop(scope.row)}}</el-button>
                            <el-button class="handle-del mr10"  size="mini" type="danger" @click="handleClickDelete(scope.$index, scope.row)">Delete</el-button>
                        </template>
                        </el-table-column>
                    </el-table>
                    </el-form>
                </template>
                </el-table-column>
                <el-table-column type="selection" width="55" align="center"></el-table-column>
                <el-table-column v-if="false" prop="id" label="ID" width="100" align="center"></el-table-column>
                <el-table-column prop="name" label="Name"></el-table-column>
                <el-table-column prop="description" label="Description"></el-table-column>
                <el-table-column prop="createdAt" label="CreatedAt" :formatter="dateFormat"></el-table-column>
                <el-table-column prop="updatedAt" label="UpdatedAt" :formatter="dateFormat"></el-table-column>
                <el-table-column prop="id" fixed="right" label="Action">
                    <template slot-scope="scope">
                        <el-button class="handle-deploy mr10" size="mini" @click="handleClickDeploy(scope.$index, scope.row)">Deploy</el-button>
                    </template>
                </el-table-column>
            </el-table> 
            <div class="pagination">
                <el-pagination
                    background
                    layout="total, prev, pager, next"
                    :current-page="query.pageIndex"
                    :page-size="query.pageSize"
                    :total="pageTotal"
                    @current-change="handlePageChange"
                ></el-pagination>
            </div>
        </div>

        <!-- 编辑弹出框 -->
        <el-dialog title="NewTrader" :visible.sync="editVisible" width="30%">
            <el-form ref="form" :model="form" label-width="100px" label-position="left">
                <!-- <el-form-item label="Algorithm:">
                    <el-input v-model="form.algorithmName" disabled></el-input>
                </el-form-item> -->
                <el-form-item label="Name:">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="Exchange:" prop="type">
                    <el-select v-model="form.exchanges" value-key="id" multiple placeholder="Select Exchange">
                    <el-option v-for="item in exchanges" 
                    :label="item.name" 
                    :value="item" 
                    :key="item.id">
                    </el-option>
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
import { exchangeListReq } from '../../api/exchange';
import { addTraderReq, traderListReq, delTraderReq, swithTrader } from '../../api/trader';
import { algorithmListReq, delAlgorithmReq } from '../../api/algorithm';
import { formatPlainDate , formatDateTime } from "../../utils/date";

export default {
    name: 'basetable',
    data() {
        return {
            query: {
                page: 1,
                size: 10
            },
            tableData: [],
            exchanges:[],
            traders:[],
            multipleSelection: [],
            editVisible: false,
            pageTotal: 0,
            form: {},
            idx: -1,
            id: -1,
        };
    },
    created() {
        this.token = localStorage.getItem("token");
        this.exchangeList();
        this.getData();
    },
    methods: {
        dateFormat(row, column, cellValue, index){
            if (!cellValue) {
                return ''
            }
            var date = new Date(cellValue);
            return formatDateTime(date);
        },
        statusFormat(row, column, cellValue, index){
            if (status == 0) {
                return "Halt"
            }
            return "Run"
        },
        runOrStop(row) {
            if (!row.status || row.status == 0){
                return "Run"
            }
            return "Stop"
        },
        getData() {
            algorithmListReq(this.query,this.token).then(res => {
                if (res.success) {
                    if (!res.data.algorithms) {
                        return 
                    }
                    for(var i = 0; i < res.data.algorithms.length; i++) {
                        res.data.algorithms[i].children = res.data.algorithms[i].traders;
                    }
                    this.tableData = res.data.algorithms;
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
        addAlgorithm() {
            this.$router.push('/editor');
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
            
            delAlgorithmReq(data,this.token).then(res => {
                if (res.success) {
                    this.multipleSelection = [];
                }else {
                    this.$message.error(res.msg || "unkown err");
                }
                
                this.getData();
            })
        },
        // 编辑操作
        handleClickDeploy(index,row) {
            var date = formatPlainDate(new Date());
            this.form.algorithmId = row.id;
            this.form.name = "New Trader @ " + date;
            this.form.algorithmName = row.name;
            this.editVisible = true;
        },
        handleClickDelete(index,row){
            delTraderReq({id:row.id},this.token).then(res => {
                if (res.success) {
                    traderListReq({algorithmId: row.algorithmId}, this.token).then(res => {
                        if (res.success) {
                            // 遍历当前页面表
                            this.tableData.forEach((temp, index) => {
                                // 找到当前点击的行，把动态获取到的数据赋值进去
                                if (temp.id === row.algorithmId) {
                                    this.tableData[index].traders = res.data.traders;
                                }
                            });
                        }else {
                            this.$message.error(res.msg || "unkown err");
                        }
                    })
                }else {
                    this.$message.error(res.msg || "unkown err");
                }
            });
        },
        handleClickRun(index,row) {
            swithTrader({id: row.id},this.token).then(res => {
                if (res.success) {    
                    row.status = !row.status
                }else { 
                    this.$message.error(res.msg || "unkown err");
                }
            })
        },
        // 保存编辑
        saveEdit() {
            if (!this.form.exchanges) {
                this.$message.error(`Please select a exchange`);
                return
            }
            console.log(this.form);
            addTraderReq(this.form, this.token).then(res => {
                if (res.success) {
                    // 动态加载展开页
                    traderListReq({algorithmId: this.form.algorithmId}, this.token).then(res => {
                        if (res.success) {
                            // 遍历当前页面表
                            this.tableData.forEach((temp, index) => {
                            // 找到当前点击的行，把动态获取到的数据赋值进去
                            if (temp.id === this.form.algorithmId) {
                                this.tableData[index].traders = res.data.traders;
                            }
                            });
                        }else {
                            this.$message.error(res.msg || "unkown err");
                        }
                    })
                }else {
                    this.$message.error(res.msg || "unkown err");
                }
            })
            this.editVisible = false;
        },
        handleEdit(row,column,event) {
            // this.$router.push('/editor',row);
            this.$router.push({path:'/editor', name:'editor', params:row})
        },
        // 分页导航
        handlePageChange(val) {
            this.$set(this.query, 'pageIndex', val);
            this.getData();
        },
        exchangeList(){
            exchangeListReq(this.query,this.token).then( res => {
                 if (res.success) {
                    this.exchanges = res.data.exchanges;
                }else {
                    this.$message.error(res.msg || "unkown err");
                }
            })
        },
        expandChange(row, expandedRows){
            // 该处是用于判断是展开还是收起行，只有展开的时候做请求，避免多次请求！
            // 展开的时候expandedRows有值，收起的时候为空.
             if (expandedRows.length > 0) {
                traderListReq({algorithmId: row.id}, this.token).then(res => {
                if (res.success) {
                    // 遍历当前页面表
                    this.tableData.forEach((temp, index) => {
                        // 找到当前点击的行，把动态获取到的数据赋值进去
                        if (temp.id === row.id) {
                            this.tableData[index].traders = res.data.traders;
                        }
                    });
                }else {
                    this.$message.error(res.msg || "unkown err");
                }
                })
             }
        },
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
.mr10 {
    margin-right: 10px;
}
.table-td-thumb {
    display: block;
    margin: auto;
    width: 40px;
    height: 40px;
}

  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }

  .unlight .light{
    width: 10px;
    height: 10px;
    background: red;
    color: red;
    /* border-radius:5px; */
    background-color:red;
  }
</style>
