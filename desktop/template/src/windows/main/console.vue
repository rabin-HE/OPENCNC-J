<template>
    <div class="console-view" :class="props.cnc.navigation.select">
        <div class="console-item">
            <div class="console-main-item">

            </div>
            <div class="console-main-item">
                <div class="console-main-header">
                    <div class="console-header-item"></div>
                    <div class="console-header-item"></div>
                    <div class="console-header-item"></div>
                </div>
                <div class="console-main-footer"></div>
            </div>
        </div>
        <div class="console-item">
            <div class="console-right">
                <div class="console-right-item">
                    <div class="title">
                        <el-text class="cnc">
                            <el-icon><ScaleToOriginal /></el-icon>
                            <span>步长(mm)</span>
                        </el-text>
                    </div>
                    <div class="box">
                        <div class="step">
                            <el-row class="cnc" :gutter="5">
                                <el-col :span="4">
                                    <div class="grid-content">0.01</div>
                                </el-col>
                                <el-col :span="4">
                                    <div class="grid-content">0.05</div>
                                </el-col>
                                <el-col :span="4">
                                    <div class="grid-content">0.1</div>
                                </el-col>
                                <el-col :span="4">
                                    <div class="grid-content">0.5</div>
                                </el-col>
                                <el-col :span="4">
                                    <div class="grid-content select">1</div>
                                </el-col>
                                <el-col :span="4">
                                    <div class="grid-content">5</div>
                                </el-col>
                                <el-col :span="4">
                                    <div class="grid-content">10</div>
                                </el-col>
                                <el-col :span="4">
                                    <div class="grid-content">20</div>
                                </el-col>
                                <el-col :span="4">
                                    <div class="grid-content">50</div>
                                </el-col>
                                <el-col :span="4">
                                    <div class="grid-content">100</div>
                                </el-col>
                                <el-col :span="4">
                                    <div class="grid-content">150</div>
                                </el-col>
                                <el-col :span="4">
                                    <div class="grid-content">360</div>
                                </el-col>
                            </el-row>
                        </div>
                    </div>
                </div>
                <div class="console-right-item">
                    <div class="title">
                        <el-text class="cnc">
                            <el-icon><Operation /></el-icon>
                            <span>速度(min)</span>
                        </el-text>
                    </div>
                    <div class="box">
                        <div class="speed-box">
                            <div class="speed-box-item">
                                <div class="item">点动速度</div>
                                <div class="item">
                                    <el-slider class="cnc" size="small" show-input :show-input-controls="false" />
                                </div>
                            </div>
                            <div class="speed-box-item">
                                <div class="item">最大速度</div>
                                <div class="item">
                                    <el-slider class="cnc" size="small" show-input :show-input-controls="false" />
                                </div>
                            </div>
                            <div class="speed-box-item">
                                <div class="item">点转速度</div>
                                <div class="item">
                                    <el-slider class="cnc" size="small" show-input :show-input-controls="false" />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="console-right-item">
                    <div class="title">
                        <el-text class="cnc">
                            <el-icon><Coordinate /></el-icon>
                            <span>调试键盘</span>
                        </el-text>
                    </div>
                    <div class="box">
                        <div class="keyboard">
                            <el-row class="cnc" :gutter="5">
                                <el-col :span="6">
                                    <div class="grid-content none"></div>
                                </el-col>
                                <el-col :span="6">
                                    <div class="grid-content">
                                        <div class="icon">
                                            <el-icon><ArrowUpBold /></el-icon>
                                        </div>
                                        <div class="name">Y+</div>
                                    </div>
                                </el-col>
                                <el-col :span="6">
                                    <div class="grid-content none"></div>
                                </el-col>
                                <el-col :span="6">
                                    <div class="grid-content">
                                        <div class="icon">
                                            <el-icon><ArrowUpBold /></el-icon>
                                        </div>
                                        <div class="name">Z+</div>
                                    </div>
                                </el-col>
                                <el-col :span="6">
                                    <div class="grid-content">
                                        <div class="icon">
                                            <el-icon><ArrowLeftBold /></el-icon>
                                        </div>
                                        <div class="name">X-</div>
                                    </div>
                                </el-col>
                                <el-col :span="6">
                                    <div class="grid-content">
                                        <div class="icon">
                                            <el-icon><ArrowDownBold /></el-icon>
                                        </div>
                                        <div class="name">Y-</div>
                                    </div>
                                </el-col>
                                <el-col :span="6">
                                    <div class="grid-content">
                                        <div class="icon">
                                            <el-icon><ArrowRightBold /></el-icon>
                                        </div>
                                        <div class="name">X+</div>
                                    </div>
                                </el-col>
                                <el-col :span="6">
                                    <div class="grid-content">
                                        <div class="icon">
                                            <el-icon><ArrowDownBold /></el-icon>
                                        </div>
                                        <div class="name">Z-</div>
                                    </div>
                                </el-col>
                                <el-col :span="6">
                                    <div class="grid-content">
                                        <div class="icon">
                                            <el-icon><ArrowLeftBold /></el-icon>
                                        </div>
                                        <div class="name">B-</div>
                                    </div>
                                </el-col>
                                <el-col :span="6">
                                    <div class="grid-content">
                                        <div class="icon">
                                            <el-icon><ArrowRightBold /></el-icon>
                                        </div>
                                        <div class="name">B+</div>
                                    </div>
                                </el-col>
                                <el-col :span="6">
                                    <div class="grid-content">
                                        <div class="icon">
                                            <el-icon><ArrowLeftBold /></el-icon>
                                        </div>
                                        <div class="name">C-</div>
                                    </div>
                                </el-col>
                                <el-col :span="6">
                                    <div class="grid-content">
                                        <div class="icon">
                                            <el-icon><ArrowRightBold /></el-icon>
                                        </div>
                                        <div class="name">C+</div>
                                    </div>
                                </el-col>
                            </el-row>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import {defineComponent, onBeforeMount, onMounted, onBeforeUnmount, onUnmounted} from "vue";
import * as icons from "@element-plus/icons";
export default defineComponent({
    name: "ConsoleMain",
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
.console-view{
    width: 100%;
    height: 100%;
    background-color: rgba(30, 31, 34, 1);
    z-index: 100;
    display: none;
}
.console-view.console{
    display: block;
}
.console-view .console-item{
    width: calc(100% - 380px);
    height: 100%;
    display: inline-block;
    vertical-align: top;
}
.console-view .console-item:last-child{
    width: 380px;
    padding: 10px;
    background-color: rgba(43, 45, 48, 1);
    border-left: solid 1px rgba(30, 31, 34, 1);
    overflow-y: auto;
}
.console-view .console-item .console-main-item{
    width: 100%;
    height: calc(100% - 240px);
}
.console-view .console-item .console-main-item:last-child{
    height: 240px;
    background-color: rgba(43, 45, 48, .5);
    border-top: solid 1px rgba(30, 31, 34, .8);
}
.console-view .console-item .console-main-item .console-main-header{
    width: 100%;
    height: 40px;
    background-color: rgba(43, 45, 48, .5);
    border-bottom: solid 1px rgba(30, 31, 34, .8);
}
.console-view .console-item .console-main-item .console-main-header .console-header-item{
    width: 33.33%;
    height: 40px;
    display: inline-block;
    vertical-align: top;
}
.console-view .console-item .console-right{
    width: 100%;
    min-height: 100%;
}
.console-view .console-item .console-right .console-right-item{
    width: 100%;
    margin-bottom: 10px;
}
.console-view .console-item .console-right .console-right-item .title{
    width: 100%;
    margin-bottom: 10px;
}
.console-view .console-item .console-right .console-right-item .box{
    width: 100%;
}
.console-view .console-item .console-right .console-right-item .box:deep(.el-row.cnc .el-col){
    margin-bottom: 5px;
}
.console-view .console-item .console-right .console-right-item .box .step .grid-content{
    width: 100%;
    height: 30px;
    line-height: 30px;
    text-align: center;
    background-color: rgba(30, 31, 34, .3);
    border-radius: 4px;
    color: #999999;
}
.console-view .console-item .console-right .console-right-item .box .step .grid-content:hover{
    cursor: pointer;
    color: #ffffff;
    background-color: rgba(30, 31, 34, .5);
}
.console-view .console-item .console-right .console-right-item .box .step .grid-content.select{
    background-color: #5e4eff;
    color: #ffffff;
}
.console-view .console-item .console-right .console-right-item .box .speed-box{
    width: 100%;
    background-color: rgba(30, 31, 34, .3);
    border-radius: 4px;
    padding: 15px;
}
.console-view .console-item .console-right .console-right-item .box .speed-box .speed-box-item{
    width: 100%;
    margin-bottom: 15px;
}
.console-view .console-item .console-right .console-right-item .box .speed-box .speed-box-item:last-child{
    margin-bottom: 0;
}
.console-view .console-item .console-right .console-right-item .box .speed-box .speed-box-item .item{
    width: calc(100% - 65px);
    height: 24px;
    display: inline-block;
    vertical-align: top;
}
.console-view .console-item .console-right .console-right-item .box .speed-box .speed-box-item .item:first-child{
    width: 65px;
    line-height: 24px;
    font-size: 12px;
    color: #666666;
}
.console-view .console-item .console-right .console-right-item .box .keyboard{
    width: 100%;
    padding: 0 15px;
}
.console-view .console-item .console-right .console-right-item .box .keyboard .grid-content{
    width: 100%;
    height: 65px;
    text-align: center;
    background-color: rgba(30, 31, 34, .3);
    border: solid 1px rgba(30, 31, 34, .5);
    border-radius: 4px;
    color: #999999;
    padding: 10px;
}
.console-view .console-item .console-right .console-right-item .box .keyboard .grid-content.none{
    background-color: rgba(30, 31, 34, 0);
    border: solid 1px rgba(30, 31, 34, 0);
}
.console-view .console-item .console-right .console-right-item .box .keyboard .grid-content:not(.none):hover{
    cursor: pointer;
    color: #ffffff;
    background-color: #5e4eff;
}
.console-view .console-item .console-right .console-right-item .box .keyboard .grid-content .icon{
    width: 100%;
    height: 25px;
    line-height: 25px;
    text-align: center;
    color: #5e4eff;
    font-size: 16px;
}
.console-view .console-item .console-right .console-right-item .box .keyboard .grid-content:not(.none):hover .icon{
    color: #ffffff;
}
.console-view .console-item .console-right .console-right-item .box .keyboard .grid-content .name{
    width: 100%;
    height: 20px;
    line-height: 20px;
    color: #999999;
}
.console-view .console-item .console-right .console-right-item .box .keyboard .grid-content:not(.none):hover .name{
    color: #ffffff;
}
</style>
