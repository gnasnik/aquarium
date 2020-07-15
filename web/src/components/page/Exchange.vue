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
                    @click="addExchange"
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
                @row-click="handleEdit"
            >
                <el-table-column type="selection" width="55" align="center"></el-table-column>
                <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
                <el-table-column prop="name" label="Name"></el-table-column>
                <el-table-column prop="type" label="Type"></el-table-column>
                <el-table-column prop="createdAt" label="CreatedAt"></el-table-column>
                <el-table-column prop="updatedAt" label="UpdatedAt"></el-table-column>

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
                <el-form-item label="Name">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="Type" prop="type" >
                    <el-select v-model="form.type" placeholder="Select Exchange Type" :disabled="edit">
                    <el-option v-for="item in exchangeTypes" :label="item" :value="item" :key="item"></el-option>
                    </el-select>
                </el-form-item>
                 <el-form-item label="AccessKey">
                    <el-input v-model="form.accessKey"></el-input>
                </el-form-item>
                 <el-form-item label="SecretKey">
                    <el-input v-model="form.secretKey"></el-input>
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
import { exchangeListReq, addExchangeReq, exchangeTypesReq, delExchangeReq } from '../../api/exchange';
export default {
    name: 'basetable',
    data() {
        return {
            query: {
                page: 1,
                size: 10
            },
            token:"",
            exchangeTypes:[],
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
        this.getExchangeTypes();
    },
    methods: {
        getData() {
            exchangeListReq(this.query,this.token).then(res => {
                if (res.success) {
                    this.tableData = res.data.exchanges;
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
        addExchange() {
            this.form = {};
            this.edit = false;
            this.dialogTitle = 'Add Exchange'
            this.editVisible = true;
        },
        getExchangeTypes() {
            exchangeTypesReq(this.query,this.token).then(res => {
                if (res.success) {
                    this.exchangeTypes = res.data.types;
                }else{
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

            delExchangeReq(data,this.token).then(res => {
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
            this.dialogTitle = 'Exchange - '+ row.name;
            this.editVisible = true;
        },
        // 保存编辑
        saveEdit() {
                addExchangeReq(this.form, this.token).then(res => {
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
</style>
