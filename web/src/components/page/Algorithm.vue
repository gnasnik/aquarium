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
                >Reload</el-button>
                  <el-button
                    class="handle-del mr10"
                    @click="addAlgorithm"
                >Add</el-button>
                <el-button
                    class="handle-del mr10"
                    @click="handleDelete"
                >Delete</el-button>
            </div>
            <el-table
                :data="tableData"
                class="table"
                ref="multipleTable"
                empty-text="No Data"
                header-cell-class-name="table-header"
                @selection-change="handleSelectionChange"
                row-key="name"
                :row-class-name="tableRowClassName"
                :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
            >
                <el-table-column type="selection" width="55" align="center"></el-table-column>
                <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
                <el-table-column prop="name" label="Name"></el-table-column>
                <el-table-column prop="description" label="Description"></el-table-column>
                <el-table-column prop="createdAt" label="CreatedAt"></el-table-column>
                <el-table-column prop="updatedAt" label="UpdatedAt"></el-table-column>
                <el-table-column prop="id" fixed="right" label="Action">
                    <template slot-scope="scope">
                        <el-button class="handle-deploy mr10" @click="handleClick(scope.row)" v-if="isTree(scope.row)">Run</el-button>
                        <el-button class="handle-deploy mr10" @click="handleClick(scope.row)" v-else >Deploy</el-button>
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
                <el-form-item label="Algorithm:">
                    <el-input v-model="form.algorithmName" disabled></el-input>
                </el-form-item>
                <el-form-item label="Name:">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="Exchange:" prop="type">
                    <el-select v-model="form.exchange" placeholder="Select Exchange">
                    <el-option v-for="item in exchanges" :label="item.name" :value="item" :key="item.id"></el-option>
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
import { addTraderReq, traderListReq } from '../../api/trader';
import { algorithmListReq, delAlgorithmReq } from '../../api/algorithm';
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
            id: -1
        };
    },
    created() {
        this.token = localStorage.getItem("token");
        this.getData();
        this.exchangeList();
    },
    methods: {
        isTree(row){
            if (row.traders) {
                return false
            }
            return true
        },
        getData() {
            algorithmListReq(this.query,this.token).then(res => {
                if (res.success) {
                    // let length = res.data.algorithms.length;
                    for(var i = 0; i < res.data.algorithms.length; i++) {
                        let traders = res.data.algorithms[i].traders;
                        if (traders) {
                            for (var j = 0; j <  traders.length; j++) {
                                res.data.algorithms[i].traders[j].id = ""; 
                            }   
                        }
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
        tableRowClassName({row, rowIndex}) {
            if (row.traders) {
                return '';
            }
            return 'success-row';;
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
        handleClick(row) {
            // algorithm
            if (row.traders) {
                this.form.algorithmId = row.id;
                this.form.algorithmName = row.name;
                this.editVisible = true;
            }
            // tree-trader
            this.$message.error(`not implement yet`);
        },
        // 保存编辑
        saveEdit() {
            if (!this.form.exchange) {
                this.$message.error(`Please select a exchange`);
                return
            }
            this.addTrader();
            this.editVisible = false;
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
        addTrader(){
            this.form.name = "trader@" + this.form.name;
            addTraderReq(this.form, this.token).then(res => {
                if (res.success) {
                    this.$message.success(`Success`);
                }else {
                    this.$message.error(res.msg || "unkown err");
                }
            })
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
    font-size: 14px;
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

.el-table .success-row {
background: #f0f9eb;
}
</style>
