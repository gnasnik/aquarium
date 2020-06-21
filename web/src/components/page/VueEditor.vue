<template>
    <div>
        <div class="container">
            <div class="editor-header">
            <el-input class="div-inline" v-model="title" placeholder="New Algorithm Name" width="200px"></el-input>
            <el-button class="div-inline" type="primary" @click="submit">Submit</el-button>
            <el-button class="div-inline" type="normal" @click="submit">Cancel</el-button>
            </div>
            <el-input type="textarea" autosize placeholder="Algorithm Description" v-model="textarea1"></el-input>
            <el-input
                type="textarea"
                :rows="2"
                placeholder="请输入内容"
                v-model="textarea">
                </el-input>
            <!-- <quill-editor ref="myTextEditor" v-model="content" :options="editorOption"></quill-editor> -->
            <el-button class="editor-btn" type="primary" @click="submit">提交</el-button>
        </div>
    </div>
</template>

<script>
    import 'quill/dist/quill.core.css';
    import 'quill/dist/quill.snow.css';
    import 'quill/dist/quill.bubble.css';
    import { quillEditor } from 'vue-quill-editor';
    export default {
        name: 'editor',
        data: function(){
            return {
                title:'New Algorithm Name',
                content: '',
                editorOption: {
                    placeholder: 'Hello World'
                }
            }
        },
        components: {
            quillEditor
        },
        methods: {
            onEditorChange({ editor, html, text }) {
                this.content = html;
            },
            submit(){
                console.log(this.content);
                this.$message.success('提交成功！');
            }
        }
    }
</script>
<style scoped>
    .div-inline{ display:inline} 
    .editor-btn{
        display:inline-block;;
        margin-top: 20px;
    }
</style>