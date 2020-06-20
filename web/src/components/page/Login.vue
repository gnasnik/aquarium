<template>
    <div class="login-wrap">
        <div class="ms-login">
            <div class="ms-title">Aquarium</div>
            <el-form :model="param" :rules="rules" ref="login" label-width="0px" class="ms-content">
                <el-form-item prop="username">
                    <el-input v-model="param.username" placeholder="username">
                        <el-button slot="prepend" icon="el-icon-lx-people"></el-button>
                    </el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input
                        type="password"
                        placeholder="password"
                        v-model="param.password"
                        @keyup.enter.native="submitForm()"
                    >
                        <el-button slot="prepend" icon="el-icon-lx-lock"></el-button>
                    </el-input>
                </el-form-item>
                <div class="login-btn">
                    <el-button type="primary" @click="submitForm()">Login</el-button>
                </div>
                <!-- <p class="login-tips">Tips : 用户名和密码随便填。</p> -->
            </el-form>
        </div>
    </div>
</template>

<script>
import { loginReq } from '../../api/index';
export default {
    data: function() {
        return {
            param: {
                username: '1',
                password: '12334434',
                login_type:"phone"
            },
            rules: {
                username: [{ required: true, message: 'username is required', trigger: 'blur' }],
                password: [{ required: true, message: 'password is required', trigger: 'blur' }],
            },
        };
    },
    methods: {
        submitForm() {
            this.$refs.login.validate(valid => {
                if (valid) {
                    this.param.user_id = parseInt(this.param.username);
                    loginReq(this.param).then(res => {
                        if (res.success) {
                            this.$message.success('Login Success');
                            localStorage.setItem('ms_username', this.param.username);
                            this.$router.push('/');
                        }else{
                            this.$message.error(res.msg?res.msg:"unknown err");
                        }
                    });
                    // this.$message.success('Login Success');
                    // localStorage.setItem('ms_username', this.param.username);
                    // this.$router.push('/');
                } else {
                    this.$message.error('Please enter username/password');
                    console.log('error submit!!');
                    return false;
                }
            });
        },

        login(){

        },
    },
};
</script>

<style scoped>
.login-wrap {
    position: relative;
    width: 100%;
    height: 100%;
    background-image: url(../../assets/img/login-bg.jpg);
    background-size: 100%;
}
.ms-title {
    width: 100%;
    line-height: 50px;
    text-align: center;
    font-size: 30px;
    color: #000;
    border-bottom: 1px solid #ddd;
}
.ms-login {
    position: absolute;
    left: 50%;
    top: 50%;
    width: 350px;
    margin: -190px 0 0 -175px;
    border-radius: 5px;
    background: rgba(255, 255, 255, 0.3);
    overflow: hidden;
}
.ms-content {
    padding: 30px 30px;
}
.login-btn {
    text-align: center;
}
.login-btn button {
    width: 100%;
    height: 36px;
    margin-bottom: 10px;
}
.login-tips {
    font-size: 12px;
    line-height: 30px;
    color: #fff;
}
</style>