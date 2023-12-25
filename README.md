![](/Picture/machine.jpg)

## **OPENCNC-J**  [📖ENGLISH_README](/README-EN.md)

* 简介
  * 这是一个基于Linuxcnc的开源整合项目, 目的为让电气工程师能更方便的使用Linuxcnc + ethercat与CIA402协议 + 树莓派(或X86 - Debian)的开发组合
* 当前版本 - 希望有更多的开发者加入
  * 0.1.0
* 个人的视频发布主页 - 详细的项目内容
  * [bilibili - OPENCNC-J开发过程合集](https://space.bilibili.com/341589947/channel/collectiondetail?sid=1918828)
* 注意 & 警告!!!
  * 此项目为开源项目, 不提供任何保障, 不提供任何赔付, 项目中涉及电气安全警告, 请自行负责
  * 此项目遵循GPL-3.0开源协议

## **安装说明 - 以树莓派安装为例**

安装[Linuxcnc](https://linuxcnc.org/)

1. 下载官方安装包 - LinuxCNC 2.9.1[Raspberry Pi OS based on Debian 12 Bookworm](https://www.linuxcnc.org/iso/linuxcnc-2.9.1-bookworm-rpi4.img.xz)
2. 制作SD卡系统 - 建议使用[树莓派官方](https://www.raspberrypi.com/software)的SD卡系统制作工具[Windows版本](https://downloads.raspberrypi.org/imager/imager_latest.exe) &[Ubuntu版本](https://downloads.raspberrypi.org/imager/imager_latest_amd64.deb)
3. 插卡启动树莓派
4. 配置Linuxcnc - [留白]

安装依赖包 - [参考安装ethercat范例](https://forum.linuxcnc.org/ethercat/45336-ethercat-installation-from-repositories-how-to-step-by-step), CIA402

1. ethercat安装
   1. 打开terminal:`sudo apt install ethercat-master libethercat-dev linuxcnc-ethercat`
   2. 如果安装失败,请参考范例
   3. 查看网口的MAC地址:`ip a`
      1. 记录MAC地址例如:`xx:aa:yy:zz:bb:cc`
   4. 输入(配置)到ethercat.conf文件
      1. `sudo geany /etc/ethercat.conf`
      2. 用geany修改文件
         1. `MASTER0_DEVICE="xx:aa:yy:zz:bb:cc"` - 替换为你刚才获取的MAC地址`DEVICE_MODULES="generic"`
   5. 将ethercat服务设置为开机启动
      1. `sudo systemctl enable ethercat.service`
         `sudo systemctl start ethercat.service`
         `sudo systemctl status ethercat.service`
         `sudo chmod 666 /dev/EtherCAT0`
      2. ethercat测试命令
         1. `ethercat slaves` - 看看是否报错, 确认安装成功
         2. 重启
2. 安装CIA402
   1. 回到用户目录:`cd`
   2. `git clone https://github.com/dbraun1981/hal-cia402` - 如果报错请确认是否已经安装git工具
      1. `sudo apt install git` - git安装命令
   3. 进入hal-cia402文件夹:`cd hal-cia402`
   4. 编译安装402工具包:`sudo halcompile --install cia402.comp`

至此完成最困难的部分👍

## **驱动器参数设置**

* 启动时显示AL221报警[编码器未回零]
  * 设置零点: 点击驱动器面板操作 → 切换到[AF.CEN] → 按S → 再按S → 设置成功  #参照驱动器说明书 P39 - 4.3.3 绝对值操作
* 电机换向
  * 如遇到电机方向不对, 可切换电机方向
    * 设置方法: 修改参数P1.01 = 1 或 0
* 电子齿轮比
  * 电子齿轮比分子P3.08 电子齿轮比分母P3.10

## **目录构成**

* CAD - 电气原理图
  ![](/Picture/CAD.png)
* linuxcnc - 包含了所需的所有配置文件

## **项目采购清单**

声明: 与采购链接无商业关系, 请自行询价比价


| 采购物品           | 网址                                                                                                                                                                            | 型号                                                              | 数量                  | 备注                                                                       |
| -------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------- | ----------------------- | ---------------------------------------------------------------------------- |
| XY轴电机+驱动      | [URL](https://detail.tmall.com/item.htm?id=651348933042&spm=2015.23436601.0.0)                                                                                                  | (新款)400W伺服+驱动+5米线<br/>电机型号: RSM-M06L1330A(无抱闸)     | 2                     | 多圈编码器<br/>绝对原点,这个是要加钱的<br/>,5米线<br/>,支持EtherCAT+CIA402 |
| Z轴电机+驱动       | [URL](https://detail.tmall.com/item.htm?id=651348933042&spm=2015.23436601.0.0&skuId=4870002170807)                                                                              | (新款)RS400E伺服+驱动+5米线<br/>电机型号: RSM-M06L1330A-Z(带抱闸) | 1                     | 多圈编码器<br/>绝对原点,这个是要加钱的<br/>,5米线<br/>,支持EtherCAT+CIA402 |
| 主轴电机+驱动      | [URL](https://detail.tmall.com/item.htm?id=651348933042&spm=2015.23436601.0.0&skuId=4870002170807)                                                                              | (新款)RS400E伺服+驱动+5米线<br/>电机型号: RSM-M06J1330A(无抱闸)   | 1                     | 普通编码器<br/>,5米线<br/>,支持EtherCAT+CIA402                             |
| 树莓派_4代开发板   | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=608798378397&ns=1&skuId=4551930747475&spm=a230r.1.14.43.2fad28eea6m3ZB)                                                   | Pi 4B/8G[不可或缺套餐]                                            | 1                     |                                                                            |
| micro hdmi转hdmi线 | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=658377041318&ns=1&skuId=4747932408936&spm=a230r.1.14.7.399d1c02A92EDD)                                                    | 3米                                                               | 1                     |                                                                            |
| EtherCat伺服总线   | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=674202033100&ns=1&sku_properties=1627207:21423176984&spm=a230r.1.14.1.4e83345adWD5Zj)                                     | 高柔性千兆屏蔽网线-绿色                                           | 3(0.5米)<br/>1(1.5米) | -                                                                          |
| 树莓派_导轨盒      | [URL](https://item.taobao.com/item.htm?spm=a230r.1.14.7.ae487264WDE38q&id=663400858369&ns=1&abbucket=2#detail)                                                                  |                                                                   | 1                     |                                                                            |
| 显示器带壳         | [URL](https://detail.tmall.com/item.htm?id=666450227637&skuId=5093251227751&spm=pc_detail.27183998.202208.1.705a7dd62FFPiD)                                                     | 7寸1024*600IPS+触摸+外壳                                          | 1                     | 不推荐,<br/>使用1个月不到就坏了,请自行选购                                 |
| 显示器支架         | [URL](https://detail.tmall.com/item.htm?abbucket=17&id=687974425740&rn=c8d5c0a97e71825d8e11d5a7d7fd0f17&spm=a1z10.5-b.w4011-23875343074.107.14993b2bXRLXEy&skuId=4901270192689) |                                                                   | 1                     |                                                                            |
| 手轮               | [URL](https://item.taobao.com/item.htm?spm=a21n57.1.0.0.10a6523cQWLmki&id=45130365561&ns=1&abbucket=0#detail)                                                                   | MACH3有线手轮4轴                                                  | 1                     |                                                                            |
| 导轨插座           | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=677701080516&ns=1&spm=a21n57.1.0.0.162c523chCpKY3&skuId=5038263790059)                                                    | 5孔                                                               | 1                     |                                                                            |
| USB线              | [URL](https://item.taobao.com/item.htm?spm=a21n57.1.0.0.19d6523c0w91YJ&id=542169862221&ns=1&abbucket=2#detail)                                                                  | Micro usb数据连接线黑色 5米                                       | 1                     |                                                                            |
| 轴流风机           | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=658927980547&ns=1&skuId=4759805705081&spm=a21n57.1.0.0.6a79523cvWfTb0)                                                    | 阻燃ZL-803+12038风机 AC220V+防护网                                | 2                     |                                                                            |
| 急停按钮           | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=552913814211&ns=1&spm=a21n57.1.0.0.347b523ctQ5vgw)                                                                        |                                                                   |                       |                                                                            |
| 断路器(施耐德)     |                                                                                                                                                                                 | IC65N-C10A/2P 10A<br/>IC65N-C6A/2P 6A                             | 1<br/>2               | 请自行选购                                                                 |
| 24V直流电源        |                                                                                                                                                                                 | 明伟电源, 120瓦                                                   | 1                     | 请自行选购                                                                 |

## 特别感谢

* 感谢桂林广陆数字测控有限公司对此开源项目的支持

  * 天猫旗舰店 - [产品购买地址](https://guanglujj.tmall.com/shop/view_shop.htm?spm=a21n57.1.0.0.396d523cnfFegq&appUid=RAzN8HWJMBXJctTqz11nosKY98Wm2AwEqKJEpqoPs9YQHKEVNDc)
    ![](/Picture/guilinguanglu.png)
* 感谢在开发过程中帮助过我的开发者们[Hakan](https://forum.linuxcnc.org/cb-profile/22448-hakan) & [rodw](https://forum.linuxcnc.org/cb-profile/rodw) & [garlicbread](https://forum.linuxcnc.org/cb-profile/garlicbread)
* 项目引用

  * [YouTube - Ethercat + linuxcnc + Raspberry pi?](https://www.youtube.com/watch?v=NQ-HnrusGJo&t=16s)
