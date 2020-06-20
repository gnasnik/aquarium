<template>
    <div class="header">
       
    </div>
</template>
<script>
import bus from '../common/bus';
export default {
    data() {
        return {
            collapse: false,
            fullscreen: false,
        };
    },
    mounted() {
        if (document.body.clientWidth < 1200) {
            this.collapseChage();
        }
    }
};
</script>
<style scoped>
.header {
    position: relative;
    box-sizing: border-box;
    width: 100%;
    height: 70px;
    font-size: 22px;
    color: #fff;
}
.collapse-btn {
    float: left;
    padding: 0 21px;
    cursor: pointer;
    line-height: 70px;
}
.header .logo {
    float: left;
    width: 250px;
    line-height: 70px;
}
</style>
