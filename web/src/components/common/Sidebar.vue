<template>
    <div class="sidebar">
       
        <el-menu
            class="sidebar-el-menu"
            :default-active="onRoutes"
            :collapse="collapse"
            background-color="#324157"
            text-color="#bfcbd9"
            active-text-color="#20a0ff"
            unique-opened
            router
        >
        <div class="header" v-if="!collapse">
        <div class="logo">Aquarium</div>
        </div>
            <template v-for="item in items">
                <template v-if="item.subs">
                    <el-submenu :index="item.index" :key="item.index">
                        <template slot="title">
                            <i :class="item.icon"></i>
                            <span slot="title">{{ item.title }}</span>
                        </template>
                        <template v-for="subItem in item.subs">
                            <el-submenu
                                v-if="subItem.subs"
                                :index="subItem.index"
                                :key="subItem.index"
                            >
                                <template slot="title">{{ subItem.title }}</template>
                                <el-menu-item
                                    v-for="(threeItem,i) in subItem.subs"
                                    :key="i"
                                    :index="threeItem.index"
                                >{{ threeItem.title }}</el-menu-item>
                            </el-submenu>
                            <el-menu-item
                                v-else
                                :index="subItem.index"
                                :key="subItem.index"
                            >{{ subItem.title }}</el-menu-item>
                        </template>
                    </el-submenu>
                </template>
                <template v-else>
                    <el-menu-item :index="item.index" :key="item.index" @click="itemClick(item.title)">
                        <i :class="item.icon"></i>
                        <span slot="title">{{ item.title }}</span>
                    </el-menu-item>
                </template>
            </template>
            <!-- 折叠按钮 -->
        <div class="bottom-collapse">
            <div class="collapse-btn" @click="collapseChage">
                <i v-if="!collapse" class="el-icon-fk-left"></i>
                <i v-else class="el-icon-fk-right"></i>
            </div>
        </div>
        </el-menu>
    </div>
</template>

<script>
import bus from '../common/bus';
export default {
    data() {
        return {
            collapse: false,
            items: [
                {
                    icon: 'el-icon-fk-compute',
                    index: 'algorithm',
                    title: 'Algorithm'
                },
                {
                    icon: 'el-icon-fk-bank',
                    index: 'exchange',
                    title: 'Exchange'
                },
                {
                    icon: 'el-icon-fk-user',
                    index: 'userinfo',
                    title: 'User'
                },
                {
                    icon: 'el-icon-fk-logout',
                    // index: 'logout',
                    title: 'Logout'
                }
            ]
        };
    },
    computed: {
        onRoutes() {
            return this.$route.path.replace('/', '');
        }
    },
    methods:{
        // 侧边栏折叠
        collapseChage() {
            this.collapse = !this.collapse;
            bus.$emit('collapse', this.collapse);
        },
        itemClick(v){
            if (v == "Logout") {
                this.open()
            }
        },
        open() {
        this.$confirm('Are you sure to logout ?',"", {
          confirmButtonText: 'OK',
          cancelButtonText: 'Cancel',
          type: 'warning',
        }).then(() => {
            localStorage.removeItem('ms_username');
            this.$router.push("/login")
        });
      }
    },
    created() {
        // 通过 Event Bus 进行组件间通信，来折叠侧边栏
        bus.$on('collapse', msg => {
            this.collapse = msg;
            bus.$emit('collapse-content', msg);
        });
    },
    mounted() {
        if (document.body.clientWidth < 1200) {
            this.collapseChage();
        }
    }
};
</script>

<style scoped>
.sidebar {
    display: block;
    position: absolute;
    left: 0;
    /* top: 50px; */
    top: 0;
    bottom: 0; 
    overflow-y: scroll;
    /* padding-bottom: 42px; */
}
.bottom-collapse {
    position: absolute;
    bottom:0;
    width: 100%;
    background-color:rgb(40,52,70);
    color:#fff;
}

.collapse-btn {
    height: 42px;
    text-align: center;
    line-height: 42px;
    cursor: pointer;
}
.sidebar::-webkit-scrollbar {
    width: 0;
}
.sidebar-el-menu:not(.el-menu--collapse) {
    width: 150px;
}
.sidebar > ul {
    height: 100%;
}
.header {
    box-sizing: border-box;
    width: 100%;
    height:70px;
    font-size: 22px;
    color: #fff;
}
.header .logo {
    float:left;
    text-align: center;
    width: 100%;
    line-height: 70px;
}
</style>
