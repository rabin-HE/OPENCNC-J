<template>
    <div class="device-view" :class="props.cnc.navigation.select">
        <div class="device-main">
            <div class="device-box">
                <div class="box-title">
                    <el-text class="cnc">
                        <el-icon><QuestionFilled /></el-icon>
                        <span>设备列表</span>
                    </el-text>
                </div>
                <div class="device-item">
                    <el-row class="cnc" :gutter="20">
                        <el-col class="cnc" :span="8" v-for="(item, index) in props.cnc.device.ips" :key="index">
                            <div class="grid-content" @click="onSelectDevice(item)">
                                <div class="device-name">{{item.name}}</div>
                                <div class="device-ip">{{item.ip}}</div>
                            </div>
                        </el-col>
                    </el-row>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import {defineComponent, onBeforeMount, onMounted, onBeforeUnmount, onUnmounted} from "vue";
import * as icons from "@element-plus/icons";
import {ElMessage} from "element-plus";
export default defineComponent({
    name: "DeviceMain",
    emits: [],
    props: ["cnc"],
    components: {},
    setup(props, context) {

        function onSelectDevice(device: any){
            (window as any).go.StartWindows.Api.DeviceRequest(device.ip + ":" + props.cnc.device.control.port, "/query/inifields/", "GET", {}).then((response: any)=>{
                console.log(response);
                if(response.code === 0){
                    if(response.MACHINE){
                        props.cnc.device.ip = device.ip;
                        (window as any).runtime.EventsEmit("event_message", {type: "connected_device"});
                    }else{
                        ElMessage.closeAll();
                        ElMessage({
                            message: "设备连接失败，请检查后重新尝试",
                            type: "warning",
                            customClass: "cnc"
                        });
                    }
                }else{
                    ElMessage.closeAll();
                    ElMessage({
                        message: "设备连接失败，请检查后重新尝试",
                        type: "warning",
                        customClass: "cnc"
                    });
                }
            });
        }

        onBeforeMount(() => {});

        onMounted(() => {});

        onBeforeUnmount(() => {});

        onUnmounted(() => {});

        return {
            props,
            icons,
            onSelectDevice
        }
    }
});
</script>

<style scoped>
.device-view{
    width: 100%;
    height: 100%;
    background-color: rgba(30, 31, 34, 1);
    overflow-y: auto;
    display: none;
}
.device-view.device{
    display: block;
}
.device-view .device-main{
    padding: 30px;
    min-height: 1000px;
}
.device-view .device-main .device-box{
    width: 850px;
    padding: 30px;
    background-color: rgba(43, 45, 48, .5);
    border-radius: 4px;
    margin: 0 auto;
}
.device-view .device-main .device-box .box-title{
    width: 100%;
    margin-bottom: 10px;
}
.device-view .device-main .device-box .box-title:deep(.el-text){
    color: #999999;
}
.device-view .device-main .device-box .device-item{
    width: 100%;
}
.device-view .device-main .device-box .device-item .grid-content{
    width: 100%;
    height: 80px;
    padding: 20px;
    background-color: rgba(43, 45, 48, .5);
    border-radius: 4px;
    position: relative;
}
.device-view .device-main .device-box .device-item .grid-content:hover{
    background-color: rgba(43, 45, 48, .8);
    cursor: pointer;
}
.device-view .device-main .device-box .device-item .grid-content .device-name{
    width: 100%;
    height: 20px;
    line-height: 20px;
    color: #999999;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}
.device-view .device-main .device-box .device-item .grid-content:hover .device-name{
    color: #ffffff;
}
.device-view .device-main .device-box .device-item .grid-content .device-ip{
    width: 100%;
    height: 20px;
    line-height: 20px;
    color: #666666;
}
</style>
