<template>
    <div>
        <div class="container">
            <div class="editor-header">
            <el-input class="editor-input" v-model="form.name" placeholder="New Algorithm Name" maxlength="100"></el-input>
            <el-button class="editor-btn" type="primary" @click="submit">Submit</el-button>
            <el-button class="editor-btn" type="normal" @click="cancel">Cancel</el-button>
            </div>
            <el-input class="editor-textarea" type="textarea" autosize placeholder="Algorithm Description" v-model="form.description"></el-input>
            <MonacoEditor
                height="300"
                width="100%"
                class="vs"
                style="position: absolute; overflow: hidden; left: 32px; width: 1038px; height: 453px;"
                language="javascript"
                :code="form.script"
                :editorOptions="options"
                @mounted="onMounted"
                @codeChange="onCodeChange"
                >
            </MonacoEditor>
        </div>
    </div>
</template>

<script>
    import bus from '../common/bus';
    import { addAlgorithmReq } from '../../api/algorithm';
    import MonacoEditor from '../../vue-monaco-editor/src/Monaco';
    export default {
        name: 'editor',
        data: function(){
            return {
                title:'New Algorithm Name',
                content: '',
                form:{
                    script:'// type your code \n function main() {\n // console.log("hello aquarium"); \n }',
                },
                token:'',
                // code: '',
                editor:null,
                options: {
                    theme: "vs",
                    selectOnLineNumbers: true,
                    roundedSelection: false,
                    readOnly: false,
                    automaticLayout: true,
                    glyphMargin: true,
                    showFoldingControls: "always",
                    formatOnPaste: true,
                    formatOnType: true,
                    folding: true,
                },
            }
        },
        components: {
            MonacoEditor
        },
        created(){
            this.token = localStorage.getItem("token");
            if (this.$route.params) {
                this.form = this.$route.params;
                console.log("form-->",this.form )
            }
        },
        methods: {
            onMounted(editor) {
            console.log('after mount!', editor, editor.getValue(), editor.getModel());
            this.editor = editor;
            },
            onCodeChange(editor) {
                console.log('code changed!', 'code:' + this.editor.getValue());
                this.form.script = this.editor.getValue();
            },
            submit(){
                if (this.form.script == "") {
                    this.form.script = this.code;
                }
                console.log("------->>", this.form);
                addAlgorithmReq(this.form, this.token).then(res => {
                if (res.success) {
                    this.$message.success(`Success`);
                    bus.$emit('close_current_tags');
                }else{
                    this.$message.error(res.msg || "unkown err");
                }  
                })
            },
            cancel(){
                bus.$emit('close_current_tags');
            },
        },
    }
</script>
<style scoped>
    .container{
        padding: 20px 20px;
        height: 100%;
    }
    .editor-input{ 
        width:30%; 
        float: left;
    } 
    .editor-btn{
        float: left;
        margin-left: 10px;
    }
    .editor-textarea{
        padding-top: 10px;
        padding-bottom: 10px;
    }
</style>