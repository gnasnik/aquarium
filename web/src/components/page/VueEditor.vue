<template>
    <div>
        <div class="container">
            <div class="editor-header">
            <el-input class="editor-input" v-model="title" placeholder="New Algorithm Name"  maxlength="10"></el-input>
            <el-button class="editor-btn" type="primary" @click="submit">Submit</el-button>
            <el-button class="editor-btn" type="normal" @click="submit">Cancel</el-button>
            </div>
            <el-input class="editor-textarea" type="textarea" autosize placeholder="Algorithm Description" v-model="textarea"></el-input>
            <MonacoEditor
                height="300"
                width="100%"
                class="vs"
                style="position: absolute; overflow: hidden; left: 32px; width: 1038px; height: 453px;"
                language="javascript"
                :code="code"
                :editorOptions="options"
                @mounted="onMounted"
                @codeChange="onCodeChange"
                >
            </MonacoEditor>
        </div>
    </div>
</template>

<script>
    import MonacoEditor from '../../vue-monaco-editor/src/Monaco';
    export default {
        name: 'editor',
        data: function(){
            return {
                title:'New Algorithm Name',
                content: '',
                textarea:'',
                code: '// type your code \n',
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
        methods: {
            onMounted(editor) {
            console.log('after mount!', editor, editor.getValue(), editor.getModel());
            this.editor = editor;
            },
            onCodeChange(editor) {
            console.log('code changed!', 'code:' + this.editor.getValue());
            },
        },
        // created(){
        //     this.options = {
        //         selectOnLineNumbers: false
        //     };
        // }
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