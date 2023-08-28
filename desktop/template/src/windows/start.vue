<template>
    <div class="cnc-main">
        <div class="main-item">
            <HeaderCommon ref="headerCommon" :cnc="props.cnc" />
        </div>
        <div class="main-item">
            <div class="main-box">
                <div class="main-box-item">
                    <NavigationCommon ref="navigationCommon" :cnc="props.cnc" />
                </div>
                <div class="main-box-item">
                    <DeviceMain ref="deviceMain" :cnc="props.cnc" />
                    <ConsoleMain ref="consoleMain" :cnc="props.cnc" />
                    <ProgramMain ref="programMain" :cnc="props.cnc" />
                    <PluginMain ref="pluginMain" :cnc="props.cnc" />
                    <BladeMain ref="bladeMain" :cnc="props.cnc" />
                    <SettingsMain ref="settingsMain" :cnc="props.cnc" />
                </div>
            </div>
        </div>
        <div class="main-item">
            <FooterCommon ref="footerCommon" :cnc="props.cnc" />
        </div>
    </div>
</template>

<script lang="ts">
import {defineComponent, nextTick, onBeforeMount, onBeforeUnmount, onMounted, onUnmounted} from "vue";
import {ElLoading} from "element-plus";
import * as icons from "@element-plus/icons";
import NoSleep from "nosleep.js";
import HeaderCommon from "./common/header.vue";
import FooterCommon from "./common/footer.vue";
import NavigationCommon from "./common/navigation.vue";
import DeviceMain from "./main/device.vue";
import ConsoleMain from "./main/console.vue";
import ProgramMain from "./main/program.vue";
import PluginMain from "./main/plugin.vue";
import BladeMain from "./main/blade.vue";
import SettingsMain from "./main/settings.vue";
export default defineComponent({
    name: "Start",
    emits: [],
    props: ["cnc"],
    components: {
        HeaderCommon,
        FooterCommon,
        NavigationCommon,
        DeviceMain,
        ConsoleMain,
        ProgramMain,
        PluginMain,
        BladeMain,
        SettingsMain
    },
    setup(props, context) {

        // 消息事件
        (window as any).runtime.EventsOn("event_message", (message: any) => {
            if(message.type && message.type === "connected_device"){
                onConnectedDevice();
            }
            if(message.type && message.type === "disconnect_device"){
                onDisconnectDevice();
            }
        });

        // 连接设备
        function onConnectedDevice(){
            if(props.cnc.device.ip !== ""){
                onConnectedDeviceControl();
            }
        }

        // 连接设备控制消息服务
        function onConnectedDeviceControl(){
            if(!props.cnc.device.control.status){
                props.cnc.device.control.socket = new WebSocket("ws://" + props.cnc.device.ip + ":" + props.cnc.device.control.port + "/websocket/", undefined);
                props.cnc.device.control.socket.onopen = function () {
                    props.cnc.device.control.status = true;
                    onConnectedDeviceMessage();
                }
                props.cnc.device.control.socket.onmessage = function (message: any) {
                    let message_json = JSON.parse(message.data);
                    console.log(message_json);
                }
                props.cnc.device.control.socket.onerror = function () {
                    onDisconnectDevice();
                }
                props.cnc.device.control.socket.onclose = function () {
                    onDisconnectDevice();
                }
            }
        }

        // 连接设备消息服务
        function onConnectedDeviceMessage(){
            if(!props.cnc.device.message.status){
                props.cnc.device.message.socket = new WebSocket("ws://" + props.cnc.device.ip + ":" + props.cnc.device.message.port + "/websocket/", undefined);
                props.cnc.device.message.socket.onopen = function () {
                    props.cnc.device.message.status = true;
                    props.cnc.navigation.select = "console";
                }
                props.cnc.device.message.socket.onmessage = function (message: any) {
                    let message_json = JSON.parse(message.data);
                    if(message_json.type && message_json.name){
                        if(message_json.name === "basicInfo"){
                            props.cnc.device.message.data.basicInfo = message_json.datas;
                        }
                        if(message_json.name === "errorInfo"){
                            console.log(message_json.datas.text);
                        }
                    }
                }
                props.cnc.device.message.socket.onerror = function () {
                    onDisconnectDevice();
                }
                props.cnc.device.message.socket.onclose = function () {
                    onDisconnectDevice();
                }
            }
        }

        // 断开设备
        function onDisconnectDevice(){
            if(props.cnc.device.control.status){
                props.cnc.device.control.status = false;
                props.cnc.device.control.socket.close();
            }
            if(props.cnc.device.message.status){
                props.cnc.device.message.status = false;
                props.cnc.device.message.socket.close();
            }
            props.cnc.navigation.select = "device";
        }

        onBeforeMount(() => {
            props.cnc.loading = ElLoading.service({
                lock: true,
                background: "rgba(0, 0, 0, .01)",
                fullscreen: true
            });
        });

        onMounted(() => {
            nextTick(()=>{
                if(!props.cnc.sleep){
                    // 禁止屏幕休眠
                    props.cnc.sleep = new NoSleep();
                    props.cnc.sleep.enable();
                }
                let ips = localStorage.getItem("cnc:device:ips");
                if(ips){
                    props.cnc.device.ips = JSON.parse(ips);
                }
                props.cnc.loading.close();
            });
        });

        onBeforeUnmount(()=>{
            onDisconnectDevice();
        })

        onUnmounted(() => {});

        return {
            props,
            icons
        }
    }
});
</script>

<style scoped>
.cnc-main{
    width: 100%;
    height: 100%;
    position: fixed;
    z-index: 1;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
}
.cnc-main .main-item{
    width: 100%;
    height: calc(100% - 65px);
}
.cnc-main .main-item:first-child{
    width: 100%;
    height: 40px;
    background-color: rgba(43, 45, 48, 1);
    border-top: solid 1px rgba(30, 31, 34, 1);
}
.cnc-main .main-item:last-child{
    width: 100%;
    height: 25px;
    background-color: rgba(43, 45, 48, .8);
    border-top: 1px solid rgba(30, 31, 34, 1);
}
.cnc-main .main-item .main-box{
    width: 100%;
    height: 100%;
    border-top: 1px solid rgba(30, 31, 34, 1);
}
.cnc-main .main-item .main-box .main-box-item{
    width: calc(100% - 40px);
    height: 100%;
    display: inline-block;
    vertical-align: top;
}
.cnc-main .main-item .main-box .main-box-item:first-child{
    width: 40px;
    height: 100%;
    display: inline-block;
    vertical-align: top;
    background-color: rgba(43, 45, 48, 1);
    border-right: solid 1px rgba(30, 31, 34, 1);
}
</style>
