<template>
    <div class="footer-main">
        <div class="footer-item"></div>
        <div class="footer-item"></div>
        <div class="footer-item">
            <div class="item">
                <el-text class="cnc">
                    <el-icon><ChatLineRound /></el-icon>
                    <span>{{props.cnc.device.control.status ? "已连接" : "未连接"}}</span>
                </el-text>
            </div>
            <div class="item">
                <el-text class="cnc">
                    <el-icon><ChatDotRound /></el-icon>
                    <span>{{props.cnc.device.message.status ? "已连接" : "未连接"}}</span>
                </el-text>
            </div>
            <div class="item">
                <el-text class="cnc">
                    <el-icon><Bell /></el-icon>
                    <span>GEEKCNC 1.0.0</span>
                </el-text>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import {defineComponent, onBeforeMount, onMounted, onBeforeUnmount, onUnmounted} from "vue";
import * as icons from "@element-plus/icons";
export default defineComponent({
    name: "FooterCommon",
    emits: [],
    props: ["cnc"],
    components: {},
    setup(props, context) {

        onBeforeMount(() => {});

        onMounted(() => {});

        onBeforeUnmount(() => {});

        onUnmounted(() => {});

        return {
            props,
            icons
        }
    }
});
</script>

<style scoped>
.footer-main{
    width: 100%;
    height: 25px;
}
.footer-main .footer-item{
    width: 33.33%;
    height: 25px;
    color: #666666;
    display: inline-block;
    vertical-align: top;
}
.footer-main .footer-item:last-child{
    text-align: right;
}
.footer-main .footer-item .item{
    width: auto;
    height: 25px;
    line-height: 25px;
    display: inline-block;
    vertical-align: top;
    padding: 0 10px;
}
.footer-main .footer-item .item:hover{
    background-color: rgba(43, 45, 48, 1);
}
.footer-main .footer-item .item:deep(.el-text){
    height: 25px;
    line-height: 25px;
}
</style>
