<template>
    <el-dialog class="cnc" v-model="props.cnc.header.dialog.status" :title="props.cnc.header.dialog.config.title" :width="props.cnc.header.dialog.config.width" draggable :modal="true" title="" :show-close="props.cnc.header.dialog.config.close" :before-close="dialogClose" :lock-scroll="true" :closeOnClickModal="false" :closeOnPressEscape="false" :destroy-on-close="true">
        <div class="new-device-dialog">
            <el-form class="cnc" :model="props.cnc.header.dialog.form" label-width="80px">
                <el-form-item label="设备IP地址">
                    <el-input class="cnc" v-model="props.cnc.header.dialog.form.ip" placeholder="请输入设备的IP地址" maxlength="140" autocomplete="off" spellcheck="false" style="width: 200px" />
                </el-form-item>
                <el-form-item label="">
                    <el-button color="#5e4eff" class="cnc" :loading="props.cnc.header.dialog.form.loading" type="primary" @click="onDevice">连接设备</el-button>
                </el-form-item>
            </el-form>
        </div>
    </el-dialog>
</template>

<script lang="ts">
import {defineComponent, onBeforeUnmount, onMounted} from "vue";
import {ElMessage} from "element-plus";
export default defineComponent({
    name: "NewDeviceDialog",
    emits: [],
    props: ["cnc"],
    components: {},
    setup(props, context) {

        function onDevice(){
            if(props.cnc.header.dialog.form.ip === "" && (props.cnc.header.dialog.form.ip.replace(/\n|\r/g, "")).trim().length === 0){
                ElMessage.closeAll();
                ElMessage({
                    message: "设备IP地址错误",
                    type: "warning",
                    customClass: "cnc"
                });
                return;
            }
            props.cnc.header.dialog.form.loading = true;
            (window as any).go.StartWindows.Api.DeviceRequest(props.cnc.header.dialog.form.ip + ":" + props.cnc.device.control.port, "/query/inifields/", "GET", {}).then((response: any)=>{
                if(response.code === 0){
                    if(response.MACHINE){
                        props.cnc.device.ip = props.cnc.header.dialog.form.ip;
                        if(props.cnc.device.ips.length > 0){
                            let check = false;
                            props.cnc.device.ips.forEach((item: any, index: any, array: any)=>{
                                if(item.ip === props.cnc.device.ip){
                                    check = true;
                                }
                            });
                            if(!check){
                                props.cnc.device.ips.push({name: response.MACHINE, ip: props.cnc.header.dialog.form.ip});
                                localStorage.setItem("cnc:device:ips", JSON.stringify(props.cnc.device.ips));
                            }
                        }else{
                            props.cnc.device.ips.push({name: response.MACHINE, ip: props.cnc.header.dialog.form.ip});
                            localStorage.setItem("cnc:device:ips", JSON.stringify(props.cnc.device.ips));
                            console.log(props.cnc.device.ips);
                        }
                        (window as any).runtime.EventsEmit("event_message", {type: "connected_device"});
                        dialogClose(false);
                    }else{
                        props.cnc.header.dialog.form.loading = false;
                        ElMessage.closeAll();
                        ElMessage({
                            message: "设备连接失败，请检查后重新尝试",
                            type: "warning",
                            customClass: "cnc"
                        });
                    }
                }else{
                    props.cnc.header.dialog.form.loading = false;
                    ElMessage.closeAll();
                    ElMessage({
                        message: "设备连接失败，请检查后重新尝试",
                        type: "warning",
                        customClass: "cnc"
                    });
                }
            });
        }

        function dialogClose(close: any){
            if(close){
                close();
            }
            props.cnc.header.dialog.status = false;
            setTimeout(()=>{
                props.cnc.header.dialog.config.type = "";
                props.cnc.header.dialog.form = {
                    loading: false
                };
            }, 20);
        }

        onMounted(() => {});

        onBeforeUnmount(() => {});

        return {
            props,
            onDevice,
            dialogClose
        }
    }
});
</script>

<style scoped>
.new-device-dialog{
    width: 100%;
    padding: 20px;
}
</style>